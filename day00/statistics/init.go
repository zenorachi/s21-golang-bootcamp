package statistics

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type StatMetrics struct {
	arrayOfIntegers []int
	lengthOfArray   int
	maxElem         int
	minElem         int
}

func (sm *StatMetrics) fillFields() {
	sort.Ints(sm.arrayOfIntegers)
	sm.lengthOfArray = len(sm.arrayOfIntegers)
	if sm.lengthOfArray > 0 {
		sm.maxElem = sm.arrayOfIntegers[sm.lengthOfArray-1]
		sm.minElem = sm.arrayOfIntegers[0]
	}
}

func (sm *StatMetrics) Init() {
	var input string
	var err error
	var reader = bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input = reader.Text()
		num, errParse := strconv.Atoi(input)
		isRightBounds := num >= -100_000 && num <= 100_000
		if errParse != nil || !isRightBounds {
			fmt.Println("Error: input must be an integer number")
			continue
		}
		if err != io.EOF {
			sm.arrayOfIntegers = append(sm.arrayOfIntegers, num)
		}
	}
	sm.fillFields()
}

func (sm *StatMetrics) SetArray(array []int) {
	if array != nil {
		sm.arrayOfIntegers = array
		sm.fillFields()
	}
}

func (sm *StatMetrics) ReadFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = strings.Trim(input, "\n")
		num, errParse := strconv.Atoi(input)
		isRightBounds := num >= -100_000 && num <= 100_000
		if errParse == nil && isRightBounds {
			sm.arrayOfIntegers = append(sm.arrayOfIntegers, num)
		}
	}
	sm.fillFields()
}
