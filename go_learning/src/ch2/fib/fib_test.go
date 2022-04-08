package try_test

import "testing"
import "fmt"

func TestFirstTry(t *testing.T)  {
  var a int = 1
  var b int = 1

  // fmt.Print(a)
  t.Log(a)

  for i :=0; i<5; i++ {
    t.Log(" ", b)
    tmp := a
    a = b
    b = tmp + a
  }


  t.Log("My First try!")
}
