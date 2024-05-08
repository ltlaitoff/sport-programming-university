package lb2

import (
  "strconv"
)

/*
Task:
Скласти програму обчислення точного значення суми 1! + 2! + 3! + ⋯ + n!
при умові, що n > 10.
*/
func Lab2Task7(base int) string {
  
  cache := map[int]string{}

  factorial(base, cache)

  result := "1"

  for _, value := range cache {
    result = summTwoBigNumbers(result, value)
  }

  return result 
}

func factorial(number int, cache map[int]string) string {
  if (number <= 0) {
    return "1"
  }

  if (cache[number] != "") {
    return cache[number]
  }

  cache[number] =  multiplyBigNumberByDigit(factorial(number - 1, cache), number)

  return cache[number]
}

func summTwoBigNumbers(first string, second string) string {
    var result string
    carry := 0
    i, j := len(first)-1, len(second)-1

    for i >= 0 || j >= 0 || carry > 0 {
        var sum int
        if i >= 0 {
            sum += int(first[i] - '0')
            i--
        }
        if j >= 0 {
            sum += int(second[j] - '0')
            j--
        }
        sum += carry
        carry = sum / 10
        sum %= 10
        result = strconv.Itoa(sum) + result
    }

    return result
}
