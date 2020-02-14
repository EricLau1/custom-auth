package console

import (
	"fmt"
	"log"
)

const (
	FIND_USER   = "3"
	DELETE_USER = "X"
)

func (c *consoleImpl) findUser() {
	field, value := "", ""
	fmt.Printf("Search for which field? ")
	fmt.Scan(&field)
	fmt.Printf("Which %s? ", field)
	fmt.Scan(&value)

	out, err := c.admin.FindUser(field, value)
	if err != nil {
		log.Println(err)
	} else {
		JSON(out)
	}
}

func (c *consoleImpl) deleteUser() {
	username := ""
	fmt.Printf("Username? ")
	fmt.Scan(&username)

	out, err := c.admin.DeleteUser(username)
	if err != nil {
		log.Println(err)
	} else {
		JSON(out)
	}
}
