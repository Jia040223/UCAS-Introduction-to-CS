package main;

import (
    "fmt"
    "strings"
)
func main(){
    var s string;
	fmt.Scanln(&s);
    fmt.Println(strings.Replace(s, " ", "", -1))
}