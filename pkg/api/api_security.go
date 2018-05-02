package api

import (
	"ftms-go/pkg/service"

	"github.com/ant0ine/go-json-rest/rest"
)

func login(w rest.ResponseWriter, r *rest.Request) {
	body := map[string]string{}
	err := r.DecodeJsonPayload(&body)
	if err != nil {
		w.WriteHeader(400)
		w.WriteJson(map[string]string{"error": err.Error()})
	}
	resp, status := service.CheckLogin(body["username"], body["password"])
	w.WriteHeader(status)
	w.WriteJson(resp)
}
