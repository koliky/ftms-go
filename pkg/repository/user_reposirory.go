package repository

import (
	"ftms-go/pkg/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(appUser entity.AppUser) (id int, err error) {
	db := getConnection()
	defer db.Close()
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(appUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	sqlQuery := `
	INSERT INTO APP_USER (
		ADDRESS,STATUS,CREATE_DATE,DEPARTMENT,EMAIL,
		EMPLOYEE_ID,FIRST_NAME,LAST_NAME,IMAGE_PROFILE,PASSWORD,
		PHONE_NUMBER,SEX,SHIFT,START_DATE,USERNAME
	)
	VALUES (
		?,?,?,?,?,
		?,?,?,?,?,
		?,?,?,?,?
	)`
	resAppUser, err := db.Exec(
		sqlQuery,
		appUser.Address,
		"ADMIN_CREATE",
		now,
		appUser.Department,
		appUser.Email,
		appUser.EmployeeID,
		appUser.FirstName,
		appUser.LastName,
		appUser.ImageProfile,
		hashedPassword,
		appUser.PhoneNumber,
		appUser.Sex,
		appUser.Shift,
		now,
		appUser.Username,
	)
	if err != nil {
		return 0, err
	}
	lastId, _ := resAppUser.LastInsertId()
	sqlQuery = `
	INSERT INTO APP_ROLE (
		CREATE_DATE,ROLE_NAME,FK_APP_USER_ID
	)
	VALUES (
		?,?,?
	)`

	for _, role := range appUser.AppRoles {
		_, err = db.Exec(
			sqlQuery,
			now,
			role,
			int(lastId),
		)
		if err != nil {
			return 0, err
		}
	}

	return int(lastId), nil
}

func CreateUserAdmin() error {
	db := getConnection()
	defer db.Close()
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("adminpassword"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	sqlQuery := `
	INSERT INTO APP_USER (
		ADDRESS,STATUS,CREATE_DATE,DEPARTMENT,EMAIL,
		EMPLOYEE_ID,FIRST_NAME,LAST_NAME,IMAGE_PROFILE,PASSWORD,
		PHONE_NUMBER,SEX,SHIFT,START_DATE,USERNAME
	)
	VALUES (
		?,?,?,?,?,
		?,?,?,?,?,
		?,?,?,?,?
	)`
	resAppUser, err := db.Exec(
		sqlQuery,
		"foamtec address",
		"ADMIN_CREATE",
		now,
		"MIS",
		"apichate@foamtecintl.com",
		"00000",
		"Admin",
		"Foamtec",
		"admin.png",
		hashedPassword,
		"814",
		"Male",
		"Office",
		now,
		"admin",
	)
	if err != nil {
		return err
	}
	lastId, _ := resAppUser.LastInsertId()
	sqlQuery = `
	INSERT INTO APP_ROLE (
		CREATE_DATE,ROLE_NAME,FK_APP_USER_ID
	)
	VALUES (
		?,?,?
	)`
	_, err = db.Exec(
		sqlQuery,
		now,
		"Admin",
		int(lastId),
	)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (entity.AppUser, error) {
	db := getConnection()
	defer db.Close()
	appUser := entity.AppUser{}
	sqlQuery := `SELECT
		ID,ADDRESS,STATUS,CREATE_DATE,DEPARTMENT,
		EMAIL,EMPLOYEE_ID,FIRST_NAME,LAST_NAME,IMAGE_PROFILE,
		PASSWORD,PHONE_NUMBER,SEX,SHIFT,START_DATE,
		USERNAME
		FROM APP_USER
		WHERE USERNAME = ?
	`
	rowsUser, err := db.Query(sqlQuery, username)
	if err != nil {
		return appUser, err
	}
	for rowsUser.Next() {
		rowsUser.Scan(
			&appUser.ID, &appUser.Address, &appUser.Status, &appUser.CreateDate, &appUser.Department,
			&appUser.Email, &appUser.EmployeeID, &appUser.FirstName, &appUser.LastName, &appUser.ImageProfile,
			&appUser.Password, &appUser.PhoneNumber, &appUser.Sex, &appUser.Shift, &appUser.StartDate,
			&appUser.Username,
		)
	}
	rowsUser.Close()
	sqlQuery = `SELECT ROLE_NAME
		FROM APP_ROLE
		WHERE ID = ?
	`
	rowsRoles, err := db.Query(sqlQuery, appUser.ID)
	if err != nil {
		return appUser, err
	}
	for rowsRoles.Next() {
		var roleName string
		rowsRoles.Scan(&roleName)
		appUser.AppRoles = append(appUser.AppRoles, roleName)
	}
	rowsRoles.Close()
	return appUser, nil
}

func GetByEmployee(empId string) (entity.AppUser, error) {
	db := getConnection()
	defer db.Close()
	appUser := entity.AppUser{}
	sqlQuery := `SELECT
		ID,ADDRESS,STATUS,CREATE_DATE,DEPARTMENT,
		EMAIL,EMPLOYEE_ID,FIRST_NAME,LAST_NAME,IMAGE_PROFILE,
		PASSWORD,PHONE_NUMBER,SEX,SHIFT,START_DATE,
		USERNAME
		FROM APP_USER
		WHERE EMPLOYEE_ID = ?
	`
	rowsUser, err := db.Query(sqlQuery, empId)
	if err != nil {
		return appUser, err
	}
	for rowsUser.Next() {
		rowsUser.Scan(
			&appUser.ID, &appUser.Address, &appUser.Status, &appUser.CreateDate, &appUser.Department,
			&appUser.Email, &appUser.EmployeeID, &appUser.FirstName, &appUser.LastName, &appUser.ImageProfile,
			&appUser.Password, &appUser.PhoneNumber, &appUser.Sex, &appUser.Shift, &appUser.StartDate,
			&appUser.Username,
		)
	}
	rowsUser.Close()
	sqlQuery = `SELECT ROLE_NAME
		FROM APP_ROLE
		WHERE ID = ?
	`
	rowsRoles, err := db.Query(sqlQuery, appUser.ID)
	if err != nil {
		return appUser, err
	}
	for rowsRoles.Next() {
		var roleName string
		rowsRoles.Scan(&roleName)
		appUser.AppRoles = append(appUser.AppRoles, roleName)
	}
	rowsRoles.Close()
	return appUser, nil
}

func GetById(id int) (entity.AppUser, error) {
	db := getConnection()
	defer db.Close()
	appUser := entity.AppUser{}
	sqlQuery := `SELECT
		ID,ADDRESS,STATUS,CREATE_DATE,DEPARTMENT,
		EMAIL,EMPLOYEE_ID,FIRST_NAME,LAST_NAME,IMAGE_PROFILE,
		PASSWORD,PHONE_NUMBER,SEX,SHIFT,START_DATE,
		USERNAME
		FROM APP_USER
		WHERE ID = ?
	`
	rowsUser, err := db.Query(sqlQuery, id)
	if err != nil {
		return appUser, err
	}
	for rowsUser.Next() {
		rowsUser.Scan(
			&appUser.ID, &appUser.Address, &appUser.Status, &appUser.CreateDate, &appUser.Department,
			&appUser.Email, &appUser.EmployeeID, &appUser.FirstName, &appUser.LastName, &appUser.ImageProfile,
			&appUser.Password, &appUser.PhoneNumber, &appUser.Sex, &appUser.Shift, &appUser.StartDate,
			&appUser.Username,
		)
	}
	rowsUser.Close()
	sqlQuery = `SELECT ROLE_NAME
		FROM APP_ROLE
		WHERE ID = ?
	`
	rowsRoles, err := db.Query(sqlQuery, appUser.ID)
	if err != nil {
		return appUser, err
	}
	for rowsRoles.Next() {
		var roleName string
		rowsRoles.Scan(&roleName)
		appUser.AppRoles = append(appUser.AppRoles, roleName)
	}
	rowsRoles.Close()
	return appUser, nil
}
