package methods 

import(
  "fmt"
  "sqrt/common"
  "os"
  "math"
)

func Bisection (d *common.Data){

  check(d.F, d.A, d.B)
  n, _ := calculateN(d.A, d.B, d.E)

  // construct record
  record := Record{
    A: d.A,
    B: d.B,
    C: (d.A + d.B) / 2,
  }
  fa, _ := common.OneVariableFunction(d.F, d.A) 
  fb, _ := common.OneVariableFunction(d.F, d.B)
  fc, _ := common.OneVariableFunction(d.F, (d.A + d.B) / 2)
  record.Fa = fa
  record.Fb = fb
  record.Fc = fc

  iterate(&record, &n, d.E)
}

func check(f string, a float64, b float64){
  // calculate f(a)
//  fa := strings.Replace(F, "x", fmt.Sprintf("%.2f", A), -1)
  ra, err := common.OneVariableFunction(f, a)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  // calculate f(b)
 // fb := strings.Replace(F, "x", fmt.Sprintf("%.2f", B), -1)
  rb, err := common.OneVariableFunction(f, b)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  if((ra * rb) > 0){
    fmt.Println("This function cannot be calculated")
    os.Exit(1)
  }
}

func calculateN(a float64, b float64, e float64) (float64, error) {
  return math.Round((math.Log2((b-a) / (e)))), nil
}

func iterate(r *Record, n *float64, e float64){
  fmt.Println(*r)
  // f()

  // swap

  // logger

  // stop condition
}
