package persistence

type DatabaseHandler interface {
	AddUser(string, string, int) (int64, error)
	GetUser(int) (User, error)
	DeleteUser(int) error
	GetAllUsers() ([]User, error)
	UpdateUser(int, int, string) error
}
