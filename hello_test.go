package main

import "testing"


func TestHello(t *testing.T) {
  x := Sum(10, 10)
  if x != 20 {
    t.Errorf("bad sum result. expected %d, received %d", 20, x)
  }
}
