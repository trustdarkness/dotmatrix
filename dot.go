// Package dot computes the dot product of two matrices.
// It can be run interactively on the console or scripted using
// csv files for input and output.
package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strconv"
	"log"
)

type matrix struct {
	rows int
	cols int
	row_major bool
    data []int
}

// Product computes the Product of two matrices and returns a matrix object.
// If the matrices cannot be multiplies, an error is printed.
func Product(m1 matrix, m2 matrix) (matrix) {
	if (m1.rows != m2.cols) {
		log.Println("You cannot compute the dot Product of these matrices")
		os.Exit(0)
	}

	// m1 should be row major and m2 should be column major.
	if !m1.row_major {
		m1 = Convert(m1)
	}

	if m2.row_major {
		m2 = Convert(m2)
	}

	new_len := m1.rows+m2.cols
	result_data := make([]int, new_len)

	for positionYdst := 0; positionYdst < m1.cols; positionYdst++ {
		for positionXdst := 0; positionXdst < m1.rows; positionXdst++ {
			for positionXsrc := 0; positionXsrc < m2.cols; positionXsrc++ {
				m2shift := positionXdst * m2.rows
				m1shift := positionYdst * m1.cols
				dstShift := positionYdst * m1.rows
				result_position := positionXdst+dstShift
				result_data[result_position] += m1.data[positionXsrc+m1shift] * m2.data[positionXsrc+m2shift]
			}
		}
	}

	var result = matrix{rows: m1.rows, cols: m2.cols, row_major: true, data: result_data}
	return result
}

// PrintMatrix prints a matrix to the console.
// Depends on PrintRowMajorMatrix
func PrintMatrix (m matrix) {
	if (!m.row_major) {
		var mrm = Convert(m)
		PrintRowMajorMatrix(mrm)
	} else {
		PrintRowMajorMatrix(m)
	}
}

// PrintRowMajorMatrix prints a row major matrix to the console.
// This is called by Print Matrix.
func PrintRowMajorMatrix(m matrix) {
	i := 0
	for c := 0; c < m.cols; c++ {
		fmt.Printf("[ ")
		for r := 0; r < m.rows; r++ {
		  fmt.Printf("%2d ", m.data[i])
		  i++
		}
		fmt.Printf("]\n")
	}
}

// Convert converts a matrix between row_major and column major.
func Convert(m matrix) (matrix) {
	var result_data = make([]int,len(m.data))
	var j = 0
	for i := 0; i < len(m.data)/2; i++ {
		result_data[j] = m.data[i]
		result_data[j+1] = m.data[m.rows+i]
		j = j+2
	}
	var ret = matrix{rows:m.rows, cols: m.cols, row_major: !m.row_major, data: result_data}
	return ret
}

// ProcessFile takes a user provided file name and attempts to build a matrix object from it.
func ProcessFile(file string) matrix {
	f, err := os.Open(file)
	if os.IsNotExist(err) {
		log.Println("%s does not seem to exist.")
		os.Exit(1)
	} else if err != nil {
		log.Println("", err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	var rows = 0
	var cols = 0
	var data = []int{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Error while processing %s: %s\n", file, err)
			os.Exit(1)
		}
		cols = len(line)
		for i := 0; i < cols; i++ {
			s, err := strconv.Atoi(line[i])
			if err != nil {
				log.Println("%s doesn't appear to be an int: %s", line[i], err)
				os.Exit(1)
			}
			data = append(data, s)
		}
		rows++
	}
	var m = matrix{rows: rows, cols: cols, row_major: true, data: data}
  	return m
}

// Write matrix writes the provided matrix to a csv file.
func WriteMatrix(m matrix, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println("", err)
		log.Println("I couldn't write the file because of the above error.")
		fmt.Println("The result matrix looks like:")
		PrintMatrix(m)
		os.Exit(1)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	i := 0
	for c := 0; c < m.cols; c++ {
		row := make([]string, m.cols)
		for r := 0; r < m.rows; r++ {
			row = append(row, strconv.Itoa(m.data[i]))
			i++
		}
		err = writer.Write(row)
		if err != nil {
			log.Fatal("Cannot write to file:", err)
		}
	}
}

func main() {
	var nameptr1 = flag.String("f1", "", "csv file containing matrix 1")
	var nameptr2 = flag.String("f2", "", "csv file containing matrix 2")
	var output = flag.String("o", "", "output file (Optional)")
	flag.Parse()
	var filename1 = *nameptr1
	var filename2 = *nameptr2
	var m1 = ProcessFile(filename1)
	var m2 = ProcessFile(filename2)
	var m = Product(m1, m2)
	if *output != "" {
		fmt.Println("Matrix 1 looks like:")
		PrintMatrix(m1)
		fmt.Println("Matrix 2 looks like:")
		PrintMatrix(m2)
		fmt.Println("Result looks like: ")
		PrintMatrix(m)
	} else {

	}

}


