// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// Determine whether a given year is a leap year
func IsLeapYear(year int) bool {    
    if year % 4 == 0 {
        if year % 100 == 0 {
            return year % 400 ==  0
        }
        return true
    }
    return false
}