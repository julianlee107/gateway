package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	realOne := &RealServer{Addr: "127.0.0.1:2003"}
	realTwo := &RealServer{Addr: "127.0.0.1:8000"}
	realOne.Run()
	realTwo.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

type RealServer struct {
	Addr string
}

func (r *RealServer) Run() {
	log.Println("starting httpserver at " + r.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.HelloHandler)
	mux.HandleFunc("/base/error/", r.ErrorHandler)
	mux.HandleFunc("/base/timeout/", r.TimeoutHandler)
	server := &http.Server{
		Addr:         r.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}

func (r *RealServer) HelloHandler(w http.ResponseWriter, req *http.Request) {

	upath := fmt.Sprintf("http://%s%s \n", r.Addr, req.URL.Path)
	realIp := fmt.Sprintf("RemoteAddr=%s,X-Forwarded-For=%s,X-Real-IP=%v\n", req.RemoteAddr, req.Header.Get("X-Forwarded-For"), req.Header.Get("X-Real-Ip"))
	header := fmt.Sprintf("headers=%v\n", req.Header)
	io.WriteString(w, upath)
	io.WriteString(w, realIp)
	io.WriteString(w, header)
}

func (r *RealServer) ErrorHandler(w http.ResponseWriter, req *http.Request) {
	upath := "error handler"
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, upath)
}

func (r *RealServer) TimeoutHandler(w http.ResponseWriter, rep *http.Request) {
	time.Sleep(6 * time.Second)
	upath := "timeout handler"
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, upath)

}
