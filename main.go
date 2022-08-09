package main

import (
	"Kluster/pkg/apis/myoperator/v1myopr"
	"fmt"
)

func main() {
	k := v1myopr.Kluster{}
	fmt.Println(k)
}