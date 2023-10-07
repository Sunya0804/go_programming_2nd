package ds_hw_0

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	// TODO: implement me
	// HINT: use for loop over `nums`
	tempSum := 0
	for num := range nums {
		tempSum += num
	}
	out <- tempSum
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	// TODO: implement me
	// HINT: use `readInts` and `sumWorkers`
	// HINT: used buffered channels for splitting numbers between workers

	// Open file
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	// Read integers from file
	ints, err := readInts(file)
	checkError(err)

	// Create buffered channel
	intsChan := make(chan int, len(ints))

	// Send integers to channel
	for _, num := range ints {
		intsChan <- num
	}

	// Create buffered channel
	sumChan := make(chan int, num)

	// Create workers
	for i := 0; i < num; i++ {
		go sumWorker(intsChan, sumChan)
	}

	// Close channel
	close(intsChan)

	// Sum integers
	sum := 0
	for i := 0; i < num; i++ {
		sum += <-sumChan
	}

	// Close channel
	close(sumChan)

	return sum
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
