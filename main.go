package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/labstack/gommon/random"
	"github.com/spf13/cast"
)

func main() {
	next()
}

func next() {
	for true {
		heart := false

		if strings.Contains("2468", random.String(1, "1234567890")) {
			heart = true
		}

		if heart {
			heartX := fmt.Sprintf("%s%s%s", random.String(1, "23456"), random.String(1, "0123456789"), random.String(1, "0123456789"))
			heartY := fmt.Sprintf("%s%s%s", random.String(1, "456"), random.String(1, "0123456789"), random.String(1, "0123456789"))
			fmt.Println("heart point ", heartX, heartY)
			if err := exec.Command("adb", "shell", fmt.Sprintf("input tap %s %s", heartX, heartY)).Run(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Millisecond*150)
			fmt.Println("heart click sleep time ", time.Millisecond*150)
			if err := exec.Command("adb", "shell", fmt.Sprintf("input tap %s %s", heartX, heartY)).Run(); err != nil {
				log.Fatal(err)
			}
			fmt.Println("heart clicked")
		}
		time.Sleep(time.Second * 3)

		startX := fmt.Sprintf("%s%s%s", random.String(1, "3456"), random.String(1, "0123456789"), random.String(1, "0123456789"))
		startY := fmt.Sprintf("1%s%s%s", random.String(1, "789"), random.String(1, "0123456789"), random.String(1, "0123456789"))
		fmt.Println("start point ", startX, startY)

		endX := fmt.Sprintf("%s%s%s", random.String(1, "4567"), random.String(1, "0123456789"), random.String(1, "0123456789"))
		endY := fmt.Sprintf("1%s%s%s", random.String(1, "012"), random.String(1, "0123456789"), random.String(1, "0123456789"))
		fmt.Println("end point ", endX, endY)

		if err := exec.Command("adb", "shell", fmt.Sprintf("input swipe %s %s %s %s", startX, startY, endX, endY)).Run(); err != nil {
			log.Fatal(err)
		}

		nextTime := cast.ToDuration(fmt.Sprintf("%s%s", random.String(1, "01"), random.String(1, "3456789")))
		sleepTime := time.Second*nextTime
		fmt.Println("next time ... ", sleepTime)
		fmt.Println("...")
		fmt.Println("..")
		fmt.Println(".")
		fmt.Println("")
		time.Sleep(sleepTime)
	}
}
