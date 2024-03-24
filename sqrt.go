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

  var context Context
  // set strategy
  switch *method {
    case "bisection":
    context.setStrategy(&methods.BisectionStrategy{})
    case "secant":
      methods.Secant(&data)
    case "modified-secant":
      methods.ModifiedSecant(&data)
    case "newton":
      methods.Newton(&data)
    case "fixed-point":
      methods.FixedPoint(&data)
    default:
      fmt.Println("choose another method")
  }

  // excute strategy
  context.executeStrategy(&data)
}


// Strategy pattern implentation
type Strategy interface {
  Execute(data *common.Data)
}
type Context struct {
  strategy Strategy
}
func newContext(strategy Strategy) *Context {
  return &Context{strategy: strategy}
}
func (c *Context) setStrategy(strategy Strategy) {
	c.strategy = strategy
}
func (c *Context) executeStrategy(d *common.Data) {
  c.strategy.Execute(d)
}


