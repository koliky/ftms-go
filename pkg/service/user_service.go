package service

import (
	"ftms-go/pkg/entity"
	"ftms-go/pkg/repository"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func CreateUserAdmin() (map[string]string, int) {
	resp := map[string]string{}
	err := repository.CreateUserAdmin()
	if err != nil {
		resp["message"] = err.Error()
		return resp, 400
	}
	resp["message"] = "Create user admin successfully"
	return resp, 200
}

func CreateTempData() (map[string]string, int) {
	resp := map[string]string{}
	for i := 1; i <= 53; i++ {
		appUser := entity.AppUser{}
		appUser.Address = "foamtec" + strconv.Itoa(i)
		appUser.Status = "ADMIN_CREATE"
		appUser.CreateDate = "19/05/1991"
		appUser.Department = "MIS" + strconv.Itoa(i)
		appUser.Email = "aaa@g.c" + strconv.Itoa(i)
		appUser.EmployeeID = "906" + strconv.Itoa(i)
		appUser.Sex = "Male"
		appUser.FirstName = "Test" + strconv.Itoa(i)
		appUser.LastName = "Last" + strconv.Itoa(i)
		appUser.PhoneNumber = "8" + strconv.Itoa(i)
		appUser.Username = "906" + strconv.Itoa(i)
		appUser.Password = "906" + strconv.Itoa(i)
		appUser.Shift = "A"
		appUser.ImageProfile = "user.png"
		appUser.AppRoles = []string{"User"}
		_, err := repository.CreateUser(appUser)
		if err != nil {
			resp["message"] = err.Error()
			return resp, 400
		}
	}
	resp["message"] = "Create data test successfully"
	return resp, 200
}

func GetUserByUsername(username string) (entity.AppUser, error) {
	return repository.GetUserByUsername(username)
}

func GetByEmployeeId(empId string) (entity.AppUser, error) {
	return repository.GetByEmployee(empId)
}

func GetById(id int) (entity.AppUser, error) {
	return repository.GetById(id)
}

func AdminCreateUser(data map[string]interface{}) (map[string]interface{}, int) {
	appUser := entity.AppUser{}
	appUser.EmployeeID = data["employeeId"].(string)
	appUser.FirstName = data["firstName"].(string)
	appUser.LastName = data["lastName"].(string)
	appUser.Sex = data["sex"].(string)
	appUser.Department = data["department"].(string)
	appUser.Shift = data["shift"].(string)
	appUser.StartDate = data["startDate"].(string)
	appUser.Username = data["employeeId"].(string)
	appUser.ImageProfile = "user.png"
	appUser.Password = data["employeeId"].(string)
	roles := []string{}
	itf := InterfaceSlice(data["roles"])
	for _, role := range itf {
		roles = append(roles, role.(string))
	}
	appUser.AppRoles = roles
	id, err := repository.CreateUser(appUser)
	resp := map[string]interface{}{}
	if err != nil {
		resp["message"] = err.Error()
		return resp, 400
	}
	resp["id"] = id
	return resp, 200
}

func UserUpdate(r *rest.Request) (map[string]interface{}, int) {
	resp := map[string]interface{}{}
	appUser := entity.AppUser{}
	appUser.EmployeeID = r.FormValue("employeeId")
	appUser.FirstName = r.FormValue("firstName")
	appUser.LastName = r.FormValue("lastName")
	appUser.Password = r.FormValue("password")
	appUser.Email = r.FormValue("email")
	appUser.Address = r.FormValue("address")
	appUser.PhoneNumber = r.FormValue("phone")
	appUser.Status = "UPDATE"

	file, typeHeader, err := r.FormFile("fileImage")
	if file != nil {
		pathName, err := uploadImage(file, typeHeader, appUser.EmployeeID)
		if err != nil {
			resp["message"] = "file error"
			return resp, 400
		}
		appUser.ImageProfile = pathName
	}
	if err != nil {
		resp["message"] = "file error"
		return resp, 400
	}

	if err != nil {
		resp["message"] = err.Error()
		return resp, 400
	}

	_, err = repository.UpdateUser(appUser)
	if err != nil {
		resp["message"] = err.Error()
		return resp, 400
	}
	defer file.Close()
	resp["message"] = "success"
	return resp, 200
}
