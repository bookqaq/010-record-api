package utils

import "fmt"

// a wrapper to concat go1.22 new url form "method /endpoint" beautifully
func RequestURL(method string, endpoint string) string {
	return fmt.Sprintf("%s %s", method, endpoint)
}
