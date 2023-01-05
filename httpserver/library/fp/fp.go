package fp

import (
	"fmt"
	"runtime"
)

func Fp(a ...any) string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("\nfile: %s, line: %d ", file, line)
		for i := range a {
			fmt.Print(a[i], "\t")

		}

		return fmt.Sprintf("\nfile: %s, line: %d ", file, line)
	}
	for i := range a {
		fmt.Print(a[i], "             ")

	}

	return ""
}
