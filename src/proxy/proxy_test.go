package proxy

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListenAndServe(t *testing.T) {
	expectedResponseBody := "I am the backend"
	targetServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expectedResponseBody))
	}))
	defer targetServer.Close()
	targetServerURL, _ := url.Parse(targetServer.URL)

	proxyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy(targetServerURL).ServeHTTP(w, r)
	}))
	defer proxyServer.Close()
	proxyServerURL, _ := url.Parse(proxyServer.URL)

	resp, err := http.Get(proxyServerURL.String())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(bodyBytes) != expectedResponseBody {
		t.Errorf("Incorect response body! Expected: %s Got: %s", expectedResponseBody, string(bodyBytes))
	}
}
