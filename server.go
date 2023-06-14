package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kpango/glg"
)

func init() {
	glg.Get().SetMode(glg.BOTH).SetLineTraceMode(glg.TraceLineShort)
}

func startHttpServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("欢迎使用，正常运行"))
	})
	http.HandleFunc("/ok", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK了"))
	})
	http.HandleFunc("/vip", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("vip"))
	})
	var port string
	if key := os.Getenv("PORT_ENV_KEY"); key != "" {
		port = os.Getenv(key)
	}
	if port == "" {
		port = "8006"
	}
	glg.Printf("使用方式 http://host:%v/", port)
	glg.Error(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func main() {
	startHttpServer()
}
