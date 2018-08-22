package main

import (
	"flag"
	"order-now/order-api/api"

	"io/ioutil"
	"runtime"

	"order-now/order-api/config"

	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	params  config.Params
	devMode = flag.Bool("dev", false, "-dev Use in developer mode ")
)

func main() {

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.Default()
	router.Use(gin.Recovery())

	setConfiguration()

	api.Setup(router)
	logo()
	router.Run(":" + config.Get().ApiPort)
}

func setConfiguration() {
	params = config.Params{
		DevMode: *devMode,
	}
	config.SetConfig(&params)
}

func logo() {
	l := `

  /$$$$$$                  /$$                              /$$$$$$            /$$
 /$$__  $$                | $$                             /$$__  $$          |__/
| $$  \ $$  /$$$$$$   /$$$$$$$  /$$$$$$   /$$$$$$         | $$  \ $$  /$$$$$$  /$$
| $$  | $$ /$$__  $$ /$$__  $$ /$$__  $$ /$$__  $$ /$$$$$$| $$$$$$$$ /$$__  $$| $$
| $$  | $$| $$  \__/| $$  | $$| $$$$$$$$| $$  \__/|______/| $$__  $$| $$  \ $$| $$
| $$  | $$| $$      | $$  | $$| $$_____/| $$              | $$  | $$| $$  | $$| $$
|  $$$$$$/| $$      |  $$$$$$$|  $$$$$$$| $$              | $$  | $$| $$$$$$$/| $$
 \______/ |__/       \_______/ \_______/|__/              |__/  |__/| $$____/ |__/
                                                                    | $$          
                                                                    | $$          
                                                                    |__/
`
	fmt.Println(l)
}
