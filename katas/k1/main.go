package main

import (
    "fmt"
)

func main() {
    // variable declaration and some types supported
    var var_1 string
    var variable_1 string = "test"
    var uint_var = 1234
    var value int = 12345
    var boolean bool = true
    var float float32 =2.3743
    var float_64 float64 = 1241234123412.2134907812348971263498761234
    var complex_ complex128 = -0.6i + 0.127

    // short variable declaration is supported too
    var_2 := "test"
    account_number := 12356456
    netflix_series := "The Witcher"
    // fmt package implements formatted I/O functions analogous 
    //   to C's printf or scanf functions. see https://golang.org/pkg/fmt/
    fmt.Println(var_1, variable_1)
    fmt.Println(uint_var)
    fmt.Println(var_2)
    fmt.Println(account_number)
    fmt.Println(value)
    fmt.Println(boolean)
    fmt.Println(float)
    fmt.Println(float_64)
    fmt.Println(complex_)
    fmt.Println(netflix_series)
    fmt.Println(boolean)

    fmt.Printf("\nCasted variables:\n\n")

    var amount int = 80
    var amount2 float64 = float64 (amount)
    var amount3 uint = uint (amount2)

    fmt.Println(amount3)
}
