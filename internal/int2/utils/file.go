package utils

import "os"

func ReadFile(path string) (string, error) {
	raw, err := os.ReadFile(path)
	return string(raw), err
}