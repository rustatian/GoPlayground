package main

import (
	"github.com/casbin/casbin"
	"net/http"
)

type User struct {
	ID   int
	Name string
	Role string
}

type Users []User

func main() {

	a := http.Request{}
	e := casbin.NewEnforcer("./accessControl/auth_model.conf", "./accessControl/policy.csv")

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	if e.Enforce(sub, obj, act) == true {
		// permit alice to read data1
	} else {
		// deny the request, show an error
	}

	roles := e.GetRolesForUser("alice")
	println(roles)
}
