package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
    tests := []struct {
        input    int
        expected bool
    }{
        {2, true},
        {3, true},
        {4, false},
        {5, true},
        {6, false},
    }

    for _, test := range tests {
        result, err := isPrime(test.input)
        if err != nil {
            t.Errorf("Error: %v", err)
            continue
        }
        if result != test.expected {
            t.Errorf("isPrime(%d) = %v; expected %v", test.input, result, test.expected)
        }
    }
}

func TestIsPrime2(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected bool
        wantErr  bool
    }{
        {"negative number", -1, false, true},
        {"zero", 0, false, true},
        {"one", 1, false, true},
        {"smallest prime", 2, true, false},
        {"even non-prime", 4, false, false},
        {"large prime", 997, true, false},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result, err := isPrime(test.input)
            
            if (err != nil) != test.wantErr {
                t.Errorf("isPrime(%d) unexpected error: %v", test.input, err)
                return
            }
            
            if !test.wantErr && result != test.expected {
                t.Errorf("isPrime(%d) = %v; expected %v", test.input, result, test.expected)
            }
        })
    }
}