package getuser

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rocksun/mvclayout/pkg/application"
)

func Do(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "hello")
	}
}
