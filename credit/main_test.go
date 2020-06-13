package main_test // взяли пакет credit, добавили _test

import (
	"fmt"
	credit "github.com/lizaMosiyash/bgo-1_homeworks_2/pkg"
)

func ExampleCalculate() { // имя функции - Example + имя проверяемой функции
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20))
	fmt.Println(credit.Calculate(10_000_00, 36, 20))
	// Output:
	// 3718400 33862300 133862300
	// 37184 338623 1338623
}