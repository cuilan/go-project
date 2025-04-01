package http

import (
	"go-project/internal/conf"
	"io"
	"net/http"
	"testing"
)

var host = "http://localhost:"

func setup() {
	conf.Unmarshal("../../configs")

	go Server(conf.HttpPort, conf.Mode)
}

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func request(t *testing.T, uri string) string {
	setup()

	resp, err := http.Get(host + conf.HttpPort + uri)
	handleError(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	handleError(t, err)
	return string(body)
}

func TestServer(t *testing.T) {
	jsonStr := request(t, Root)
	t.Log(jsonStr)
}

func TestHealth(t *testing.T) {
	jsonStr := request(t, PushGroup+Health)
	t.Log(jsonStr)
}
