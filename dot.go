// Package dot computes the dot product of two matrices.
// It can be run interactively on the console or scripted using
// csv files for input and output.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Matrix is a struct representation of a matrix with an underlying
// array of int in either row major or column major ordering
type Matrix struct {
	rows     int
	cols     int
	rowMajor bool
	data     []int
}

// Product computes the Product of two matrices and returns a Matrix object.
// If the matrices cannot be multiplies, an error is printed.
func Product(a Matrix, b Matrix) Matrix {
	if a.cols != b.rows {
		log.Println("The dot product of these matrices is not defined,")
		log.Print("see https://en.wikipedia.org/wiki/Dot_product.")
		os.Exit(0)
	}

	// m1 should be row major and m2 should be column major.
	if !a.rowMajor {
		a = Convert(a)
	}

	if b.rowMajor {
		b = Convert(b)
	}

	newLen := a.rows * b.cols
	resultData := make([]int, newLen)

	dstPos := 0
	bPos := 0
	for dstYPos := 0; dstYPos < a.rows; {
		for acol := 0; acol < a.cols; acol++ {
			if !(bPos < len(b.data)) {
				dstYPos++
				bPos = 0
				dstPos--
				break
			} else {
				aPos := acol + (dstYPos * a.cols)
				term1 := a.data[aPos]
				term2 := b.data[bPos]
				resultData[dstPos] += term1 * term2
				bPos++
			}
		}
		dstPos++
	}
	var result = Matrix{a.rows, b.cols, true, resultData}
	return result
}

// PrintMatrix prints a Matrix to the console.
// Depends on PrintRowMajorMatrix
func PrintMatrix(m Matrix) {
	if !m.rowMajor {
		var mrm = Convert(m)
		PrintRowMajorMatrix(mrm)
	} else {
		PrintRowMajorMatrix(m)
	}
}

// PrintRowMajorMatrix prints a row major Matrix to the console.
// This is called by Print Matrix.
func PrintRowMajorMatrix(m Matrix) {
	i := 0
	for r := 0; r < m.rows; r++ {
		fmt.Printf("[ ")
		for c := 0; c < m.cols; c++ {
			fmt.Printf("%2d ", m.data[i])
			i++
		}
		fmt.Printf("]\n")
	}
}

// Convert converts a Matrix between rowMajor and column major.
func Convert(m Matrix) Matrix {
	var resultData = make([]int, len(m.data))
	dstPos := 0
	for col := 0; col < m.cols; col++ {
		srcPos := 0
		for row := 0; row < m.rows; row++ {
			srcPos = row*m.cols + col
			resultData[dstPos] = m.data[srcPos]
			dstPos++
		}
	}
	var ret = Matrix{m.rows, m.cols, !m.rowMajor, resultData}
	return ret
}

// ProcessFile takes a user provided file name
// and attempts to build a Matrix object from it.
func ProcessFile(file string) Matrix {
	f, err := os.Open(file)
	if os.IsNotExist(err) {
		log.Printf("%s does not seem to exist.", file)
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
				log.Printf("%s doesn't appear to be an int: %s\n", line[i], err)
				log.Println("Check the format of your Matrix")
				os.Exit(1)
			}
			data = append(data, s)
		}
		rows++
	}
	var m = Matrix{rows: rows, cols: cols, rowMajor: true, data: data}
	SanityCheck(m, file)
	return m
}

// SanityCheck runs checks to make sure our Matrix has the right number
// of elements, etc.
func SanityCheck(m Matrix, name string) {
	// length of m.data should be rows * cols
	length := m.rows * m.cols
	if length != len(m.data) {
		log.Printf("Matrix %s doesn't appear to be valid\n", name)
		os.Exit(1)
	}
}

// WriteMatrix writes the provided Matrix to a csv file.
func WriteMatrix(m Matrix, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println("", err)
		log.Println("I couldn't write the file because of the above error.")
		fmt.Println("The result Matrix looks like:")
		PrintMatrix(m)
		os.Exit(1)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	i := 0
	for r := 0; r < m.rows; r++ {
		row := make([]string, m.cols)
		for c := 0; c < m.cols; c++ {
			row[c] = strconv.Itoa(m.data[i])
			i++
		}
		err = writer.Write(row)
		if err != nil {
			log.Fatal("Cannot write to file:", err)
		}
	}
}

// Morph is a silly easter egg.
func Morph() {
	lines := make([]string, 10)
	ascii := make([]string, 14)
	lines[0] = "What is real? How do you define 'real'?"
	ascii[0] = "         ;;;;;;;;;,    "
	lines[1] = "The body cannot live without the mind"
	ascii[4] = "    ;';;@@@@@'@@@@@;;''  "
	ascii[7] = "     ';;;;;;;;;;;;;;;' "
	lines[2] = "Unfortunately, no one can be told what the Matrix is. You have to see it for yourself."
	ascii[8] = "      ;;;;;''''';;;;;  "
	lines[3] = "Fate, it seems, is not without a sense of irony."
	ascii[11] = "       ++;;';;;;;;++   "
	lines[4] = "You've felt it your entire life, that there's something wrong with the world."
	ascii[1] = "       :;;;;;;;;;;;; "
	lines[5] = "You take the blue pill - the story ends, you wake up in your bed and believe whatever you want to believe. You take the red pill - you stay in Wonderland and I show you how deep the rabbit-hole goes."
	ascii[2] = "      `;;;;;;;;;;;;;;  "
	ascii[5] = "    ,';;;@@@+;;@@@#;;''   "
	lines[6] = "You can feel it when you go to work... when you go to church... when you pay your taxes."
	ascii[13] = "     '+++;;;;';;;;++++"
	ascii[10] = "        +;;;;;;;;;+:"
	lines[7] = "Have you ever had a dream... that you were so sure was real?"
	ascii[9] = "       ;;;;;;;;;;;;;;"
	lines[8] = "They will never be as strong, or as fast, as *you* can be."
	ascii[3] = "    '+;''@@@#;;@@@@';''"
	ascii[12] = "      +++;;;''';;;+++ "
	lines[9] = "What you know, you can't explain. But you feel it."
	ascii[6] = "     ';;;;;;;;;;;;;;;';"
	for i := 0; i < len(ascii); i++ {
		fmt.Println(ascii[i])
	}
	rand.Seed(time.Now().Unix())
	fmt.Println(lines[rand.Intn(len(lines))])
}

func main() {
	var nameptr1 = flag.String("a", "", "csv file containing Matrix a")
	var nameptr2 = flag.String("b", "", "csv file containing Matrix b")
	var output = flag.String("o", "", "output file (Optional)")
	var morph = flag.Bool("morph", false, "Red Pill?")
	var hal = flag.Bool("hal", false, "Open the pod bay doors")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Computes and returns the dot product of two matrices\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *morph {
		Morph()
	}
	if *hal {
		fmt.Println("I'm sorry, Dave. I'm afraid I can't do that.")
	}

	var filename1 = *nameptr1
	var filename2 = *nameptr2
	if filename1 == "" || filename2 == "" {
		fmt.Println("I need you to specifiy -a and -b in order to do anything useful.")
		fmt.Println("Each of these should point to a file that is a csv representation of a Matrix")
		fmt.Println("where one line of the csv is one row of the Matrix.")
		fmt.Println("I will then compute c = a*b and give you c.")
		os.Exit(1)
	}
	var b = ProcessFile(filename1)
	var a = ProcessFile(filename2)
	var m = Product(b, a)

	if *output == "" {
		fmt.Println("Matrix b looks like:")
		PrintMatrix(b)
		fmt.Println("Matrix a looks like:")
		PrintMatrix(a)
		fmt.Println("Result looks like: ")
		PrintMatrix(m)
	} else {
		WriteMatrix(m, *output)
	}

}
