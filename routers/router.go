package routers

import (
	"github.com/anonymousnet/hls_server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
