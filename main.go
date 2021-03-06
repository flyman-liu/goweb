package main

import (
	"net/http"

	"asyons.com/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("public"))
	router.GET("/", controllers.Index)
	router.GET("/item/:id", controllers.Detail)
	router.GET("/logout", controllers.Logout)
	router.GET("/reg", controllers.Reg)

	router.POST("/regpost", controllers.RegPost)
	router.POST("/login", controllers.Login)

	router.GET("/add", controllers.Add)
	router.POST("/upload", basicAuth(controllers.Upload))
	router.POST("/post", basicAuth(controllers.Post))

	router.GET("/user/:id", controllers.User)
	router.POST("/search", controllers.Search)

	http.ListenAndServe(":80", router)
}

func basicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if ok := check(r); ok {
			h(w, r, ps)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func check(r *http.Request) bool {
	cookie, _ := r.Cookie("username")
	if cookie == nil {
		return false
	}
	return true
}
