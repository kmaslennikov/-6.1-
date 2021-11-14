package main

import "fmt"

func main() {
  fmt.Println("Введите число:")
  var count int
  fmt.Scan(&count)

  for i:=0; i<=count; i++{
    fmt.Println(i)
  }
}