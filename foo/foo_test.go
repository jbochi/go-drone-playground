package foo

import "testing"

func TestSomething(t *testing.T) {
  if addsTwo(10) != 12 {
      t.Errorf("Can't sum!")
  }
}
