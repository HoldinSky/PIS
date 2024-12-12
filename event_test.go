package poems_lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEvent(t *testing.T) {
	user := NewUser("organizer1")
	event, err := user.CreateEvent("Poetry Night")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	assert.Equal(t, "Poetry Night", event.Name)
	assert.Equal(t, user, event.Organizer)

	isOnTheList, keyPresent := user.Events[event]

	assert.Len(t, user.Events, 1)
	assert.True(t, keyPresent)
	assert.True(t, isOnTheList)
}

func TestJoinEvent(t *testing.T) {
	organizer := NewUser("organizer1")
	event, _ := organizer.CreateEvent("Poetry Night")
	attendee := NewUser("attendee1")

	err := event.Join(attendee)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	attends, keyPresent := event.Attendees[attendee]

	assert.Len(t, event.Attendees, 1)
	assert.True(t, keyPresent)
	assert.True(t, attends)
}

func TestJoinEventTwice(t *testing.T) {
	organizer := NewUser("organizer1")
	event, _ := organizer.CreateEvent("Poetry Night")
	attendee := NewUser("attendee1")

	_ = event.Join(attendee)
	err := event.Join(attendee)
	if assert.Error(t, err) {
		assert.EqualError(t, err, EventError_AlreadyJoined)
	}
}
