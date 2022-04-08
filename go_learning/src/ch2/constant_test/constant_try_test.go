package contsant_test

import "testing"

const(
  Monday = iota + 1
  Tuesday
  Wednesday

)

func TestConstantTry(t *testing.T) {
  t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
  t.Log(Monday, Tuesday)
}
