package poems_lib

import "errors"

const (
	EventError_EmptyName           = "Event name cannot be empty"
	EventError_AlreadyJoined       = "User already joined the event"
	EventError_OrganizerCannotJoin = "Organizer cannot join the event"
)

// CreateEvent creates a public event for anyone to attend with. Needs an Event name as a parameter.
// Makes a user it is called on an organizer of the Event
//
// Returns an error if an empty name is passed
func (u *User) CreateEvent(name string) (*Event, error) {
	if name == "" {
		return nil, errors.New(EventError_EmptyName)
	}

	event := &Event{
		Name:      name,
		Organizer: u,
		Attendees: map[*User]bool{},
	}
	u.Events[event] = true

	return event, nil
}

// Join adds user as an attendee to the event it is called on.
//
// Returns an error if organizer tries to join or the user is already listed in the attendees.
func (e *Event) Join(user *User) error {
	if e.Organizer == user {
		return errors.New(EventError_OrganizerCannotJoin)
	}
	if e.Attendees[user] {
		return errors.New(EventError_AlreadyJoined)
	}
	e.Attendees[user] = true

	return nil
}
