package tools

import(
	"fmt"
	"time"
)
func TypePrint(speed int, text string) {
	if speed<0 {
		speed = 0
	}
	if if len(s) > 0 {
		for _, char := range text {
			fmt.Print(string(char))
			time.Sleep(time.Duration(speed) * time.Millisecond)
		}
	}
}
