package hamming

import "errors"

func Distance(a, b string) (int, error) {
    
    if len(a) != len(b) {
        return 0, errors.New("Strings differ in length")
    }

    var hamming_distance int
    for index := range a {
        if a[index] != b[index] {
            hamming_distance ++
        }
    }
    
    return hamming_distance, nil
}