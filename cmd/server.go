package main

import (
	"encoding/json"
	"fmt"
	"ftms-go/pkg/api"
	"ftms-go/pkg/repository"
	"io/ioutil"
	"net/http"
)

var jsonData []byte

func main() {
	loadConfig()
	repository.SetupDatabase(getConfig("databaseURL"), getConfig("ddl-auto"))
	api := api.NewAPI(api.NewRouter())
	fmt.Println("server start... => http://localhost:" + getConfig("port"))
	http.ListenAndServe(":"+getConfig("port"), api.MakeHandler())
}

func loadConfig() {
	loadConfg, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	jsonData = loadConfg
}

func getConfig(key string) string {
	var objmap map[string]*json.RawMessage
	json.Unmarshal(jsonData, &objmap)
	var value string
	json.Unmarshal(*objmap[key], &value)
	return value
}
