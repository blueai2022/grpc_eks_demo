package api

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func setupProxy() gin.HandlerFunc {

	target := "api.lifeai.us"

	return func(ctx *gin.Context) {
		director := func(req *http.Request) {
			r := ctx.Request

			req.URL.Scheme = "https"
			req.URL.Host = target
			req.Host = target
			req.Method = r.Method

			backPath := fmt.Sprintf("/backend%s", r.URL.Path)
			req.URL.Path = backPath
			req.RequestURI = backPath

			// req.Header["my-header"] = []string{r.Header.Get("my-header")}
			// Golang camelcases headers
			delete(req.Header, "Authorization")
			// log.Printf("orig req details: %v", ctx.Request)
			// log.Printf("req details: %v", req)
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
		// log.Printf("proxy called %v", proxy.Director)
	}
}
