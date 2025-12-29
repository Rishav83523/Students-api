package types

type Student struct {
	Id    int64
	Name  string `validate:"required,min=2,max=100"` //adding validation tags for name field
	Email string `validate:"required,email"`         //adding validation tags for email field
	Age   int    `validate:"required,min=1,max=120"` //adding validation tags for age field
} //defining the Student struct with fields Id , Name , Email and Age to send in request body