package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rocksun/mvclayout/cmd/api/handlers/getuser"
	"github.com/rocksun/mvclayout/pkg/application"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/users", getuser.Do(app))
	return mux
}
