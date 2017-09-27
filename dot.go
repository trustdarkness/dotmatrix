// Package dot computes the dot product of two matrices.
// It can be run interactively on the console or scripted using
// csv files for input and output.
package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/csv"
	"math/rand"
	"io"
	"strconv"
	"log"
	"time"
)

type matrix struct {
	rows int
	cols int
	row_major bool
    data []int
}

// Product computes the Product of two matrices and returns a matrix object.
// If the matrices cannot be multiplies, an error is printed.
func Product(a matrix, b matrix) (matrix) {
	if (a.cols != b.rows) {
		log.Println("The dot product of these matrices is not defined,")
		log.Println("see https://en.wikipedia.org/wiki/Dot_product for more information.")
		os.Exit(0)
	}

	// m1 should be row major and m2 should be column major.
	if !a.row_major {
		a = Convert(a)
	}

	if b.row_major {
		b = Convert(b)
	}

	new_len := a.rows*b.cols
	result_data := make([]int, new_len)

	dstPos := 0
	bPos := 0
	for dstYPos := 0; dstYPos < a.rows; {
		//for dstXPos := 0; dstXPos < b.rows; dstXPos++ {*/
			//for bcol := 0; bcol < b.cols; bcol++ {

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
						result_data[dstPos] += term1 * term2
						bPos++
					}
				}
				dstPos++
			//}

		//}
	}

	var result = matrix{rows: a.rows, cols: b.cols, row_major: true, data: result_data}
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
	for r := 0; r < m.rows; r++ {
		fmt.Printf("[ ")
		for c := 0; c < m.cols; c++ {
		  fmt.Printf("%2d ", m.data[i])
		  i++
		}
		fmt.Printf("]\n")
	}
}

// Convert converts a matrix between row_major and column major.
func Convert(m matrix) (matrix) {
	var result_data = make([]int,len(m.data))
	dstPos := 0
	for col := 0; col < m.cols; col++ {
		srcPos := 0
		for row := 0; row < m.rows; row++ {
			srcPos = row * m.cols + col
			result_data[dstPos] = m.data[srcPos]
			dstPos++
		}
	}
	var ret = matrix{rows:m.rows, cols: m.cols, row_major: !m.row_major, data: result_data}
	return ret
}

// ProcessFile takes a user provided file name and attempts to build a matrix object from it.
func ProcessFile(file string) matrix {
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
func Morph(){
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
	var nameptr1 = flag.String("a", "", "csv file containing matrix a")
	var nameptr2 = flag.String("b", "", "csv file containing matrix b")
	var output = flag.String("o", "", "output file (Optional)")
	var morph = flag.Bool("morph", false, "Red Pill?")
	var hal = flag.Bool("hal", false, "Open the pod bay doors")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Computes the dot product of two matrices c = b*a and returns c.\n")
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
		fmt.Println("Each of these should point to a file that is a csv representation of a matrix")
		fmt.Println("where one line of the csv is one row of the matrix.")
		fmt.Println("I will then compute c = ba and give you c.")
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


