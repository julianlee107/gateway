package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)
// 函数是一等公民，可作为一个参数进行传递
type FunctionHandler func(http.ResponseWriter, *http.Request)

func (f FunctionHandler) ServerHttp(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func main()  {
	function := FunctionHandler(HelloWorld)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET","/",bytes.NewBuffer([]byte("test")))
	function.ServerHttp(res,req)
	bts,_ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bts))
}


func HelloWorld(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("hello world"))

}