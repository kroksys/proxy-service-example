package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/kroksys/proxy-service-example/src/utils"
)

// Start proxy http server using provided configuration.
// Starts logging to the file based on config.Proxy.Log file name.
func ListenAndServe(cfg *utils.Config) error {

	// parse target url
	targetUrl, err := url.Parse(cfg.Proxy.Target)
	if err != nil {
		return err
	}

	// start loging to a file
	f, err := os.OpenFile(cfg.Proxy.Log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/", proxy(targetUrl).ServeHTTP)
	return http.ListenAndServe(cfg.Proxy.Listen, nil)
}

// Proxy with logging and SSL redirection.
func proxy(targetUrl *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.Director = func(request *http.Request) {
		log.Printf("[REQUEST] %+v", request)
		request.URL.Host = targetUrl.Host
		request.URL.Scheme = targetUrl.Scheme
		request.Header.Set("X-Forwarded-Host", request.Header.Get("Host"))
		request.Host = targetUrl.Host
	}
	proxy.ModifyResponse = func(response *http.Response) error {
		log.Printf("[RESPONSE] %+v", response)
		return nil
	}
	return proxy
}
