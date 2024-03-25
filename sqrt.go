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
    context.setStrategy(&methods.SecantStrategy{})
    case "modified-secant":
    context.setStrategy(&methods.ModifiedSecantStrategy{})
    case "newton":
    context.setStrategy(&methods.NewtonStrategy{})
    case "fixed-point":
    context.setStrategy(&methods.FixedPointStrategy{})
    default:
      fmt.Println("choose another method")
  }

  // excute strategy
  context.executeStrategy(&data)
}

// Strategy pattern implentation
type Strategy interface {
  Execute(data *common.Data) float64
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

