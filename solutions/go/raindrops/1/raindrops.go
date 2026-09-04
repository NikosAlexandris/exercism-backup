package raindrops
import (
    "strconv"
)

// fmt.Println("Debug message")
var result string

func Convert(number int) string {
    output := ""
    if number % 3 == 0 {
        output += "Pling"
    }
    if number % 5 == 0 {
        output += "Plang"
    } 
    if number % 7 == 0 {
        output += "Plong"
    } 
    if output == "" {
        output = strconv.Itoa(number)
    }
    return output
}