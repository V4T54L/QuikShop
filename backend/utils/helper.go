package utils

import "strconv"

func GetInt(s string) (*int, error) {
	if len(s) == 0 {
		return nil, nil
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	return &val, nil
}
