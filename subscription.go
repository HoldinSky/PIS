package poems_lib

import "errors"

const (
	SubscriptionError_AlreadySubscribed = "Already subscribed to this author"
	SubscriptionError_SubscribeToSelf   = "Cannot subscribe to self"
)

// Subscribe makes a user to follow the other user.
// User to be followed is passed as a parameter.
//
// Returns an error if user tries to subscribe to himself or if he is already subscribed to a chosen author.
func (u *User) Subscribe(author *User) error {
	if u == author {
		return errors.New(SubscriptionError_SubscribeToSelf)
	}

	if u.Subscriptions[author] {
		return errors.New(SubscriptionError_AlreadySubscribed)
	}

	u.Subscriptions[author] = true
	return nil
}
