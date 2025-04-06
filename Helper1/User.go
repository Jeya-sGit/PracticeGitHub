package helper1

import (
	"fmt"
)

type User struct {
	Name     string
	Mobile   int
	Username string
	Password string
	Userid   int
}

func UserOption1() {
	fmt.Println("1.View Products|2.Order Product|3.OrderHistory")
}

func (user *User) RegisterUser(exslice *[]*User, newUser User) {
	var id int
	fmt.Println(user.Userid)
	for _, i := range *exslice {
		id = i.Userid
	}
	*exslice = append(*exslice, &User{
		Name:     newUser.Name,
		Mobile:   newUser.Mobile,
		Username: newUser.Username,
		Password: newUser.Password,
		Userid:   id + 1,
	})
}
