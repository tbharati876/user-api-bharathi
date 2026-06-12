package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	dob := time.Now().AddDate(-20, 0, 0)

	age := CalculateAge(dob)

	if age != 20 {
		t.Errorf("expected age 20, got %d", age)
	}
}
