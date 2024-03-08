package functionalcore

// this package holds the functional core / "business logic" of the application

// New erstellt eine neue Members-Instanz basierend auf initMembers.
// Diese Funktion ist funktional, da sie keine externen Zustände verändert.
func New(initMembers []Member) (result Members) {
	for _, member := range initMembers {
		result.Items = append(result.Items, member)
	}
	return result
}

// ReadyMembers filtert die Mitglieder basierend auf ihrem Ready-Status.
// Diese Funktion ist funktional, da sie keine externen Zustände verändert.
func ReadyMembers(members Members) (readyMembers Members) {
	for _, member := range members.Items {
		if member.IsReady {
			readyMembers.Items = append(readyMembers.Items, member)
		}
	}
	return readyMembers
}

// AddMember fügt ein neues Mitglied hinzu und gibt eine neue Members-Instanz zurück.
// Diese Funktion ist funktional, da sie eine neue Instanz zurückgibt statt den ursprünglichen Zustand zu ändern.
func AddMember(members Members, name string, isReady bool) Members {
	newMembers := Members{
		Items: append([]Member(nil), members.Items...), // Kopie der ursprünglichen Items-Liste erstellen
	}
	newMember := Member{Name: name, IsReady: isReady}
	newMembers.Items = append(newMembers.Items, newMember)
	return newMembers
}
