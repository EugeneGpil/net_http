package app

import (
	"github.com/EugeneGpil/net_http/app/modules/middleware"
	"github.com/EugeneGpil/net_http/app/modules/pathParam"
)

func Run() {
	pathParam.Run()
	if false {
		middleware.Run()
	}
}
