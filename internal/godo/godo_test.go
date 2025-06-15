package godo

import (
	"os"
	"testing"
)

func TestAddGodo(t *testing.T) {
	list, cleanup := setUpTestDB(t)
	defer cleanup()
	t.Run("should add valid godo item", func(t *testing.T) {
		_, err := list.Add("Water the plants")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if list.String() == "No godos in list" {
			t.Error("expected godo to be added to list")
		}
	})
	t.Run("should reject empty or whitespace-only godo", func(t *testing.T) {
		_, err := list.Add("  \n \t  ")
		if err == nil {
			t.Error("expected error for empty godo, got nil")
		}
	})
}

func TestCompleteGodo(t *testing.T) {
	list, cleanup := setUpTestDB(t)
	defer cleanup()
	t.Run("should complete existing godo", func(t *testing.T) {
		godo, err := list.Add("Buy milk")
		if err != nil {
			t.Fatal(err)
		}

		err = list.Complete(godo.ID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
	t.Run("should reject completing non-existent godo", func(t *testing.T) {
		list.Add("Walk the dog")
		err := list.Complete(999)
		if err == nil {
			t.Errorf("expected error when completing non-existent godo")
		}
	})
}

func setUpTestDB(t *testing.T) (*List, func()) {
	t.Helper()
	f, err := os.CreateTemp("", "godo-test-*.db")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	// Set up the migrations directory path relative to the test file
	origDir := "db/migrations"
	if _, err := os.Stat(origDir); os.IsNotExist(err) {
		// If running tests from package directory, look up one level
		origDir = "../../db/migrations"
	}
	list, err := NewList(f.Name())
	if err != nil {
		os.Remove(f.Name())
		t.Fatal(err)
	}
	return list, func() {
		list.Close()
		os.Remove(f.Name())
	}
}
