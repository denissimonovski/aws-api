package persistence

type (
	// User e novo definiran resurs
	User struct {
		Name   string `json:"name"`
		Gender string `json:"gender"`
		Age    int    `json:"age"`
		Id     int    `json:"id"`
	}
)
