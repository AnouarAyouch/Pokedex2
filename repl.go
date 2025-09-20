package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		for {
			bold := color.New(color.FgCyan, color.Bold).SprintFunc()
			fmt.Print(bold("Pokedex > "))
			if !reader.Scan() { // handle EOF
				done <- true
				return
			}
			words := cleanInput(reader.Text())
			if len(words) == 0 {
				continue
			}
			commandName := words[0]
			args := words[1:]
			command, exists := Commands()[commandName]
			if exists {
				b := &Config{}
				err := command.callback(b, args)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("\n Unknown command ")
				continue
			}
		}
	}()
	select {
	case <-sigs:
		fmt.Printf("\n")

	case <-done:
		fmt.Printf("\nkeyboard interrupt ")
	}

}

/*
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
*/
