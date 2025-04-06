package main

import (
	helper1 "UIMProject/Helper1"
	"fmt"
)

type Product struct {
	productid    int
	productname  string
	productprice float64
	availcount   int
	availstatus  string
}
type Order struct {
	orderid   int
	buyerid   int
	buyername string
	pid       int
	pname     string
	quantity  float64
	amount    float64
}

// func UsersOption() {
// 	fmt.Println("1.View Products|2.Order Product|3.OrderHistory")
// }

func main() {

	// var p []*helper1.Product

	// p = []*helper1.Product{&helper1.Product{}}

	uslice := []*helper1.User{
		&helper1.User{
			Name: "Jeya", Mobile: 8765432, Username: "jeya321", Password: "jeya$321", Userid: 1,
		},
	}

	products := []*Product{
		&Product{productid: 1, productname: "HPEliteBook", productprice: 70000, availcount: 2, availstatus: "Available"},
		&Product{productid: 2, productname: "HPProBook", productprice: 62000, availcount: 1, availstatus: "Available"},
	}

	orders := []*Order{}

	var user1 helper1.User
	fmt.Println("Welcome to UIM")
	fmt.Println("1.Login|2.Register")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		fmt.Println("Welcome")
		fmt.Println("Enter the Username:")
		var uname string
		fmt.Scanln(&uname)
		fmt.Println("Enter the Password:")
		var pass string
		fmt.Scanln(&pass)

		for _, u := range uslice {
			if uname == u.Username && pass == u.Password {
				user1 = *u
				fmt.Printf("Welcome %v\n", user1.Name)
				helper1.UserOption1()
				var opt1 int
				fmt.Scanln(&opt1)
				switch opt1 {
				case 1:
					for _, p := range products {
						fmt.Printf("ProductID:%v|ProductName:%v|ProductPrice:%2.f|AvailabilityCount:%d|AvailabilityStatus:%v\n", p.productid, p.productname, p.productprice, p.availcount, p.availstatus)
					}
					fmt.Println("1.Wish to order?")
					var o int
					fmt.Scanln(&o)
					if o == 1 {
						fmt.Println("Enter the ProductID:")
						var productid int
						fmt.Scanln(&productid)
						var pName string
						var pPrice float64
						for _, p := range products {
							if productid == p.productid {
								pName = p.productname
								pPrice = p.productprice
							}
						}
						var oid int
						for _, Or := range orders {
							oid = Or.orderid
						}
						fmt.Println("Quantity:")
						var count float64
						fmt.Scanln(&count)

						orders = append(orders, &Order{orderid: oid + 1, buyerid: u.Userid, buyername: user1.Name, pid: productid, pname: pName, quantity: count, amount: pPrice * count})
						fmt.Println("Your Order Summary")
						var orderid int
						var piid int
						var pnname string
						var qu float64
						var total float64
						for _, fo := range orders {
							if fo.buyerid == user1.Userid {
								orderid = fo.orderid
								piid = fo.pid
								pnname = fo.pname
								qu = fo.quantity
								total = fo.amount
							}
						}
						for _, p1 := range products {
							if piid == p1.productid {
								p1.availcount = p1.availcount - int(qu)
							}
							if piid == p1.productid && p1.availcount < 1 {
								p1.availstatus = "Not Available"
							}
						}

						fmt.Printf("OrderId:%v\nProductId:%v\nProductName:%v\nQuantity:%f\nTotal:%f", orderid, piid, pnname, qu, total)
						for _, p2 := range products {
							fmt.Printf("\nProductID:%v|ProductName:%v|ProductPrice:%2.f|AvailabilityCount:%d|AvailabilityStatus:%v\n", p2.productid, p2.productname, p2.productprice, p2.availcount, p2.availstatus)
						}
					}
				case 2:
					fmt.Println("Order")
				case 3:
					fmt.Println("History")
				}
			}
		}
	case 2:
		fmt.Println("Enter the Following Details")
		fmt.Println("Enter Your Name:")
		var name string
		fmt.Scanln(&name)
		fmt.Println("Enter Your MobileNumber:")
		var mn int
		fmt.Scanln(&mn)
		fmt.Println("Create a UserName:")
		var un string
		fmt.Scanln(&un)
		fmt.Println("Create a Password:")
		var pass string
		fmt.Scanln(&pass)
		newUser := helper1.User{Name: name, Mobile: mn, Username: un, Password: pass}
		usernew := &helper1.User{}
		usernew.RegisterUser(&uslice, newUser)
		for _, ii := range uslice {
			fmt.Println(ii.Userid, "|", ii.Name)
		}
	}
}
