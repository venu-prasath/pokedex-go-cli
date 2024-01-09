package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func startRepl() {
  scanner := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()

    words := cleanInput(scanner.Text())
    if len(words) == 0 {
      continue
    }
    commandName := words[0]

    command, exists := getCommands()[commandName]
    if exists {
      err := command.callback()
      if err != nil {
        fmt.Println(err)
      }
      continue
    } else {
      fmt.Println("Unknown command")
      continue
    }

  }
}

func cleanInput(text string) []string {
  output := strings.ToLower(text)
   words := strings.Fields(output)
   return words
}

type cliCommand struct {
    name string
    description string
    callback func() error
}

func getCommands() map[string]cliCommand {
  return map[string]cliCommand {
      "help": {
          name: "help",
          description: "Display help message",
          callback: commandHelp,
        },
      "exit": {
          name: "exit",
          description: "Exit the app",
          callback: commandExit,
        },
    }
}

