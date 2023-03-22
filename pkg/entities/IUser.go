package entities

import "fmt"

type IUser struct {
	Uid       string `json:"uid"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Dob       string `json:"dob"`
	Deleted   bool   `json:"deleted"`
}

func (u *IUser) Eliminar() {
	fmt.Printf("Usuario eliminado - Nombre: %v Apellido %v", u.Firstname, u.Lastname)
}
