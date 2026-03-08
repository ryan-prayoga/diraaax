package repositories

import (
	"context"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestMatchesPINWithBcryptHash(t *testing.T) {
	repo := &AuthRepository{}
	hash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate bcrypt hash: %v", err)
	}

	matches, err := repo.matchesPIN(context.Background(), "123456", string(hash))
	if err != nil {
		t.Fatalf("expected bcrypt hash to be checked without error, got: %v", err)
	}
	if !matches {
		t.Fatal("expected correct PIN to match bcrypt hash")
	}
}

func TestMatchesPINWithBcryptHashRejectsInvalidPIN(t *testing.T) {
	repo := &AuthRepository{}
	hash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate bcrypt hash: %v", err)
	}

	matches, err := repo.matchesPIN(context.Background(), "654321", string(hash))
	if err != nil {
		t.Fatalf("expected mismatched bcrypt hash to return no error, got: %v", err)
	}
	if matches {
		t.Fatal("expected incorrect PIN not to match bcrypt hash")
	}
}
