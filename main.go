package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "flag"
)

type Data struct {
  F string  `json:"f"`
  A float32 `json:"a"`
  B float32 `json:"b"`
  E float32 `json:"e"`
}

func main () {
  // Parse command-line arguments
  dataPath := flag.String("d", "./data.json", "Path to data file")
  // method := flag.String("m", "bisection", "Method")
  flag.Parse()

  // Read JSON file
  file, err := ioutil.ReadFile(*dataPath)
  if err != nil {
    fmt.Println("Error reading JSON file:", err)
    return
  }

  // sericalize json file
  var data Data 
  if err := json.Unmarshal(file, &data); err != nil {
    fmt.Println("Error unmarshaling JSON:", err)
    return
  }
}
