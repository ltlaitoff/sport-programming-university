package lb2

import "math/big"

/*
Task:
- Скласти програму обчислення точного значення n^n, де n > 10.
*/

func Lab2Task2(number *big.Int, power *big.Int) *big.Int {

	if power.Sign() == 0 {
		return big.NewInt(1)
	}

	if power.Sign() < 0 {
		return big.NewInt(-1)
	}

	if big.NewInt(0).Mod(power, big.NewInt(2)).Sign() == 1 {
		newPower := big.NewInt(0).Div(big.NewInt(0).Sub(power, big.NewInt(1)), big.NewInt(2))
		res := Lab2Task2(number, newPower)
		resultForReturn := big.NewInt(0).Mul(res, res)
		return big.NewInt(0).Mul(resultForReturn, number)
	}

	newPower := big.NewInt(0).Div(power, big.NewInt(2))
	res := Lab2Task2(number, newPower)

	return big.NewInt(0).Mul(res, res)
}
