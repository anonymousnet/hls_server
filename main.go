package main

import (
	_ "github.com/anonymousnet/hls_server/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

