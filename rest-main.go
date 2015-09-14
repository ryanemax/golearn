// Package  provides ...
// 182.92.5.210:4567

package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

type User struct {
	Id   string
	Name string
}

func AuthLogin(w rest.ResponseWriter, req *rest.Request) {
	user := User{
		Id:   req.PathParam("id"),
		Name: "ryn",
	}
	w.WriteJson(&user)
}
func RawWrite(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(req.PathParam("raw.json"))
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	//配置API跨域CORS
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, req *rest.Request) bool {
			return origin == req.Header.Get("Origin")
		},
		AllowedMethods: []string{"GET", "POST", "PUT"},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})

	//设置API基本Router
	router, err := rest.MakeRouter(
		rest.Post("/:raw.json", RawWrite),
		rest.Post("/auth/login/:id", AuthLogin),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
