package validations

type CreateUserRequest struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	MobileNumber string `json:"mobile_number" validate:"required"`
	Password     string `json:"password" validate:"required,min=6"`
}
