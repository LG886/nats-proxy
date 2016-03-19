package natsproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/nats-io/nats"
)

func TestProxy(t *testing.T) {

	var reqEvent string
	var reqSession string

	clientConn, _ := nats.Connect(nats_url)
	natsClient, _ := NewNatsClient(clientConn)
	natsClient.Subscribe("POST", "/test/:event/:session", func(c *Context) {
		reqEvent = c.PathVariable("event")
		reqSession = c.PathVariable("session")

		fmt.Println(c.Request.GetURL())
		if reqEvent != "12324" {
			fmt.Println(reqEvent)
			t.Error("Event path variable assertion failed")
		}

		if reqSession != "123" {
			fmt.Println(reqSession)
			t.Error("Session path variable assertion failed")
		}

		respStruct := struct {
			User string
		}{
			"Radek",
		}

		c.JSON(200, respStruct)
		headerKey := "X-AUTH"
		c.Response.Header = &HeaderMap{
			Items: []*HeaderItem{
				&HeaderItem{
					Key:   &headerKey,
					Value: []string{"12345"},
				},
			},
		}
	})
	defer clientConn.Close()

	proxyConn, _ := nats.Connect(nats_url)
	proxyHandler, _ := NewNatsProxy(proxyConn)
	http.Handle("/", proxyHandler)
	defer proxyConn.Close()

	log.Println("initializing proxy")
	go http.ListenAndServe(":3000", nil)
	time.Sleep(1 * time.Second)

	log.Println("Posting request")
	reader := bytes.NewReader([]byte("testData"))
	resp, err := http.Post("http://127.0.0.1:3000/test/12324/123?name=testname", "multipart/form-data", reader)
	if err != nil {
		log.Println(err)
		t.Error("Cannot do post")
		return
	}

	out, _ := ioutil.ReadAll(resp.Body)
	respStruct := &struct {
		User string
	}{}

	json.Unmarshal(out, respStruct)
	log.Println(respStruct)
	if respStruct.User != "Radek" {
		t.Error("Response assertion failed")
	}

	if reqEvent != "12324" {
		t.Error("Path variable doesn't match")
	}
}
