package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func init() {

	resize = make(map[string]func())
	resize["darwin"] = func() {
		cmd := exec.Command("resize -s 35 120")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	resize["linux"] = func() {
		cmd := exec.Command("resize -s 35 120")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	resize["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "mode con:cols=80 lines=30") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clear = make(map[string]func()) //Initialize it
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	logpath := filepath.Join(".", "logs")
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		os.MkdirAll(logpath, os.ModePerm)
	}

	jobpath := filepath.Join(".", "jobs")
	if _, err := os.Stat(jobpath); os.IsNotExist(err) {
		os.MkdirAll(jobpath, os.ModePerm)
	}
}

func checkErr(err error) {
	if err != nil {
		// log.Fatalln(err)
		log.Println(err)
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	//if we defined a clear func for that platform:
	if ok {
		value() //we execute it
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(") //unsupported platform
	}
}

func ResizeWindow() {
	value, ok := resize[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	//if we defined a clear func for that platform:
	if ok {
		value() //we execute it
	} else {
		panic("Your platform is unsupported! I can't resize terminal screen :(") //unsupported platform
	}
}
