package fp

import (
	"fmt"
	"runtime"
	"time"
)

func Fp(a ...any) string {

	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("\n time: %s file: %s, line: %d ", time.Now().String(), file, line)
		for i := range a {
			fmt.Print(a[i], "             ")

		}

		return fmt.Sprintf("\nfile: %s, line: %d ", file, line)
	}
	for i := range a {
		fmt.Print(a[i], "             ")

	}
	fmt.Println("")
	return ""
}