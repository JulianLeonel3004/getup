package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Minutes parameter not found")
		return
	}

	minutes, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Minutes parameter should be an integer")
		return
	}

	fmt.Println("run")
	getUp(minutes)

}

func getUp(minutes int) {

	now := time.Now()

	for now.Hour() >= 9 && now.Hour() < 18 {

		time.Sleep(time.Duration(minutes) * time.Minute)
		alertUp()

		now = time.Now()
	}

}

func alertUp() {
	fmt.Println("levantate")

	script := `tell application "Finder"
						activate
						display dialog "LEVANTATE."
					end tell`

	cmd := exec.Command("osascript", "-e", script)
	cmd.Run()
}
