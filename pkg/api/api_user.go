package api

import (
	"ftms-go/pkg/service"

	"github.com/ant0ine/go-json-rest/rest"
)

func getDataUser(w rest.ResponseWriter, r *rest.Request) {
	resp := map[string]string{}
	cms, err := service.CheckToken(r)
	if err != nil {
		resp["message"] = "JWT_invalid"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	strUsername := cms.Get("username").(string)
	data, err := service.GetUserByUsername(strUsername)
	if err != nil {
		resp["message"] = "invalid user"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	w.WriteHeader(200)
	w.WriteJson(data)
}

func userUpdateProfile(w rest.ResponseWriter, r *rest.Request) {
	_, err := service.CheckToken(r)
	if err != nil {
		resp := map[string]string{}
		resp["message"] = "JWT_invalid"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	res, status := service.UserUpdate(r)
	w.WriteHeader(status)
	w.WriteJson(res)
}
