package main

import (
  "os"
  "mathgen"
)

func main() {
  file, err := os.Create("Output.html")
  if err != nil {
    panic("Cannot create output file")
  }

  mathgen.Generate(file, 25, 30)

  file.Close()
}
