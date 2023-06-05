package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name       string
	Method     string
	Path       string // localhost:8000/helloworld
	HandleFunc func(*gin.Context)
}

type routes struct {
	router *gin.Engine
}

type Routes []Route

func (r routes) Translator(g *gin.RouterGroup) {
	group := g.Group("/" + os.Getenv("TranslateRoute"))
	for _, singleRoute := range translateRoutes {
		switch singleRoute.Method {
		case http.MethodPost:
			group.POST(singleRoute.Path, singleRoute.HandleFunc)
		case http.MethodGet:
			group.GET(singleRoute.Path, singleRoute.HandleFunc)
		case http.MethodPut:
			group.PUT(singleRoute.Path, singleRoute.HandleFunc)
		case http.MethodDelete:
			group.DELETE(singleRoute.Path, singleRoute.HandleFunc)
		}
	}
}

func Routing() {
	r := routes{
		router: gin.Default(),
	}
	grouping := r.router.Group(os.Getenv("ApiVersion"))
	r.Translator(grouping)
	r.router.Run(":" + os.Getenv("Port"))
}
