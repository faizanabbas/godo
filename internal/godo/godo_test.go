package godo

import (
	"os"
	"strings"
	"testing"
)

func TestGodoOperations(t *testing.T) {
	tests := []struct {
		name   string
		testFn func(*testing.T, *List)
	}{
		{
			name: "adding valid godo",
			testFn: func(t *testing.T, l *List) {
				godo, err := l.Add("Water the plants")
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if godo.Text != "Water the plants" {
					t.Errorf("expected text 'Water the plants', got %q", godo.Text)
				}
				if godo.Done {
					t.Error("new godo should not be marked as done")
				}
			},
		},
		{
			name: "rejecting empty godo",
			testFn: func(t *testing.T, l *List) {
				_, err := l.Add("  \n \t  ")
				if err == nil {
					t.Error("expected error for empty godo, got nil")
				}
				if !strings.Contains(err.Error(), "empty") {
					t.Errorf("expected 'empty' error message, got %q", err.Error())
				}
			},
		},
		{
			name: "completing existing godo",
			testFn: func(t *testing.T, l *List) {
				godo, err := l.Add("Buy milk")
				if err != nil {
					t.Fatal(err)
				}
				err = l.Complete(godo.ID)
				if err != nil {
					t.Errorf("unexpected error completing godo: %v", err)
				}
				output := l.String()
				if !strings.Contains(output, "[✓] Buy milk") {
					t.Errorf("expected completed godo to be marked with ✓, got: %s", output)
				}
			},
		},
		{
			name: "rejecting completion of non-existent godo",
			testFn: func(t *testing.T, l *List) {
				err := l.Complete(999)
				if err == nil {
					t.Error("expected error when completing non-existent godo")
				}
				if !strings.Contains(err.Error(), "not found") {
					t.Errorf("expected 'not found' error message, got %q", err.Error())
				}
			},
		},
		{
			name: "string representation shows correct format",
			testFn: func(t *testing.T, l *List) {
				if output := l.String(); output != "No godos in list" {
					t.Errorf("expected empty list message, got %q", output)
				}
				godo, _ := l.Add("Task 1")
				l.Complete(godo.ID)

				output := l.String()
				if !strings.Contains(output, "Godo list:") {
					t.Error("expected output to start with 'Godo list:'")
				}
				if !strings.Contains(output, "[✓] Task 1") {
					t.Error("expected completed task to be marked with ✓")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, cleanup := setUpTestDB(t)
			defer cleanup()
			tt.testFn(t, list)
		})
	}
}

func setUpTestDB(t *testing.T) (*List, func()) {
	t.Helper()
	f, err := os.CreateTemp("", "godo-test-*.db")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	list, err := NewList(f.Name())
	if err != nil {
		os.Remove(f.Name())
		t.Fatal(err)
	}
	cleanup := func() {
		list.Close()
		os.Remove(f.Name())
	}
	return list, cleanup
}
