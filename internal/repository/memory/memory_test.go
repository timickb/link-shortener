package memory

import "testing"

func TestCreateShortening(t *testing.T) {
	repo := New()

	err := repo.CreateShortening("shortened1", "original1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetOriginal(t *testing.T) {
	repo := New()

	url := "original1"

	err := repo.CreateShortening("shortened1", url)
	if err != nil {
		t.Fatal(err)
	}

	original, err := repo.GetOriginal("shortened1")

	if url != original {
		t.Fatal("expected url = original")
	}
}

func TestGetOriginalNotFound(t *testing.T) {
	repo := New()
	if _, err := repo.GetOriginal("something"); err == nil {
		t.Fatal("expected err")
	}
}
