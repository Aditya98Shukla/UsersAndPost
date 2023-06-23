package entities

type Resp struct{
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lat		 string	`json:"lat"`
	Long 	 string	`json:"lng"`
	CompName string	`json:"companyName"`
	Title	 string	`json:"title"`
	Body	 string	`json:"body"`
}