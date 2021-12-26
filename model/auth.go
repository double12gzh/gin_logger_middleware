package model

type Auth struct {
	Tel     string   `json:"tel"`
	Addr    string   `json:"addr"`
	Age     int      `json:"age"`
	Persons []Person `json:"persons"`
}

type Person struct {
	// https://github.com/gin-gonic/examples/tree/master/custom-validation
	Name    string `json:"name" binding:"required,visitorNameValidation"`
	Sex     string `json:"sex"`
	Company string `json:"company"`
}

type Visitor struct {
	Email string `json:"email"`
	Person
}
