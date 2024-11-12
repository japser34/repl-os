package startup

import (
	"fmt"
	"time"
)

func Start() {
	asciiArt :=`
  _____  ______ _____  _                   
 |  __ \|  ____|  __ \| |                  
 | |__) | |__  | |__) | |  ______ ___  ___ 
 |  _  /|  __| |  ___/| | |______/ _ \/ __|
 | | \ \| |____| |    | |____   | (_) \__ \
 |_|  \_\______|_|    |______|   \___/|___/ GO-version `
	fmt.Println("Loading")
	for i in range(3) {
		fmt.Print(".")
		time.Sleep(.5 * time.Second)
	}
	fmt.Println(asciiArt)
}
