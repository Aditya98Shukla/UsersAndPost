package entities

type Users struct {
	List []User
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
} 

type Address  struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     `json:"geo"`
}

type Geo struct{
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  `json:"address"`
	Company `json:"company"`
}