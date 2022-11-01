/* RZFeeser | Alta3 Research
   Create a struct named person */

package main

import "fmt"

type astro struct {
    name string
    age  int
    mission string
    isneeded bool
}

func main() {

    p1 := astro{name: "Bob", age: 20}
    p2 := astro{name: "Fred", age: 30, mission: "mars"}
    p3 := astro{"George", 40, "SpaceX Crew-3", true}

    myslice := []astro{p1, p2, p3}
    fmt.Println(myslice)

    fmt.Println(myslice[2].mission)
}

