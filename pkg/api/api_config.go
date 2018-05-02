package api

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func NewAPI(router rest.App) (api *rest.Api) {
	api = rest.NewApi()
	// api.Use(rest.DefaultDevStack...)
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders := []string{
		"Accept",
		"Authorization",
		"X-Real-IP",
		"Content-Type",
		"X-Custom-Header",
		"Language",
		"Origin",
	}
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
		AllowedMethods:                allowedMethods,
		AllowedHeaders:                allowedHeaders,
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})
	loginMiddle := &LoginMiddleware{}
	api.Use(loginMiddle)
	api.SetApp(router)
	return api
}

type LoginMiddleware struct {
}

func (login *LoginMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		handler(w, r)
	}
}
