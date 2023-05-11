package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var err error
    hamming_distance := 0
    if len(a) != len(b) {
        err = errors.New("Strigns are not of equal length!")
        return 0, err
    }
	for i, _ := range a{
        if a[i] != b[i] {
            hamming_distance += 1
        }
    }
	return hamming_distance, err
}
