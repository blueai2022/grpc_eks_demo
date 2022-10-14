package api

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

const (
	backChanProtocol    = "http"
	jsonContentType     = "application/json"
	authorizationHeader = "Authorization"
	backUrlPathPrefix   = "/backend"
)

func setupGinProxy(target string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		director := func(req *http.Request) {
			r := ctx.Request

			req.Method = r.Method

			req.URL.Scheme = backChanProtocol
			req.URL.Host = target
			req.Host = target

			backPath := backUrlPathPrefix + r.URL.Path
			req.URL.Path = backPath
			req.RequestURI = backPath

			delete(req.Header, authorizationHeader)
			// log.Printf("orig req: %v", ctx.Request)
			// log.Printf("outbound req: %v", req)

		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
