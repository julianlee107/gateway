package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

var (
	proxyAddr = "http://127.0.0.1:2003"
	port      = "2002"
)

func main() {
	http.HandleFunc("/",handler)
	log.Println("start on serving on port"+port)
	err := http.Liste nAndServe(":"+port,nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter,r *http.Request)  {
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		panic("proxy can't be nil")
	}
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host

	transport := http.DefaultTransport
	resp,err := transport.RoundTrip(r)

	if err != nil {
		log.Print(err)
		return
	}

	for key,value := range resp.Header{
		for _,v := range value{
			w.Header().Add(key,v)
		}
	}
	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(w)

}