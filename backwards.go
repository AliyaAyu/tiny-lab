package main

import "fmt"

func main() {
    word := "океан"     
    runes := []rune(word)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    fmt.Println("Было:", word)
    fmt.Println("Стало:", string(runes))
}
