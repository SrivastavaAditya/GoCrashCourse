package main

import (
	"fmt"
	"math"

	"github.com/SrivastavaAditya/go_crash_course/03_packages/strutils"
)

func main() {
	fmt.Println(math.Floor(33.4))
	fmt.Println(math.Ceil(33.6))
	fmt.Println(math.Sqrt(36))
	fmt.Println(strutils.Reverse("olleh"))
}
