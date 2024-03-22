package methods

import (
	"fmt"
	"math"
	"os"
	"sqrt/common"
)

func Bisection(d *common.Data) {

	// check step
	check(d.F, d.A, d.B)
	// calculate number of iteration
	n, _ := calculateN(d.A, d.B, d.E)

	// construct record
	record := Record{
		A: d.A,
		B: d.B,
	}
	// iterate
	iterate(&record, &d.F, &n, d.E)

	fmt.Println("------------------- ")
	fmt.Println("*** The Final Result is: ", record.C)
	fmt.Println("------------------- ")
}

func calculateN(a float64, b float64, e float64) (float64, error) {
  return math.Round((math.Log2((b-a) / (e)))), nil
}

func check(f string, a float64, b float64) {
	fa, _ := common.OneVariableFunction(f, a)
	fb, _ := common.OneVariableFunction(f, b)
	if (fa * fb) > 0 {
		fmt.Println("This function will not be calculated")
		os.Exit(1)
	}
}

func iterate(r *Record, f *string, n *float64, e float64) {
	var i = 1
	for {
		// f() calcualtions
		r.C = (r.A + r.B) / 2
		fa, _ := common.OneVariableFunction(*f, r.A)
		fb, _ := common.OneVariableFunction(*f, r.B)
		fc, _ := common.OneVariableFunction(*f, r.C)
		r.Fa = fa
		r.Fb = fb
		r.Fc = fc

		// logger
		fmt.Println(i, "===>  a= ", r.A, ", b= ", r.B, ", c= ", r.C, ", f(a)=", r.Fa, ", f(b)=", r.Fb, ", f(c)=", r.Fc)

		// stop condition check
		if float64(i) == *n || math.Abs(r.Fc) < e {
			break
		}

		// swap
		if r.Fa*r.Fc > 1 {
			r.A = r.C
		} else {
			r.B = r.C
		}

		// increment the counter
		i++
	}
}
