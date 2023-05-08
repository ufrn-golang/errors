// Example to demonstrate the creation of custom errors as (global) variables
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: May 8, 2023

package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("denominator is 0")

func divide(numerator, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, ErrDivideByZero
    }
    return (numerator / denominator), (numerator % denominator), nil
}

func main() {
    a := 10
    b := 0
    quotient, remainder, err := divide(a, b)
    if err == nil {
        fmt.Printf("%d / %d = %d (%d)\n", a, b, quotient, remainder)
    } else {
        fmt.Println(err)
    }
}
