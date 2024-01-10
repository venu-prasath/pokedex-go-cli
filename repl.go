package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"

  "github.com/venu-prasath/pokedex-go-cli/internal/pokeapi"
)

type config struct {
  pokeapiClient pokeapi.Client
  nextLocationUrl *string
  previousLocationUrl *string
}

func startRepl(cfg *config) {
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
      err := command.callback(cfg)
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
      "map": {
          name: "map",
          description: "Get next set of locations",
          callback: commandMapf
      },
      "mapb": {
          name: "mapb",
          description: "Get previous set of locations",
          callback: commandMapb
      }
    }
}


