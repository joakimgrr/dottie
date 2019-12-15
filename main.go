package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"

  "github.com/joakimgrr/dottie/commands/variable"
  "github.com/joakimgrr/dottie/commands/link"
)

func main() {
  fmt.Println("Dottie v0.01")

  file, _ := os.Open("dottie.conf")

  configScanner := bufio.NewScanner(file)

  for configScanner.Scan() {
    handleLine(configScanner.Text())
  }
}

func handleLine(line string) {
  if line == "" {
    fmt.Println("Line empty, skipping")
    return
  }

  commands := strings.Fields(line)
  verb := commands[0]
  fmt.Println("verb: ", verb)

  handleVerb(verb, commands)
}

func handleVerb(verb string, commands []string) {
  switch verb {
  case "VAR":
    fmt.Println("variable found!")
    variable.Handle(commands)

  case "LINK":
    fmt.Println("Link found!")
    link.Handle(commands)
  }
}
