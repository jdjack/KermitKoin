package main

import "fmt"

type Address struct {
  Front int
}

func (ad Address) toHex() string {
  pt1Hex := fmt.Sprintf("%X", ad.Front)

  return pt1Hex
}