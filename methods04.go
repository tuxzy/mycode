/* Alta3 Research | RZFeeser
   Methods - defining         */

// Go program to illustrate the
// method with struct type receiver
package main
 
import "fmt"
 
type Virtmach struct {
    ip      string
    hostname    string
    diskgb     int
    ram    int
}
 
// Method with a receiver
// of author type
func (v Virtmach) show() {
 
    fmt.Println("ip: ", v.ip)
    fmt.Println("hostname: ", v.hostname)
    fmt.Println("diskgb: ", v.diskgb)
    fmt.Println("ram: ", v.ram)
}

func (a *Virtmach) doubleDiskGb() {
    a.diskgb = a.diskgb * 2;
}

func (a *Virtmach) changeHostname(newHostname string) {
    a.hostname = newHostname
}
 
// Main function
func main() {
 
    // Initializing the values
    // of the author structure
    res := Virtmach{
        ip:      "1.192.0.1",
        hostname:    "zachary-hertig.com",
        diskgb:     128,
        ram:        4,
    }
 
    // Calling the method
    res.show()
    res.doubleDiskGb()
    res.changeHostname("zack.com")
    res.show()

}

