package request

type SignUp struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=25"`
	LastName  string `json:"last_name" binding:"required,min=1,max=25"`
	Phone     string `json:"phone" binding:"required,min=10,max=10"`
	Password  string `json:"password" binding:"required,min=6,max=25"`
}
