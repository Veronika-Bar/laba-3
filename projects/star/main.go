package main

import (
 "fmt"
 "strings"
)

func main() {
 var input string
 fmt.Scan(&input)

 // Добавляем '*' между буквами
 result := strings.Join(strings.Split(input, ""), "*")

 fmt.Println(result)
}