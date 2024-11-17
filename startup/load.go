package startup

import (
	"fmt"
	"time"
	"github.com/japser34/repl-os/tools"
	"os/exec"
)

func Start() {
	asciiArt := `
  _____  ______ _____  _
 |  __ \|  ____|  __ \| |
 | |__) | |__  | |__) | |  _____   ___  ___
 |  _  /|  __| |  ___/| | |_____| / _ \/ __|
 | | \ \| |____| |    | |____    | (_) \__ \
 |_|  \_\______|_|    |______|    \___/|___/ `
	goText := "Go-version \n \n"
	
	fmt.Print("Loading")
	tools.TypePrint(500, "...\n")
	
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	tools.TypePrint(15, asciiArt)
	time.Sleep(150 * time.Millisecond)
	tools.TypePrint(130, goText)
}
