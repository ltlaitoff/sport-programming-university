package lb1

// import "fmt"
import "math/big"

/*
Task:
Звести число a в натуральну ступінь n за якомога меншу кількість множень.
*/
func Task8(number *big.Int, power *big.Int) *big.Int {

	if power.Sign() == 0 {
		return big.NewInt(1)
	}

	if power.Sign() < 0 {
		return big.NewInt(-1)
	}

	if big.NewInt(0).Mod(power, big.NewInt(2)).Sign() == 1 {
		newPower := big.NewInt(0).Div(big.NewInt(0).Sub(power, big.NewInt(1)), big.NewInt(2))
		res := Task8(number, newPower)
		resultForReturn := big.NewInt(0).Mul(res, res)
		return big.NewInt(0).Mul(resultForReturn, number)
	}

	newPower := big.NewInt(0).Div(power, big.NewInt(2))
	res := Task8(number, newPower)

	return big.NewInt(0).Mul(res, res)
}

func task8Default(number *big.Int, power *big.Int) *big.Int {
	result := number

	i := big.NewInt(0)
	for ; i.Cmp(big.NewInt(0).Sub(power, big.NewInt(1))) < 0; i.Add(i, big.NewInt(1)) {
		result = big.NewInt(0).Mul(result, number)
	}

	return result
}
