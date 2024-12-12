// Package poems_lib. Take your art to the world!
//
// PoemsLib is a project of a social media targeted for poetry lovers.
// Publish, share and read poems from wherever you are whenever you want.
package poems_lib

type Publication struct {
	Author  *User
	Content string
}

// Event represents a public event that can be created by one User - organizer and attended by anyone who wants.
type Event struct {
	Name      string
	Organizer *User
	Attendees map[*User]bool
}

// User represents an entity of an app user who:
//
// May be subscribed to a number of other users (Subscriptions)
//
// May have posts (Publications)
//
// May have organized events (Events)
type User struct {
	Name          string
	Subscriptions map[*User]bool
	Publications  map[*Publication]bool
	Events        map[*Event]bool
}

// String makes a representation of user in format: "User <Name>"
func (u User) String() string {
	return "User <" + u.Name + ">"
}

// NewUser creates an instance of a User by his name
func NewUser(name string) *User {
	return &User{
		Name:          name,
		Subscriptions: map[*User]bool{},
		Publications:  map[*Publication]bool{},
		Events:        map[*Event]bool{},
	}
}
