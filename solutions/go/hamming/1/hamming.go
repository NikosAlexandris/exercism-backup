package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
    _error := errors.New("")
    
    if len(a) != len(b) {
        return 0, errors.New("true")
    }
    if a == b {
        return 0, nil
    }

    var hamming_distance int
    for index := range a {
        if a[index] != b[index] {
            hamming_distance += 1
            _error = nil
        }
    }
    
    return hamming_distance, _error
}