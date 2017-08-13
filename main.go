package main

import (
	"fmt"
	"os/user"
)

func main() {
	var response string
	fmt.Println("Which User should I look up?")
	fmt.Scanln(&response)
	fmt.Println(getCurrentUser())
	lookedupBool := lookupUser(response)
	fmt.Println(userExists(lookedupBool, response))
}

func lookupUser(username string) bool {
	lookupname, _ := user.Lookup(username)
	lookedup := lookupname.Username

	if lookedup == username {
		return true
	} else {
		return false
	}
}

func userExists(lookup bool, usr string) string {
	hasUser := lookup
	fmt.Println(hasUser)
	success := "The User " + usr + " exists"
	failure := "The User " + usr + " could not be found"
	if hasUser {
		return success
	} else if !hasUser {
		return failure
	} else {
		return failure
	}
}

func getCurrentUser() string {
	usr, _ := user.Current()
	return usr.Username
}
