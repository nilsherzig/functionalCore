package main

import (
	"fmt"

	fc "nilsherzig.com/imperativeShell/functionalCore"
)

// this file holds the entire imperative shell

func main() {
	imperativeShell()
}

func imperativeShell() {
	// this is the imperative shell

	// creating an example object
	members := fc.Members{}

	// adding two members
	members = members.AddMember("John", true)
	members = members.AddMember("Doe", false)

	// print current members
	fmt.Println(members)

	// filter map for only ready members
	members = members.ReadyMembers()

	// print filtered members
	fmt.Println(members)
}
