package main

import "fmt"

func main() {
    number := 54   
    isSimple := true

    for i := 2; i*i <= number; i++ {
        if number%i == 0 {
            isSimple = false
            break
        }
    }

    if isSimple && number > 1 {
        fmt.Println(number, "— простое")
    } else {
        fmt.Println(number, "— не простое")
    }
}
