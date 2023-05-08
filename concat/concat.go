// Example to demonstrate the creation of custom errors via variables
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: May 8, 2023

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func concat(words ...string) (string, error) {
    if len(words) <= 1 {
        return "", errors.New("no words enough have been provided")
    }
    return strings.Join(words, " "), nil
}

func main() {
    args := os.Args[1:]
    if result, err := concat(args...); err != nil {
        fmt.Printf("Error: %s\n", err)
    } else {
        fmt.Printf("Concatenated string: '%s'\n", result)
    }
}
