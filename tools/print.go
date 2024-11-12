package tools

import(
	"fmt"
	"time"
)
func TypePrint(speed int, text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}
