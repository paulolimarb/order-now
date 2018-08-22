package main

import (
	"flag"
	"fmt"
	"order-now/order-consumer/config"
	"order-now/order-consumer/queue"
	"runtime"
)

var (
	params  config.Params
	devMode = flag.Bool("dev", false, "-dev Use in developer mode ")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	setConfiguration()
	logo()
	queue.ReadMessage()
}

func setConfiguration() {
	params = config.Params{
		DevMode: *devMode,
	}
	config.SetConfig(&params)
}

func logo() {
	l := `
  _____   ______ ______  _______  ______     _______  _____  __   _ _______ _     _ _______ _______  ______
 |     | |_____/ |     \ |______ |_____/ ___ |       |     | | \  | |______ |     | |  |  | |______ |_____/
 |_____| |    \_ |_____/ |______ |    \_     |_____  |_____| |  \_| ______| |_____| |  |  | |______ |    \_
`
	fmt.Println(l)
}
