package gx

import (
	"errors"
	"testing"
)

func TestMust_Success(t *testing.T) {
	// Test with int
	result := Must(42, nil)
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}

	// Test with string
	str := Must("hello", nil)
	if str != "hello" {
		t.Errorf("Expected 'hello', got '%s'", str)
	}

	// Test with struct
	type testStruct struct {
		Value int
	}
	s := Must(testStruct{Value: 10}, nil)
	if s.Value != 10 {
		t.Errorf("Expected Value=10, got %d", s.Value)
	}
}

func TestMust_Panic(t *testing.T) {
	testErr := errors.New("test error")

	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected panic, but didn't panic")
		}
		if r != testErr {
			t.Errorf("Expected panic with testErr, got %v", r)
		}
	}()

	// This should panic
	Must(42, testErr)
}

func TestMust_PanicWithDifferentTypes(t *testing.T) {
	testErr := errors.New("another error")

	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected panic, but didn't panic")
		}
		if r != testErr {
			t.Errorf("Expected panic with testErr, got %v", r)
		}
	}()

	// This should panic with string type
	Must("value", testErr)
}
