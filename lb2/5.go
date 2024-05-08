package lb2

import "strconv"


/*
Task:
- Обчислити 7^123
*/
func Lab2Task5(base int, power int) string {
  result := "1"

  for i := 1; i <= power; i++ {
    result = multiplyBigNumberByDigit(result, base)
  }

  return result
}

func multiplyBigNumberByDigit(number string, digit int) string {
  result := ""
  tmp := number 
  length := len(number)
  carryDigit := 0

  for i := length - 1; i >= 0; i-- {
    numberDigit, _ := strconv.Atoi(string(tmp[i]))

    expr := numberDigit * digit + carryDigit
    last := expr % 10
    carryDigit = expr / 10

    result = strconv.Itoa(last) + result
  }

  if carryDigit > 0 {
    return strconv.Itoa(carryDigit) + result // Prepend carry if any
  }

  return result
}
