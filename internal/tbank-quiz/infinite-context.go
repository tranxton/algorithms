package tbank_quiz

import (
	"fmt"
	"math"
)

func KostyaNextPeriodPayment() {
	var montlyPayment, montlyPackageSize, overdraftPrice, totalUsed int

	fmt.Scan(&montlyPayment)
	fmt.Scan(&montlyPackageSize)
	fmt.Scan(&overdraftPrice)
	fmt.Scan(&totalUsed)

	if totalUsed <= montlyPackageSize {
		fmt.Println(montlyPayment)

		return
	}

	total_overdraft := totalUsed - montlyPackageSize

	fmt.Println(montlyPayment + (total_overdraft * overdraftPrice))
}

func MinimalCutsCount() {
	var persons float64

	fmt.Scan(&persons)

	fmt.Println(math.Ceil(math.Log2(persons)))
}
