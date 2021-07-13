package main

import (
	"fmt"

	"github.com/pluralsite/webservice/models"
)

func Learn() {
	//VARIABLE DECLARATION
	//verbose
	var n int
	n = 42
	fmt.Println(n)

	//semi verbose
	var f float32 = 3.14
	fmt.Println(f)

	//implicit
	firstName := "Christian"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, i := real(c), imag(c)
	fmt.Println(r, i)

	//POINTERS
	var name *string = new(string)
	*name = "Arthur"
	fmt.Println(*name)

	lastName := "Doe"
	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)

	//CONSTANTS
	const pi = 3.14
	fmt.Println(pi)

	const d = 3
	fmt.Println(d + 3)
	fmt.Println(d + 1.2)

	//ARRAYS
	//verbose
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	//implicit
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//SLICES
	slice := arr[:]
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	slice2 := []int{1, 2, 3}
	slice2 = append(slice2, 4, 42, 27)
	fmt.Println(slice2)

	slice3 := slice2[1:]
	slice4 := slice2[:2]
	slice5 := slice2[1:2] //[included : excluded]
	fmt.Println(slice3, slice4, slice5)

	//MAPS
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)

	//STRUCTS
	//definition
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	//verbose
	var u user //fields are initialized to its "zero" value
	u.ID = 1
	u.FirstName = "Christian"
	u.LastName = "Augustyn"
	fmt.Println(u)

	//implicit
	u2 := user{
		ID:        2,
		FirstName: "John",
		LastName:  "Doe",
	}
	fmt.Println(u2)

	//FOR LOOPS
	counter := 0
	for counter < 5 {
		println(counter)
		counter++
		if counter == 3 {
			continue
		}
		println("continuing...")
	}

	for i := 0; i < 5; i++ {
		println(i)
		if i == 4 {
			break
		}
	}

	s := []int{1, 2, 3}
	for i := 0; i < len(s); i++ {
		println(fmt.Sprint(i) + ":" + fmt.Sprint(s[i]))
	}

	for i, v := range s {
		println(fmt.Sprint(i) + ":" + fmt.Sprint(v))
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts { //can ignore v and just have the first return value. To get the second only you need _, v :=
		println(k, v)
	}

	//IF STATEMENTS
	user1 := models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}

	user2 := models.User{
		ID:        2,
		FirstName: "Sussy",
		LastName:  "Bakka",
	}

	if user1 == user2 {
		println("same user")
	} else if user1.ID == user2.ID {
		println("similar user")
	} else {
		println("different users")
	}

	//SWITCH STATEMENTS
	type HTTPRequest struct {
		Method string
	}

	req := HTTPRequest{Method: "HEAD"}

	switch req.Method {
	case "GET":
		println("Get request")

	case "POST":
		println("Post request")

	case "PUT":
		println("Put request")

	case "DELETE":
		println("Delete request")

	default:
		println("unhandled method")
	}
}
