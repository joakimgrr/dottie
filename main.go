package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "path/filepath"
)

var variables = make(map[string] string)

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
    handleVariableCommand(commands)

  case "LINK":
    fmt.Println("Link found!")
    handleLinkCommand(commands)
  }
}

func handleVariableCommand(commands []string) {
  variable := commands[1]
  value := commands[2]
  variables[variable] = value
  fmt.Println("Added variable: ", variable, " with value: ", value)
}

func getVariableValue(possibleVariable string) string {
  if strings.ContainsAny(possibleVariable, "[]") {
    variableCleaner := strings.NewReplacer("[", "", "]", "")
    return variables[variableCleaner.Replace(possibleVariable)]
  }

  return possibleVariable;
}

// TODO: handle absolute and relative paths correctly
func handleLinkCommand(commands []string) {
  sourcePath := commands[1]
  // TODO: use MkdirAll for paths
  endLocation := commands[2]
  actualLocation := endLocation

  // TODO: handle errors
  actualSourcePath, _ := filepath.Abs(sourcePath)
  fmt.Println("actual source ", actualSourcePath)

  // handle variables that are defined before using LOCATION
  actualLocation = getVariableValue(endLocation)
  fmt.Println("actual end: ", actualLocation)

  filename := filepath.Base(sourcePath)
  err := os.Symlink(actualSourcePath, filepath.Join(actualLocation, filename))

  if err != nil {
    fmt.Println(err)
  }
}
