package main

import (
	"errors"
	"fmt"
)

type DivisionError struct { 
   num, denom int
   message string
}

func (e *DivisionError) Error() string { 
   return e.message
}

func divide(a, b int) (int, error) {
   if b == 0 { 
		return 0, &DivisionError{ 
			message: fmt.Sprintf("cannot divide '%d' by zero", a), 
        	num: a, denom: b, 
    	}
   } 
   return (a / b), nil 
}

func main() {
    a, b := 10, 0
    result, err := divide(a, b)
    if err != nil {
        var divErr *DivisionError
        switch {
        case errors.As(err, &divErr):
            fmt.Printf("%d / %d is not mathematically valid: %s\n",
              divErr.num, divErr.denom, divErr.Error())
        default:
            fmt.Printf("unexpected division error: %s\n", err)
        }
        return
    }

    fmt.Printf("%d / %d = %d\n", a, b, result)
}
