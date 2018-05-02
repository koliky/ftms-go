package api

import (
	"ftms-go/pkg/service"
	"io"
	"net/http"
	"os"

	"github.com/ant0ine/go-json-rest/rest"
)

func imageProfile(w rest.ResponseWriter, r *rest.Request) {
	fileName, _ := service.GetImageProfile(r.PathParam("username"))
	out, err := os.Open("./images-profile/" + fileName)
	if err != nil {
		respErr := map[string]string{}
		respErr["message"] = "invalid profile"
		w.WriteHeader(400)
		w.WriteJson(respErr)
		return
	}
	defer out.Close()
	w.Header().Set("Content-Disposition", "inline; filename="+fileName)
	w.Header().Set("Content-Type", "image/png")
	io.Copy(w.(http.ResponseWriter), out)
}
