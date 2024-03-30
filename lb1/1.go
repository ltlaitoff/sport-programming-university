package lb1

/*
Task
Задано масив M[1:N] натуральних чисел, упорядкований за не спаданням, тобто:
M[1] ← M[2] ← ⋯ ← M[N].
Знайти перше натуральне число, яке не представляється сумою ніяких елементів цього
масиву, при цьому сума може складатися і з одного доданка, але кожен елемент масиву може
входити в неї тільки один раз.
*/
func Task1() int {
	return 1
}

// /* Its genious */
// function chatVersion(arr: number[]): number {
// 	let summ = 1

// 	for (const num of arr) {
// 		if (num > summ) {
// 			break
// 		}

// 		summ += num
// 	}

// 	return summ
// }

// /* Its our :D */
// function main(set: Set<number>, array: number[]) {
// 	if (array.length === 1) return array

// 	array.forEach(item => set.delete(item))

// 	const firstElement = array[0]
// 	const slicedArray = array.slice(1)

// 	main(set, slicedArray)

// 	slicedArray.map((_, index, array) => {
// 		const arr = [...array]
// 		arr[index] += firstElement
// 		main(set, arr)
// 	})

// 	return set
// }

// const a = [1, 2, 7, 10]

// console.time('my')

// const set = new Set(
// 	Array.from(
// 		{ length: a.reduce((acc, item) => acc + item, 0) + 1 },
// 		(_: any, index: number) => index + 1
// 	)
// )

// main(set, a)

// console.log('our: ', Array.from(set)[0])

// console.timeEnd('my')

// console.time('chat')

// console.log('chat:', chatVersion(a))

// console.timeEnd('chat')
