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
	members = fc.AddMember(members, "John", true)
	members = fc.AddMember(members, "Doe", false)

	// print current members
	fmt.Println(members)

	// filter map for only ready members
	members = fc.ReadyMembers(members)

	// print filtered members
	fmt.Println(members)
}
