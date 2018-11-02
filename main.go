package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const version = "1.0"

func main() {
	if len(os.Args) < 3 {
		printUsage()
		return
	}
	seconds := parseTimeString(os.Args[1])
	if seconds <= 0 {
		fmt.Println("[autoexit] Invalid argument:", os.Args[1])
		printUsage()
		os.Exit(2)
		return
	}
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()
	fmt.Println("[autoexit] Exec:", strings.Join(os.Args[2:], " "))
	fmt.Printf("[autoexit] Automatic exit at %s\n", time.Now().Add(seconds).Format(time.RFC3339))

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
		return
	}
	fmt.Println("[autoexit] Start process with PID:", cmd.Process.Pid)

	go func() {
		time.Sleep(seconds)
		fmt.Println("[autoexit] It's the time to exit.")
		err := cmd.Process.Kill()
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(1)
	}()

	err = cmd.Wait()
	if err != nil {
		fmt.Println("[autoexit] Process", err)
		os.Exit(3)
	} else {
		fmt.Println("[autoexit] Process finished with exit status 0")
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf("autoexit v%s by Zongmin Lei <leizongmin@gmail.com> Copyright 2018\n", version)
	fmt.Println("Project: https://github.com/leizongmin/autoexit")
	fmt.Println("Usage:   autoexit <seconds> <command>")
	fmt.Println("Example: ")
	fmt.Println("         autoexit 3600 cmd arg1 arg2 ...           # auto exit after 3600 seconds")
	fmt.Println("         autoexit 3600s cmd arg1 arg2 ...          # auto exit after 3600 seconds")
	fmt.Println("         autoexit 2d+3h+4m+5s cmd arg1 arg2 ...    # auto exit after 3 days and 3 hours and 4 minutes and 5 seconds")
}
