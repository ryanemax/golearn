//已经尝试实用ParseForm解析Get、Post请求及字符串拼接Json，转向encoding/json直接使用
//已经尝试手动配置CORS，转向CORS直接解决跨域问题
//已经学习参考过net及net/http，转向go-json-rest继续进行RESTful工程开发
package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
	fmt.Println("User has Connected!\n")
}
func AuthLogin(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var rst string
	if len(req.Form["username"]) > 0 && len(req.Form["password"]) > 0 {

		rst = req.Host + "\nUserName:" + req.Form["username"][0] + "\nPassWord:" + req.Form["password"][0]
	} else {
		rst = "else"
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Write([]byte(rst))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth/login", AuthLogin)
	//http.HandleFunc("/auth/login", AuthLogin)
	//http.HandleFunc("/hello", SayHello)
	http.Handle("/", &MyServer{r})
	//header for CORS
	//error http.Header.Set("Access-Control-Allow-Origin", "*")
	http.ListenAndServe(":8001", nil)
}

//config the CORS HTTP Headers
type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}
