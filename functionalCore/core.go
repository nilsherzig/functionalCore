package functionalcore

// this package holds the functional core / "business logic" of the application

func (ms Members) New(initMembers []Member) (result Members) {
	for _, member := range initMembers {
		result.Items = append(result.Items, member)
	}
	return result
}

func (members Members) ReadyMembers() (readyMembers Members) {
	for _, member := range members.Items {
		if member.IsReady {
			readyMembers.Items = append(readyMembers.Items, member)
		}
	}
	return readyMembers
}

func (members Members) AddMember(name string, isReady bool) Members {
	members.Items = append(members.Items, Member{Name: name, IsReady: isReady})
	return members
}
