package router

import "github.com/gin-gonic/gin"

func DefineRoutes() {
	r := gin.Default()
	testController(r)
	r.Run()
}
