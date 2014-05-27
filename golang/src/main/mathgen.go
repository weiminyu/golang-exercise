package main

import (
  "os"
  "mathgen"
)

func main() {
  outFileName := "Output.html"
  if len(os.Args) > 1 {
    outFileName = os.Args[1]
  }
  file, err := os.Create(outFileName)
  if err != nil {
    panic("Cannot create output file " + outFileName)
  }

  mathgen.Generate(file, 25, 40)

  file.Close()
}
