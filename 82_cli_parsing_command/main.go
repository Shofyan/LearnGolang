package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var app = kingpin.New("app", "simple app")
var commandAdd = app.Command("add", "add new user")
var commandAddFlagOverride = commandAdd.Flag("override", "override existing user").Short('o').Bool()
var commandAddArgsUser = commandAdd.Arg("user", "username").Required().String()

var commandUpdate = app.Command("update", "update user")
var commandUpdateArgOldUser = commandUpdate.Arg("old", "old username").Required().String()
var commandUpdateArgNewUser = commandUpdate.Arg("new", "new username").Required().String()

var commandDelete = app.Command("delete", "delete user")
var commandDeleteFlagForce = commandDelete.Flag("force", "force delete").Short('f').Bool()
var commandDeleteArgsUser = commandDelete.Arg("user", "username").Required().String()

func main() {
	commandAdd.Action(func(context *kingpin.ParseContext) error {
		user := *commandAddArgsUser
		override := *commandAddFlagOverride
		fmt.Printf("adding user %s, override %t \n", user, override)

		return nil
	})

	commandUpdate.Action(func(context *kingpin.ParseContext) error {
		oldUser := *commandUpdateArgOldUser
		newUser := *commandUpdateArgNewUser
		fmt.Printf("updating user from %s %s \n", oldUser, newUser)

		return nil
	})

	commandDelete.Action(func(context *kingpin.ParseContext) error {
		user := *commandDeleteArgsUser
		force := *commandDeleteFlagForce
		fmt.Printf("deleting user %s, force %t \n", user, force)

		return nil
	})

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
