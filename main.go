package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 解析命令行参数
	// 指定监听的端口
	var port int
	flag.IntVar(&port, "port", 80, "Port to serve on")
	flag.IntVar(&port, "p", 80, "Port to serve on")
	flag.Parse()
	log.Printf("The port is %d\n", port)

	// 获取程序当前所在的目录
	dir, err := os.Getwd()
	log.Printf("The work dir is %s\n", dir)

	if err != nil {
		log.Fatal(err)
	}
	// 将文件目录作为静态文件服务器的根目录
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	// 启动HTTP服务器
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Server started on port %d\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
