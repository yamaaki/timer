package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/nsf/termbox-go"
)

var x, y int

var empty = []string{
	"        ",
	"        ",
	"        ",
	"        ",
	"        ",
}

var digits = [][]string{
	{
		" ██████ ",
		"██  ████",
		"██ ██ ██",
		"████  ██",
		" ██████ ",
	},
	{
		"     ██ ",
		"  █████ ",
		"     ██ ",
		"     ██ ",
		"     ██ ",
	},
	{
		" ██████ ",
		"      ██",
		"  █████ ",
		" ██     ",
		" ███████",
	},
	{
		" ██████ ",
		"      ██",
		"  █████ ",
		"      ██",
		" ██████ ",
	},
	{
		" ██   ██",
		" ██   ██",
		" ███████",
		"      ██",
		"      ██",
	},
	{
		" ███████",
		" ██     ",
		" ███████",
		"      ██",
		" ███████",
	},
	{
		" ██████ ",
		"██      ",
		"███████ ",
		"██    ██",
		" ██████ ",
	},
	{
		" ███████",
		"      ██",
		"     ██ ",
		"    ██  ",
		"    ██  ",
	},
	{
		"  █████ ",
		" ██   ██",
		"  █████ ",
		" ██   ██",
		"  █████ ",
	},
	{
		"  █████ ",
		" ██   ██",
		"  ██████",
		"      ██",
		"  █████ ",
	},
}

func print(n int) {
	if n < 0 || n > 30 {
		fmt.Println("Number out of range")
		return
	}

	var left, right []string
	if n <= 9 {
		left = empty
	} else {
		left = digits[n/10]
	}
	right = digits[n%10]

	lines := make([]string, 5)
	for i := 0; i < 5; i++ {
		lines[i] = left[i] + right[i]
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for i, line := range lines {
		j := 0
		for _, c := range line {
			j++
			termbox.SetCell(x+j, y+i, rune(c), termbox.ColorWhite, termbox.ColorDefault)
		}
	}
	termbox.Flush()
}

func sound() {
	cmd := exec.Command("afplay", "/System/Library/Sounds/Tink.aiff")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("音の再生に失敗しました: %v", err)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetCursor(-1, -1)

	_, height := termbox.Size()

	x = 8
	y = height - 12

	go func() {
		for i := 1; i <= 30; i++ {
			go sound()
			print(i)
			time.Sleep(1000 * time.Millisecond)
		}

		termbox.Close()
		os.Exit(0)
	}()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' || ev.Key == termbox.KeyEsc {
				termbox.Close()
				os.Exit(0)
			}
		}
	}
}
