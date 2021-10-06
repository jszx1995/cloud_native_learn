package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//1.接收客户端 request，并将 request 中带的 header 写入 response header
	//fmt.Fprintf(w, "Header = %q\n", r.Header)
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}

	//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	//fmt.Fprintf(w, "VERSION = %q\n", os.Getenv("VERSION"))
	io.WriteString(w, fmt.Sprintf("%s=%s\n", "VERSION", os.Getenv("VERSION")))

	//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	log.Printf("%s,%s", r.RemoteAddr, r.URL) // HTTP 返回码

}

//4.当访问 localhost/healthz 时，应返回200
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	io.WriteString(w, "200")
}
