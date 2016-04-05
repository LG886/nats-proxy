package natsproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nats-io/nats"
)

// TestProxy integration test to
// test complete proxy and client.
func TestProxy(t *testing.T) {

	var reqEvent string

	// Initialize NATS client
	//
	clientConn, _ := nats.Connect(nats_url)
	natsClient, _ := NewNatsClient(clientConn)
	natsClient.Use(func(c *Context) {

		c.Response.Header["Middleware"] = &Values{
			[]string{"Mok"},
		}

		if c.HeaderVariable("X-Auth") == "" {
			c.AbortWithJSON("Not authenticated")
		}
	})
	natsClient.Subscribe("POST", "/test/:event/:session", func(c *Context) {
		c.ParseForm()
		reqEvent = c.PathVariable("event")

		if reqEvent != "12324" {
			fmt.Println("ReqEvent: " + reqEvent)
			t.Error("Path variable doesn't match")
		}

		// Assert that the form
		// is also parsed for the
		// query params

		nameVal := c.FormVariable("name")
		if nameVal != "testname" {
			t.Error("Form value assertion failed")
		}

		// Assets that the form params
		// are also parsed for post forms
		nameVal = c.FormVariable("post")
		if nameVal != "postval" {
			fmt.Println("postval: " + nameVal)
			t.Error("Form value assertion failed")
		}

		// Assert method
		if c.Request.Method != "POST" {
			t.Error("Method assertion failed")
		}

		if reqEvent != "12324" {
			fmt.Println(reqEvent)
			t.Error("Event path variable assertion failed")
		}

		respStruct := struct {
			User string
		}{
			"Radek",
		}

		if v, ok := c.Request.Header["X-Auth"]; ok {
			if len(v.Arr) == 0 || v.Arr[0] != "xauthpayload" {
				t.Error("Header assertion failed")
			}
		}

		formVal := c.FormVariable("both")

		if formVal != "y" {
			fmt.Println(c.Request.Form)
			t.Error("Form assertion failed")
		}

		// Generate response
		c.JSON(200, respStruct)
		c.Response.GetHeader().Set("X-Auth", "12345")
	})
	defer clientConn.Close()

	proxyConn, _ := nats.Connect(nats_url)
	proxyHandler, _ := NewNatsProxy(proxyConn)
	proxyHandler.AddHook(".*", func(r *Response) {
		r.GetHeader().Set("Hook", "Hok")
	})
	defer proxyConn.Close()

	reader := strings.NewReader("post=postval&both=y")
	req, _ := http.NewRequest("POST", "http://127.0.0.1:3000/test/12324/123?name=testname&both=n", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("X-AUTH", "xauthpayload")

	rw := httptest.NewRecorder()
	proxyHandler.ServeHTTP(rw, req)

	if rw.Header().Get("Middleware") != "Mok" {
		t.Error("Middleware usage assertion failed")
	}
	if rw.Header().Get("Hook") != "Hok" {
		t.Error("Hook usage assertion failed")
	}

	out, _ := ioutil.ReadAll(rw.Body)
	respStruct := &struct {
		User string
	}{}

	json.Unmarshal(out, respStruct)
	if respStruct.User != "Radek" {
		t.Error("Response assertion failed")
	}

	//Test aborting request
	reader = strings.NewReader("post=postval")
	req, _ = http.NewRequest("POST", "http://127.0.0.1:3000/test/12324/123?name=testname", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	rw = httptest.NewRecorder()
	proxyHandler.ServeHTTP(rw, req)

	out, _ = ioutil.ReadAll(rw.Body)
	if string(out) != "Not authenticated" && rw.Code != 500 {
		t.Errorf("Abort assertion failed code: %d , resp: %s", rw.Code, string(out))
	}

}

func TestProxyServeHttpError(t *testing.T) {
	proxyConn, _ := nats.Connect(nats_url)
	proxyHandler, _ := NewNatsProxy(proxyConn)
	defer proxyConn.Close()
	rw := httptest.NewRecorder()
	proxyHandler.ServeHTTP(rw, nil)

	if rw.Code != http.StatusInternalServerError {
		t.Error()
	}

	req, _ := http.NewRequest("", "", nil)
	proxyHandler.ServeHTTP(rw, req)
	if rw.Code != http.StatusInternalServerError {
		t.Error()
	}
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}
