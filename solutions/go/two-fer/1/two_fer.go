package twofer // In some English accents "two fer" is a two-for-one offer : buy one, get one for free

import "fmt"

func ShareWith(name string) string {
    // Share free cookie with someone
    if (name == "") {
        name = "you"
    }
    return fmt.Sprintf("One for %s, one for me.", name)
}