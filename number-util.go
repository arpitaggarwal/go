package main

import (
	"fmt"

	"github.com/arpitaggarwal/numberutil"
)

func main() {
	i23 := int64(23)
	fmt.Printf("Decimal to Binary for 23 is %s \n", numberutil.DecimalToBinary(i23))
}
