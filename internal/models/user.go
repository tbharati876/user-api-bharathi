package models

type UserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	DOB  string `json:"dob" validate:"required"`
}

type CreateUserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}