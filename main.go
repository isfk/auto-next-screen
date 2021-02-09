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
		swipePoint()
		heartPoint()
	}
}

func heartPoint() (x,y string) {
	heart := false
	if strings.Contains("2468", random.String(1, "1234567890")) {
		heart = true
	}

	if heart {
		heartX := fmt.Sprintf("%s%s%s", random.String(1, "3456"), random.String(1, "0123456789"), random.String(1, "0123456789"))
		heartY := fmt.Sprintf("%s%s%s", random.String(1, "456789"), random.String(1, "0123456789"), random.String(1, "0123456789"))
		fmt.Println("heart point ", heartX, heartY)
		if err := exec.Command("adb", "shell", fmt.Sprintf("input tap %s %s", heartX, heartY)).Run(); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 5)
		if err := exec.Command("adb", "shell", fmt.Sprintf("input tap %s %s", heartX, heartY)).Run(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("heart clicked")
	}
	time.Sleep(time.Second * 1)
	return
}

func swipePoint() {
	nextTimes := []int{
		5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	}

	startX4 := ""
	startX3 := random.String(1, "3456")
	startX2 := random.String(1, "0123456789")
	startX1 := random.String(1, "0123456789")
	startX := fmt.Sprintf("%s%s%s%s", startX4, startX3, startX2, startX1)

	startY4 := "1"
	startY3 := random.String(1, "789")
	startY2 := random.String(1, "0123456789")
	startY1 := random.String(1, "0123456789")
	startY := fmt.Sprintf("%s%s%s%s", startY4, startY3, startY2, startY1)

	endX4 := ""
	endX3 := cast.ToInt(startX3) + 3
	endX2 := random.String(1, "0123456789")
	endX1 := random.String(1, "0123456789")
	endX := fmt.Sprintf("%s%d%s%s", endX4, endX3, endX2, endX1)

	endY4 := "1"
	endY3 := cast.ToInt(startY3) - 5
	endY2 := random.String(1, "0123456789")
	endY1 := random.String(1, "0123456789")
	endY := fmt.Sprintf("%s%d%s%s", endY4, endY3, endY2, endY1)

	if err := exec.Command("adb", "shell", fmt.Sprintf("input swipe %s %s %s %s", startX, startY, endX, endY)).Run(); err != nil {
		log.Fatal(err)
	}

	nextTime := cast.ToDuration(nextTimes[cast.ToInt(random.String(1, "1234567890"))])
	sleepTime := time.Second * nextTime

	fmt.Println("next time ... ", sleepTime)
	fmt.Println("...", "..", ".", "")
	time.Sleep(sleepTime)
	return
}
