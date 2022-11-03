/* Alta3 Research | RZFeeser
   Custom Sorting - Sorting structs with custom functions   */

package main

import (
    "fmt"
    "sort"
)

// record a "Person" Name, Age
type Person struct {
    Name   string
    Age    int
    Height int
}


/* In Go, if you define "String()" for a struct
   the result is a custom "ToString". That is to say,
   we can control the "view" that is returned when someone tries
   to print our struct */
//func (p Person) String() string {
//    return fmt.Sprintf("%s: %d", p.Name, p.Age)
//}

// ByAge implements sort.Interface for []Person based on
// the Age field
type ByAge []Person

func (a ByAge) Len() int {
    return len(a)
}
func (a ByAge) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
    return a[i].Height < a[j].Height
}

func main() {

    // Name, Age
    people := []Person{
        {"Bob", 31, 55},
        {"John", 42, 54},
        {"Michael", 17, 56},
        {"Jenny", 26, 53},
    }

    /* If you commented out the "String()" function,
    the way the people struct is displayed will change.
    The formatting it is following is provided by the "String()"
    function */
    fmt.Println(people)

    // We defined the interface for sort.Sort
    // a set of methods for the slice type, as with ByAge, and
    // call sort.Sort.
    sort.Sort(ByAge(people))
    fmt.Println(people)

}

