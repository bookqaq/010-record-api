package utils

import (
	"fmt"
	"os"
)

// a simulated panic is a throw-error-and-show wrapper, should only use before logger initilization complete.
func SimulatedPanic(msg any) {
	fmt.Println(msg)
	fmt.Scanln()
	os.Exit(1)
}
