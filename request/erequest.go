// Example to demonstrate handling multiple errors
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: May 8, 2023

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrTimeout = errors.New("the request timed out")
	ErrRejected = errors.New("the request was rejected")
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func sendRequest(req string) (string, error) {
    switch random.Intn(3) {
    case 0:
        return "Success", nil
    case 1:
        return "", ErrRejected
    default:
        return "", ErrTimeout
    }
}

func main() {
    response, err := sendRequest("Hello")
    for err == ErrTimeout {
        fmt.Println("Timeout. Retrying.")
        response, err = sendRequest("Hello")
    }

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println(response)
    }
}