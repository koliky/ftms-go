package api

import (
	"ftms-go/pkg/service"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func createUserAdmin(w rest.ResponseWriter, r *rest.Request) {
	code := r.URL.Query().Get("code")
	if code != "pppassword" {
		respErr := map[string]string{}
		respErr["message"] = "Code invalid."
		w.WriteHeader(400)
		w.WriteJson(respErr)
		return
	}
	resp, status := service.CreateUserAdmin()
	w.WriteHeader(status)
	w.WriteJson(resp)
}

func adminValidateEmpId(w rest.ResponseWriter, r *rest.Request) {
	resp := map[string]string{}
	cms, err := service.CheckToken(r)
	if err != nil {
		resp["message"] = "JWT_invalid"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	if service.CheckRole("Admin", cms.Get("roles").(string)) {
		resp["message"] = "Invalid_role"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	body := map[string]interface{}{}
	err = r.DecodeJsonPayload(&body)
	if err != nil {
		w.WriteHeader(400)
		w.WriteJson(map[string]string{"error": err.Error()})
	}
	data, err := service.GetByEmployeeId(body["employeeId"].(string))
	if err != nil {
		resp["message"] = "invalid user"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	w.WriteHeader(200)
	w.WriteJson(data)
}

func adminCreateUser(w rest.ResponseWriter, r *rest.Request) {
	resp := map[string]interface{}{}
	cms, err := service.CheckToken(r)
	if err != nil {
		resp["message"] = "JWT_invalid"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	if service.CheckRole("Admin", cms.Get("roles").(string)) {
		resp["message"] = "Invalid_role"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	body := map[string]interface{}{}
	err = r.DecodeJsonPayload(&body)
	if err != nil {
		w.WriteHeader(400)
		w.WriteJson(map[string]string{"error": err.Error()})
	}
	resp, status := service.AdminCreateUser(body)
	w.WriteHeader(status)
	w.WriteJson(resp)
}

func adminFindById(w rest.ResponseWriter, r *rest.Request) {
	resp := map[string]interface{}{}
	cms, err := service.CheckToken(r)
	if err != nil {
		resp["message"] = "JWT_invalid"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	if service.CheckRole("Admin", cms.Get("roles").(string)) {
		resp["message"] = "Invalid_role"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	body := map[string]interface{}{}
	err = r.DecodeJsonPayload(&body)
	if err != nil {
		w.WriteHeader(400)
		w.WriteJson(map[string]string{"error": err.Error()})
	}
	strId := body["id"].(string)
	id, _ := strconv.Atoi(strId)
	appUser, err := service.GetById(id)
	if err != nil {
		resp["message"] = "Invalid_user"
		w.WriteHeader(400)
		w.WriteJson(resp)
		return
	}
	w.WriteHeader(200)
	w.WriteJson(appUser)
}

func createDataTest(w rest.ResponseWriter, r *rest.Request) {
	code := r.URL.Query().Get("code")
	if code != "pppassword" {
		respErr := map[string]string{}
		respErr["message"] = "Code invalid."
		w.WriteHeader(400)
		w.WriteJson(respErr)
		return
	}
	resp, status := service.CreateTempData()
	w.WriteHeader(status)
	w.WriteJson(resp)
}
