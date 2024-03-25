package methods

import (
  "testing"
  "sqrt/common"
)

func TestBisection(t *testing.T){
  testData := &common.Data{
    F: "x ^ 2 - 3",
    A: 1,
    B: 2,
    E: 0.01,
  }
  strategy := &BisectionStrategy{}
  result := strategy.Execute(testData)
 
  if result < 1.7 {
    t.Fatalf("not working")
  }
}
