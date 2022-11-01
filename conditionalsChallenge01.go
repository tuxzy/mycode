package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {

  rand.Seed(time.Now().UnixNano())

  mySlice := []string{"s0", "s1", "s2", "s3", "s4", "s5"}

  index := rand.Intn(5 - 0)
  selection := mySlice[index]

  if selection == "s0" {
      fmt.Println("A")
  } else if selection == "s1" {
      fmt.Println("B")
  } else if selection == "s2" {
      fmt.Println("C")
  } else if selection == "s3" {
      fmt.Println("D")
  } else if selection == "s4" {
      fmt.Println("E")
  } else {
      fmt.Println(selection)
  }

}
