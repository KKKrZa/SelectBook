package Sever

import "net/http"

type CorsMiddleware struct {
	Next http.Handler
}

func (cors *CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cors.Next == nil {
		cors.Next = http.DefaultServeMux
	}
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许跨域
	cors.Next.ServeHTTP(w, r)
}
