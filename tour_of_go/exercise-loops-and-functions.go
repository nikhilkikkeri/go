package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	zn := 1.0
	//count := 0
	for {
		znp1 := zn - (((zn*zn)-x)/2*zn)
		//fmt.Println(znp1)
        //count++
		if(math.Abs(zn-znp1) < 0.0001) {
			break;
		} else {			
	    	zn = znp1
		}
	}
	return zn
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}


package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	zn := 1.0
    for {
		if znp1 := zn - (((zn*zn)-x)/2*zn); math.Abs(zn-znp1) < 0.0001 {
			break;
		} else {			
	    	zn = znp1
		}
	}
	return zn
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
