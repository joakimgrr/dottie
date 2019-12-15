package variable

import (
  "strings"
  "fmt"
)

var variables = make(map[string] string)

func Handle(commands []string) {
  name := commands[1]
  value := commands[2]
  SetValue(name, value)
  fmt.Println("Added variable: ", name, " with value: ", value)
}

func SetValue(name string, value string) {
  variables[name] = value
}

func Value(possibleVariable string) string {
  if strings.ContainsAny(possibleVariable, "[]") {
    variableCleaner := strings.NewReplacer("[", "", "]", "")
    return variables[variableCleaner.Replace(possibleVariable)]
  }

  return possibleVariable;
}
