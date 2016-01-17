package slurp

import "fmt"
import "testing"


func Test_Readfile_returns_a_uint8_array(t *testing.T) {

  expected := "[]uint8"

  lines,_ := slurp("foo.dat")
  actual := fmt.Sprintf("%T", lines)
  if actual != expected {
    fmt.Println("expected: ", expected, "--- actual: ", actual)
    t.Fail()
  }
}
