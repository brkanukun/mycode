package main

import (
    "regexp"
    "fmt"
)

func main() {
    var re = regexp.MustCompile(`(?m)fmt.Println("You are using the latest version of GoLang")`)
    var str = `
`
    
    for i, match := range re.FindAllString(str, -1) {
        fmt.Println(match, "found at index", i)
    }
}

