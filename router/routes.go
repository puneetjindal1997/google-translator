package router

import (
	"net/http"
	"translate/constant"
	"translate/controller"
)

// this routes represent user
var translateRoutes = Routes{
	Route{"Translate from to language", http.MethodPost, constant.TranslatePath, controller.Translate},
}
