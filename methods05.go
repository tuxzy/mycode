/* Alta3 Research | RZFeeser
   Methods - defining         */

// Go program to illustrate the
// method with struct type receiver
package main
 
import "fmt"
 
type Player struct {
    inventory []string
    lives     int
    stage     int
}
 
// Method with a receiver
// of author type
func (v Player) show() {
 
    fmt.Println("lives: ", v.lives)
    fmt.Println("stage: ", v.stage)
    fmt.Println("inventory: ", v.inventory)
}

func (a *Player) Greenmushroom() {
    a.lives++;
}

func (a *Player) Pickup(item string) {
    a.inventory = append(a.inventory, item)
}

func (a *Player) CanWhistle() bool {
    if a.stage >= 5 {
      return true;
    } else {
      return false;
    }
}
 
// Main function
func main() {
 
    // Initializing the values
    // of the author structure
    res := Player{
        inventory:      []string{"mushroom"},
        lives:    3,
        stage:    5,
    }
 
    // Calling the method
    res.show()
    res.Greenmushroom()
    res.Pickup("hammer")
    fmt.Println(res.CanWhistle())
    res.show();

}

