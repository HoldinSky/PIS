package poems_lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePublication(t *testing.T) {
	user := NewUser("author1")
	pub, err := user.Publish("A beautiful poem")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	assert.Equal(t, "A beautiful poem", pub.Content)
	assert.Equal(t, user, pub.Author)

	published, keyPresent := user.Publications[pub]

	assert.Len(t, user.Publications, 1)
	assert.True(t, keyPresent)
	assert.True(t, published)
}

func TestCreateEmptyPublication(t *testing.T) {
	user := NewUser("author1")
	_, err := user.Publish("")

	if assert.Error(t, err) {
		assert.EqualError(t, err, PublicationError_EmptyContent)
	}
}

func TestCreateLongPublication(t *testing.T) {
	user := NewUser("author1")
	longContent := make([]byte, 1001) // 1001 символ для перевірки обмеження
	for i := range longContent {
		longContent[i] = 'a'
	}

	_, err := user.Publish(string(longContent))
	if assert.Error(t, err) {
		assert.EqualError(t, err, PublicationError_TooLargeContent)
	}
}
