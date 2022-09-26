package part2

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

/*
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthz 时，应返回 200
*/

func Start() {
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {

	// 接收客户端 request，并将 request 中带的 header 写入 response header
	for key := range r.Header {
		val := r.Header.Get(key)
		w.Header().Add(key, val)
	}

	// 设置环境值的值
	os.Setenv("VERSION", "Geek Test Value")
	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	w.Header().Set("Version", version)

	// 获取Client IP
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("get ip err:", err)
	}
	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	if net.ParseIP(ip) != nil {
		fmt.Println("IP is:", ip)
		log.Println("IP is:", ip)
	}

	fmt.Println("http Status Code:", http.StatusOK)
	log.Println("http Status Code:", http.StatusOK)

	// 当访问 localhost/healthz 时，应返回 200
	w.WriteHeader(http.StatusOK)
}
