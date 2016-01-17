package slurp

import "fmt"
import "strings"
import "testing"


func Test_Readfile_returns_a_uint8_array(t *testing.T) {

  // byte is an alias of uint8
  expected := "[]uint8"

  lines,_ := slurp("foo.dat")
  actual := fmt.Sprintf("%T", lines)
  if actual != expected {
    fmt.Println("expected: ", expected, "--- actual: ", actual)
    t.Fail()
  }
}

func Test_Readfile_returns_the_whole_file(t *testing.T) {
  expected := 9

  fileByteArray,_ := slurp("foo.dat")
  lines := strings.Split(string(fileByteArray), "\n")
  actual := len(lines)

  if expected != actual {
    fmt.Println("expected: ", expected, "--- actual: ", actual)
    t.Fail()
  }
}

