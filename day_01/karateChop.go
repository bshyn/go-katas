package karateChop

func BinaryChop(numberToFind int, orderedArr []int) int {
	if len(orderedArr) == 0 {
		return -1
	}

	middle := (len(orderedArr)/2)
	middleValue := orderedArr[middle]

	if middleValue == numberToFind {
		return middle
	}

	if middleValue > numberToFind {
		return BinaryChop(numberToFind, orderedArr[0:middle])
	}
	return BinaryChop(numberToFind, orderedArr[middle:]) + middle
}
