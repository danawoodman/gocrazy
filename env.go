package gocrazy

import "os"

/*
Getenv returns the value of the environment variable named by the key or the
fallback value if the variable is not set.
*/
func Getenv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
