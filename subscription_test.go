package poems_lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubscribeToAuthor(t *testing.T) {
	user := NewUser("reader1")
	author := NewUser("author1")

	err := user.Subscribe(author)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	isSubscribedTo, keyPresent := user.Subscriptions[author]

	assert.Len(t, user.Subscriptions, 1)
	assert.True(t, keyPresent)
	assert.True(t, isSubscribedTo)
}

func TestSubscribeToSameAuthorTwice(t *testing.T) {
	user := NewUser("reader1")
	author := NewUser("author1")

	_ = user.Subscribe(author)
	err := user.Subscribe(author)
	if assert.Error(t, err) {
		assert.EqualError(t, err, SubscriptionError_AlreadySubscribed)
	}
}

func TestSubscribeToSelf(t *testing.T) {
	user := NewUser("reader1")

	err := user.Subscribe(user)
	if assert.Error(t, err) {
		assert.EqualError(t, err, SubscriptionError_SubscribeToSelf)
	}
}
