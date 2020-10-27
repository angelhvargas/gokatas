package main

import(
    "fmt"
)

func main() {
    fmt.Printf("concatenate" +" strings in \n" + "golang is similar to JS.\n\n")

    // what about literals and operations?
    fmt.Printf("in golang we can operate, e.g. multiply 2 by 2: %d\n\n", 2*2)
    // okay the previous sample is very similar to C... right?
    // let's try now a more go-ish way:
    fmt.Println("multiply 2 by 2: ", 2*2)

    // bool operators also can be printed out:
    fmt.Println("if something is \"true\" and the other thing is \"false\", then: ", true && false)
}
