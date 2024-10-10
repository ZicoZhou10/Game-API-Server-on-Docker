package main

import (
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Any("/players", reverseProxy("http://player_management:8080"))
	r.Any("/players/*path", reverseProxy("http://player_management:8080"))
	r.Any("/rooms", reverseProxy("http://room_management:8080"))
	r.Any("/rooms/*path", reverseProxy("http://room_management:8080"))
	r.Any("/challenges", reverseProxy("http://challenge_system:8080"))
	r.Any("/challenges/*path", reverseProxy("http://challenge_system:8080"))
	r.Any("/logs", reverseProxy("http://log_collector:8080"))
	r.Any("/logs/*path", reverseProxy("http://log_collector:8080"))
	r.Any("/payments", reverseProxy("http://payment_system:8080"))
	r.Any("/payments/*path", reverseProxy("http://payment_system:8080"))

	r.Run(":8080")
}

func reverseProxy(target string) gin.HandlerFunc {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)

	return func(c *gin.Context) {
		// Remove the trailing slash if it exists and it's not the root path
		if c.Request.URL.Path != "/" {
			c.Request.URL.Path = strings.TrimSuffix(c.Request.URL.Path, "/")
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
