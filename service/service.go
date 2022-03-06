package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// 通用service
func Start(ctx context.Context, serviceName, host, port string, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()                           //注册service
	ctx = startService(ctx, serviceName, host, port) //启动service

	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port
	go func() {
		log.Panicln(srv.ListenAndServe())
		cancel()
	}()
	go func() {
		fmt.Printf("%v started. Press any key to stop. \n", serviceName)
		var s string
		fmt.Scanln(&s) //给用户选项，可以手动停止掉service
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
