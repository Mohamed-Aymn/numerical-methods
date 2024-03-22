package main 

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "flag"
  "sqrt/methods"
  "sqrt/common"
) 

func main () {
  // Parse command-line arguments
  dataPath := flag.String("d", "./data.json", "Path to data file")
  method := flag.String("m", "bisection", "Method, available methods: bisection, secant, modified-secant, newton-raphson and fixed-point")
  flag.Parse()

  // Read JSON file
  file, err := ioutil.ReadFile(*dataPath)
  if err != nil {
    fmt.Println("Error reading JSON file:", err)
    return
  }

  // serialize json file
  var data common.Data 
  if err := json.Unmarshal(file, &data); err != nil {
    fmt.Println("Error unmarshaling JSON:", err)
    return
  }

  // excute method
  switch *method {
  case "bisection":
    methods.Bisection(&data)
//  case "secant":
 // case "modified-secant":
  //case "newton-raphson":
  //case "fixed-point":
  default:
    fmt.Println("choose another method")
}
}
