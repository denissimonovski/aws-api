package persistence

type (
	// User represents the structure of our resource
	User struct {
		Name   string `json:"name"`
		Gender string `json:"gender"`
		Age    int    `json:"age"`
		Id     int    `json:"id"`
	}
)
