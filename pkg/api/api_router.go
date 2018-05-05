package api

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
)

func NewRouter() rest.App {
	router, err := rest.MakeRouter(
		rest.Get("/security/createuseradmin", createUserAdmin),
		rest.Get("/security/createdatatest", createDataTest),
		rest.Post("/security/login", login),

		rest.Get("/api/user/imageprofile/:username", imageProfile),
		rest.Post("/api/user/getdatauser", getDataUser),
		rest.Post("/api/user/updateprofile", userUpdateProfile),

		rest.Post("/api/admin/validateempid", adminValidateEmpId),
		rest.Post("/api/admin/createuser", adminCreateUser),
		rest.Post("/api/admin/findbyid", adminFindById),
	)
	if err != nil {
		log.Fatal(err)
	}
	return router
}
