package web

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/api", apiHandler)
	return router
}

func main() {
	fmt.Println("Start monitor...")
	fmt.Println("Open http://127.0.0.1:20001")
	handles := RegisterHandlers()
	http.ListenAndServe(":20001", handles)
}
