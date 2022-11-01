/* Alta3 Research | RZFeeser
   Map - hosts to IP:port resolution  */

package main

import (  
    "fmt"
)

func main() {  
    
    // create a map type, "hostResolution"
    extensions := map[string]string{
        "Python": ".py",
        "Java": ".java",
        "C++": ".cpp",
    }
    
    fmt.Println(extensions)
    delete(extensions, "C++")
    extensions["Julia"] = ".jl"
    fmt.Println(extensions)

}

