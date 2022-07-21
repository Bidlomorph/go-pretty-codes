package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	result, err := prettyNumber(9, 3, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("collection", result)

	result, err = prettyNumberSlice(9, 3, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("slice", result)
}

// prettyNumber can return pretty code dependent on codeLength(length of returned code),
// orderLength(length of maximum same numbers) and orderCount(number of sets of identical numbers)
func prettyNumber(codeLength, orderLength, orderCount uint64) (uint64, error) {
	rand.Seed(time.Now().UnixNano())

	maxCountOfValues := int(codeLength - (orderLength * orderCount) + orderCount)

	if maxCountOfValues <= 1 {
		return 0, errors.New("[error] the orderLength does not match the codeLength \n" +
			"it depends on expression codeLength - (orderLength * orderCount) + orderCount")
	}

	collection := make(map[int]string)

	for i := 0; maxCountOfValues > i; i++ {
		number := rand.Int63n(9 - 1)
		collection[i] = strconv.FormatInt(number, 10)
	}

	for i := 0; int(orderCount) > i; i++ {
		multipleValue := strings.Repeat(collection[i], int(orderLength))
		collection[i] = multipleValue
	}

	var finalString string
	for _, v := range collection {
		finalString += v
	}

	return strconv.ParseUint(finalString, 10, 64)
}

// prettyNumberSlice can return pretty code dependent on codeLength(length of returned code),
// orderLength(length of maximum same numbers) and orderCount(number of sets of identical numbers)
func prettyNumberSlice(codeLength, orderLength, orderCount uint64) (uint64, error) {
	rand.Seed(time.Now().UnixNano())

	maxCountOfValues := int(codeLength - (orderLength * orderCount) + orderCount)

	if maxCountOfValues <= 1 {
		return 0, errors.New("[error] the orderLength does not match the codeLength \n" +
			"it depends on expression codeLength - (orderLength * orderCount) + orderCount")
	}

	sliceCollection := make([]string, 0, maxCountOfValues)

	for i := 0; maxCountOfValues > i; i++ {
		number := rand.Int63n(9 - 1)
		sliceCollection = append(sliceCollection, strconv.FormatInt(number, 10))
	}

	for i := 0; int(orderCount) > i; i++ {
		multipleValue := strings.Repeat(sliceCollection[i], int(orderLength))
		sliceCollection[i] = multipleValue
	}

	rand.Shuffle(len(sliceCollection), func(i, j int) {
		sliceCollection[i], sliceCollection[j] = sliceCollection[j], sliceCollection[i]
	})

	return strconv.ParseUint(strings.Join(sliceCollection[:], ""), 10, 64)
}
