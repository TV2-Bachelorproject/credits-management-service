package user

import (
	"testing"

	"github.com/TV2-Bachelorproject/server/pkg/db"
)

func TestUserVerification(t *testing.T) {
	db.Migrate(User{})
	defer db.Reset(User{})

	user, err := New("John Doe", "john@example.com", "123456", Admin)

	if err != nil {
		t.Error(err)
	}

	db.Create(&user)

	// invalid email
	_, err = Authorize("john@example", "123456")

	if err == nil {
		t.Errorf("authorized user with an invalid email")
	}

	// test invalid password
	_, err = Authorize("john@example.com", "")

	if err == nil {
		t.Errorf("authorized user with an invalid password")
	}

	_, err = Authorize("john@example.com", "123456")

	if err != nil {
		t.Error("credentials are valid and should be granted access")
	}

}
