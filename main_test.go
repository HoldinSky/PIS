package poems_lib

func ExampleNewUser() {
	NewUser("Nazar")
}

func ExampleUser_CreateEvent() {
	user := NewUser("Nazar")
	event, _ := user.CreateEvent("Poetry Night")
	println(event.Organizer)
	// Output:
	// User "Nazar"
}

func ExampleUser_CreateEvent2() {
	userNazar := NewUser("Nazar")
	userAndrew := NewUser("Andrew")
	event, _ := userNazar.CreateEvent("Poetry Night")
	event.Join(userAndrew)
}
