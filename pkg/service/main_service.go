package service

import (
	"ftms-go/pkg/repository"
	"io"
	"mime/multipart"
	"os"
	"reflect"
)

func GetImageProfile(username string) (string, error) {
	appUser, err := repository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	return appUser.ImageProfile, nil
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func uploadImage(file multipart.File, fName *multipart.FileHeader, empid string) (string, error) {
	mkdirPath := "./images-profile/" + empid
	os.MkdirAll(mkdirPath, os.ModePerm)
	pathName := empid + "/" + fName.Filename
	drFile, err := os.OpenFile("./images-profile/"+pathName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(drFile, file)
	if err != nil {
		return "", err
	}
	defer drFile.Close()
	return pathName, nil
}
