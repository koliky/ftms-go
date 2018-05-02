package service

import (
	"ftms-go/pkg/repository"
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
