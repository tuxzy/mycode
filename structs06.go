/* RZFeeser | Alta3 Research
   Create a struct named person */

package main

import "fmt"

type nasaMission struct {
    people []astro
    number int
    message string
}

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

    myMission := nasaMission{myslice, 2, "hey you did it"}
    fmt.Println(myMission)
    fmt.Printf("%+v", myMission)
}

