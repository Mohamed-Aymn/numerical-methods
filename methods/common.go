package methods

type Record struct {
  A float64 
  B float64
  C float64
  Fa float64
  Fb float64
  Fc float64
}

type NewtonRecord struct{
  X float64
  Fx float64
  Fdx float64
}

type FixedPointRecord struct{
  X float64
  Fx float64
  Gx float64
}
