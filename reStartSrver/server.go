/**
 * server.go：在4242端口启动一个HTTP服务
 * 设置endless服务，当该进程收到SIGHUP信号的时候，fork子进程接受新的连接，处理父进程现有任务并后退出
 * Create By ChenWenBo
 * 2019-10-15
 */
package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"net/http"
	"syscall"
)

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "Hello golang http!")
}

func main() {
	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", index)

	server := endless.NewServer(":4242", nil)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
		// save it somehow
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
