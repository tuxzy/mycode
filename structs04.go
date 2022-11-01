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
    p3 := astro{"George", 40, "moon", true}

    fmt.Println(p1)
    fmt.Println(p2)
    fmt.Println(p3)

}

