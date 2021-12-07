package utils

import (
	"os"
	"strings"
)

func ReadFile(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
