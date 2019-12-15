package link

import (
  "fmt"
  "os"
  "path/filepath"
)

func Create(commands []string) {
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
