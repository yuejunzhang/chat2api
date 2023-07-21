package main

import (
	"chat2api/config"
	"chat2api/v1/chat"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	s.SetPort(config.PORT)
	chatGroup := s.Group("/v1/chat")
	chatGroup.Middleware(MiddlewareCORS)

	chatGroup.ALL("/completions", chat.Completions)
	s.Run()
}
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
