package methods 

import (
  "fmt"
  "math"
  "sqrt/common"
)

type FixedPointRecord struct{
  X float64
  Fx float64
  Gx float64
}

type FixedPointStrategy struct {}

func (s *FixedPointStrategy) Execute(d *common.Data) {

  // construct record
  record := FixedPointRecord{
    X: d.X,
  }

  // iterate
  fixedPointIterate(&record, &d.F, &d.G , d.E, &d.N)

  fmt.Println("------------------- ")
  fmt.Println("*** The Final Result is: ", record.Gx)
  fmt.Println("------------------- ")
}

func fixedPointIterate (r *FixedPointRecord, f *string, g *string, e float64, n *int) {
  var i = 1
  for{
    // f() calculations
    fx, _ := common.OneVariableFunction(*f, r.X)
    gx, _ := common.OneVariableFunction(*g, r.X)
    r.Fx = fx
    r.Gx = gx

    // logger
    fmt.Println(i, "===>  x= ", r.X, ", f(xi)= ", r.Fx, ", G(xi)= ", r.Gx) 

    // stop condition check
    //fmt.Println("n is", *n)
    //break
    if i == *n || math.Abs(r.Fx) < e {
      break
    }

    // swap
    r.X = r.Gx

    // increament
    i++
  }
}
