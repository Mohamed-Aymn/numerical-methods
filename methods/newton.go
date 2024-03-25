package methods 

import (
  "fmt"
  "math"
  "sqrt/common"
)

type NewtonRecord struct{
  X float64
  Fx float64
  Fdx float64
}

type NewtonStrategy struct {}

func (s *NewtonStrategy) Execute(d *common.Data) float64 {

  // construct record
  record := NewtonRecord{
    X: d.X,
  }

  // iterate
  newtonIterate(&record, &d.F, &d.D , d.E, &d.N)

  fmt.Println("------------------- ")
  fmt.Println("*** The Final Result is: ", record.X)
  fmt.Println("------------------- ")

  return record.X
}

func newtonIterate (r *NewtonRecord, f *string, d *string, e float64, n *int) {
  var i = 1
  for{
    // f() calculations
    fx, _ := common.OneVariableFunction(*f, r.X)
    fdx, _ := common.OneVariableFunction(*d, r.X)
    r.Fx = fx
    r.Fdx = fdx

    // logger
    fmt.Println(i, "===>  x= ", r.X, ", f(xi)= ", r.Fx, ", F`(xi)= ", r.Fdx) 

    // stop condition check
    //fmt.Println("n is", *n)
    //break
    if i == *n || math.Abs(r.Fx) < e {
      break
    }

    // swap
    r.X = r.X - (r.Fx / r.Fdx)

    // increament
    i++
  }
}

