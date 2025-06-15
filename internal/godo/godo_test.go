package godo

import "testing"

func TestAddGodo(t *testing.T) {
	t.Run("should add valid godo item", func(t *testing.T) {
		list := NewList()
		err := list.Add("Water the plants")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got := len(list.Godos); got != 1 {
			t.Fatalf("expected 1 item in list, got %d", got)
		}
		godo := list.Godos[0]
		assertGodoEquals(t, godo, "Water the plants", false)
	})
	t.Run("should reject empty or whitespace-only godo", func(t *testing.T) {
		list := NewList()
		err := list.Add("  \n \t  ")
		if err == nil {
			t.Error("expected error for empty godo, got nil")
		}
	})
}

func TestCompleteGodo(t *testing.T) {
	t.Run("should complete existing godo", func(t *testing.T) {
		list := NewList()
		list.Add("Buy milk")
		err := list.Complete(0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !list.Godos[0].Done {
			t.Error("expected item to be marked as done")
		}
	})
	t.Run("should reject completing non-existent godo", func(t *testing.T) {
		list := NewList()
		list.Add("Buy milk")
		err := list.Complete(-1)
		if err == nil {
			t.Errorf("expected error when completing non-existent godo")
		}
		err = list.Complete(1)
		if err == nil {
			t.Errorf("expected error when completing non-existent godo")
		}
	})
}

func assertGodoEquals(t *testing.T, godo Godo, wantText string, wantDone bool) {
	t.Helper()
	if godo.Text != wantText {
		t.Errorf("godo text = %q, want %q", godo.Text, wantText)
	}
	if godo.Done != wantDone {
		t.Errorf("godo done = %v, want %v", godo.Done, wantDone)
	}
}
