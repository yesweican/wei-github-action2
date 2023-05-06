package main

import "testing"

func TestGreeting(t *testing.T) {
    result := Greeting("Bob")
    expected := "Hello, Bob!"

    if result != expected {
        t.Errorf("Greeting(\"Bob\") returned %q, expected %q", result, expected)
    }
}

func TestAppendElement(t *testing.T) {
    slice := []int{1, 2, 3}
    expected := []int{1, 2, 3, 4}

    result := AppendElement(slice, 4)

    if len(result) != len(expected) {
        t.Errorf("AppendElement(%v, 4) returned slice with length %d, expected length %d", slice, len(result), len(expected))
    }

    for i := range expected {
        if expected[i] != result[i] {
            t.Errorf("AppendElement(%v, 4) returned %v, expected %v", slice, result, expected)
            break
        }
    }
}


