package model

type Greeting struct {
	ID      int    `json:"id,omitempty" sql:"id"`
	Message string `json:"message" sql:"message"`
}
