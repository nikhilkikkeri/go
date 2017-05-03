package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v",float64(e))
}


func Sqrt(x float64) (zn float64, e error) {
	if(x < 0) {
		e = ErrNegativeSqrt(x)
		zn = x
	} else {
		zn = 1.0
		for {
			znp1 := zn - (((zn*zn)-x)/2*zn)
			if(math.Abs(zn-znp1) < 0.001) {
				break;
			} else {			
	    		zn = znp1
			}
		}
	}
	return zn, e
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}