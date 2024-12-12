package poems_lib

import "errors"

const (
	PublicationError_EmptyContent    = "Content cannot be empty"
	PublicationError_TooLargeContent = "Content exceeds maximum length of 1000 characters"
)

// Publish makes user to publish a post. Takes content as a parameter.
//
// Returns an error if content is empty or exceeds in total of 1000 characters.
func (u *User) Publish(content string) (*Publication, error) {
	if content == "" {
		return nil, errors.New(PublicationError_EmptyContent)
	}

	if len(content) > 1000 {
		return nil, errors.New(PublicationError_TooLargeContent)
	}

	pub := &Publication{
		Author:  u,
		Content: content,
	}
	u.Publications[pub] = true

	return pub, nil
}
