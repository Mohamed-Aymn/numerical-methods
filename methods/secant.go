package methods 

import (
  "fmt"
  "math"
  "sqrt/common"
)

func Secant (d *common.Data) {
  // calculateN
  n, _ := secantCalculateN(d.A, d.B, d.E)

  // construct record
  record := Record{
    A: d.A,
    B: d.B,
  }

  // iterate
  secantIterate(&record, &d.F, &n, d.E)

  fmt.Println("------------------- ")
  fmt.Println("*** The Final Result is: ", record.C)
  fmt.Println("------------------- ")
}

func secantCalculateN(a float64, b float64, e float64) (float64, error) {
  return math.Ceil((math.Log2((b-a) / (e)))), nil
}

func secantIterate (r *Record, f *string, n *float64, e float64) {
  var i = 1
  for{
    // f() calculations
    if i == 1{
      fa, _ := common.OneVariableFunction(*f, r.A)
      fb, _ := common.OneVariableFunction(*f, r.B)
      r.Fa = fa
      r.Fb = fb
    }

    r.C = (r.B - r.Fb) * ( (r.B - r.A) / (r.Fb - r.Fa) )
    fc, _ := common.OneVariableFunction(*f, r.C)
    r.Fc = fc

    // logger
    fmt.Println(i, "===>  a= ", r.A, ", b= ", r.B, ", c= ", r.C, ", f(a)=", r.Fa, ", f(b)=", r.Fb, ", f(c)=", r.Fc)

    // stop condition check
    if float64(i) == *n || math.Abs(r.Fc) < e {
      break
    }

    // swap
    r.A = r.B
    r.B = r.C
    r.Fa = r.Fb
    r.Fb = r.Fc

    // increament
    i++
  }
}

