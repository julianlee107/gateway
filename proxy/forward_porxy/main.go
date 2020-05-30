package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Proxy struct {
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Header)
	fmt.Printf("Received request %s %s %s \n", r.Method, r.Host, r.RemoteAddr)
	transport := http.DefaultTransport

	outReq := new(http.Request)
	// 拷贝请求对象
	outReq = r
	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err != nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ",") + "," + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)

		//	请求下游
		res, err := transport.RoundTrip(outReq)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		//	把下游请求内容放回给上游
		for key, value := range res.Header {
			for _, v := range value {
				w.Header().Add(key, v)
			}
		}
		w.WriteHeader(res.StatusCode)
		io.Copy(w, res.Body)
		res.Body.Close()
	}
}

func main() {
	fmt.Println("server on 8080")
	http.Handle("/", &Proxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
