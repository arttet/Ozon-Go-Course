package test

import "strconv"

func CalcSum(x, y int) int {
	return x + y
}

func CalcSumFromStrings(x, y string) (int, error) {
	a, err := strconv.Atoi(x)
	if err != nil {
		return 0, err
	}

	b, err := strconv.Atoi(y)
	if err != nil {
		return 0, err
	}

	return a + b, nil
}
