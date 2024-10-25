package main

import (
	"fmt"
	"investment-calculator/user"
)

func main() {
	var userX user.User
	userX, err := user.New("Wayne", "Mortenson")
	fmt.Println(userX.GetUserName(), err)
}

/*
	Previous section...
*/

// package main

// import (
// 	"fmt"
// 	"math"
// )

// const inflationRate = 2.5

// func main() {
// 	var investmentAmount float64
// 	var expectedReturnRate float64
// 	var years float64

// 	fmt.Println("Enter an investment amount")
// 	fmt.Scan(&investmentAmount)

// 	fmt.Println("Enter expected return rate")
// 	fmt.Scan(&expectedReturnRate)
// 	fmt.Println("Enter number of years")
// 	fmt.Scan(&years)

// 	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

// 	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
// 	formattedFRV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)

// 	fmt.Print(formattedFV, formattedFRV)
// }

// func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, frv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	frv = fv / math.Pow(1+inflationRate/100, years)
// 	return
// }
