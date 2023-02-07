package latihanpack

func sum(number ...int) int {
	var totalSum int = 0
	for _, num := range number {
		totalSum = totalSum + num
	}

	return totalSum
}
