package lb3

// Lab3Task8
// Маємо N каменів ваги A1, A2,..., АN. Необхідно розбити їх на дві купи таким
// чином, щоб ваги Куп відрізнялися не більше ніж в 2 рази. Якщо цього зробити не можна,
// то вказати це.

type Result struct {
	weight int
	stones []int
}

func Lab3Task8(stones []int) (stones1 []int, stones2 []int, err bool) {
	totalWeight := 0

	for _, stone := range stones {
		totalWeight += stone
	}

	target := totalWeight / 2

	var innerLab3Task8 func(index int, weint int, selectedStones []int) *Result
	innerLab3Task8 = func(index int, weight int, selectedStones []int) *Result {
		if weight > target {
			return nil
		}

		if index == len(stones) {
			return &Result{weight: weight, stones: selectedStones}
		}

		includeStone := append(selectedStones, stones[index])
		include := innerLab3Task8(index+1, weight+stones[index], includeStone)
		exclude := innerLab3Task8(index+1, weight, selectedStones)

		if include != nil && (exclude == nil || include.weight > exclude.weight) {
			return include
		}

		return exclude
	}

	result := innerLab3Task8(0, 0, []int{})

	if result == nil {
		return nil, nil, true
	}

	weight1 := result.weight
	stones1 = result.stones

	stones1Map := make(map[int]int)
	for _, stone := range stones1 {
		stones1Map[stone]++
	}

	for _, stone := range stones {
		if stones1Map[stone] > 0 {
			stones1Map[stone]--
		} else {
			stones2 = append(stones2, stone)
		}
	}

	weight2 := 0
	for _, stone := range stones2 {
		weight2 += stone
	}

	if weight1 <= 2*weight2 && weight2 <= 2*weight1 {
		return stones1, stones2, false
	}

	return nil, nil, true
}
