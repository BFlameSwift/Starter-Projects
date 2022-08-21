package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"hertz-demo/router"
)

func main() {
	h := server.Default()
	_ = router.InitRouter(*h)
	h.Spin()
}
