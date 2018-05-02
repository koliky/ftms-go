package entity

type AppUser struct {
	ID           int
	Address      string
	Status       string
	CreateDate   string
	Department   string
	Email        string
	EmployeeID   string
	FirstName    string
	LastName     string
	ImageProfile string
	Password     string
	PhoneNumber  string
	Sex          string
	Shift        string
	StartDate    string
	Username     string
	AppRoles     []string
}
