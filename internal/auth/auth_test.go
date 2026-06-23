package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	userId := uuid.New()
	secret := "secret"
	currentTime := time.Hour

	token, err := MakeJWT(userId, secret, currentTime)
	if err != nil {
		t.Fatal("Couldnt create a signed token")
	}

	gotUserId, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatal("Couldnt validate a token")
	}

	if userId != gotUserId {
		t.Errorf("An error occured")
	}
}

func TestBearerToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer abc123")

	got, err := GetBearerToken(headers)
	want := "abc123"

	if err != nil {
		t.Fatal("Couldn`t get a token")
	}

	if got != want {
		t.Errorf("wrong token")
	}
}
