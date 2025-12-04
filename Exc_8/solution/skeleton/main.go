package main

import (
	"exc8/client"
	"exc8/server"
	"time"
)

func main() {
	go func() {
		// todo start server
		if err := server.StartGrpcServer(); err != nil {
			println("server error:", err.Error())
		}
	}()
	time.Sleep(1 * time.Second)
	// todo start client
	c, err := client.NewGrpcClient()
	if err != nil {
		println("client error:", err.Error())
		return
	}
	if err := c.Run(); err != nil {
		println("run error:", err.Error())
		return
	}
	println("Orders complete!")
}
