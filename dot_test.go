
package main
import (
	"testing"
	"reflect"
	"os/exec"
	"math/rand"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Test2x2dot2x2(t *testing.T) {
	m1 := matrix{2, 2, true, []int{1, 2, 3, 4}}
	m2 := matrix{2, 2, true, []int{3, 4, 5, 6}}
	good := matrix{2, 2, true, []int{13,16,29,36}}
	product := Product(m1, m2)
	if !reflect.DeepEqual(product, good) {
		t.Error("2x2dot2x2 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}

func Test2x4dot4x2(t *testing.T) {
	m1 := matrix{2, 4, true, []int{3, 4, 5, 6, 7, 8, 9, 10}}
	m2 := matrix{4, 2, true, []int{1, 2, 3, 4, 5, 6, 7, 8}}
	good := matrix{2, 2, true, []int{82, 100, 146, 180}}
	product := Product(m1, m2)
	if !reflect.DeepEqual(product, good) {
		t.Error("2x4dot4x2 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}

func Test1x9dot9x1(t *testing.T) {
	good := matrix{1,1, true, []int{237}}
	m1 := ProcessFile("testdata/1x9.csv")
	m2 := ProcessFile("testdata/9x1.csv")
	product := Product(m1, m2)
	if !reflect.DeepEqual(product, good) {
		t.Error("4x2dot2x2 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}

func Test9x1dot1x9(t *testing.T) {
	good := ProcessFile("testdata/9x9_out.csv")
	m2 := ProcessFile("testdata/1x9.csv")
	m1 := ProcessFile("testdata/9x1.csv")
	product := Product(m1, m2)
	if !reflect.DeepEqual(product, good) {
		t.Error("4x2dot2x2 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}

func Test4x4dot4x4andInFiles(t *testing.T) {
	m1 := ProcessFile("testdata/4x4_1.csv")
	m2 := ProcessFile("testdata/4x4_2.csv")
	good := ProcessFile("testdata/4x4_out.csv")
	product := Product(m1, m2)
	if !reflect.DeepEqual(product, good) {
		t.Error("4x4dot4x4 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}

func Test2x2dot2x4andInFiles(t *testing.T) {
	a := ProcessFile("testdata/2x2.csv")
	b := ProcessFile("testdata/2x4.csv")
	good := ProcessFile("testdata/2x4_out.csv")
	product := Product(a, b)
	if !reflect.DeepEqual(product, good) {
		t.Error("2x2dot2x4 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}

func Test1x1dot1x1andInFiles(t *testing.T) {
	a := ProcessFile("testdata/1x1_1.csv")
	b := ProcessFile("testdata/1x1_2.csv")
	good := matrix{1, 1, true, []int{15} }
	product := Product(a, b)
	if !reflect.DeepEqual(product, good) {
		t.Error("2x2dot2x4 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}

func Test2x3dot3x2andInFiles(t *testing.T) {
	a := ProcessFile("testdata/2x3.csv")
	b := ProcessFile("testdata/3x2.csv")
	good := matrix{2, 2, true, []int{58, 64, 139, 154} }
	product := Product(a, b)
	if !reflect.DeepEqual(product, good) {
		t.Error("2x2dot2x4 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}

func Test1x3dot3x4(t *testing.T) {
	a := matrix{1, 3, true, []int{3, 4, 2}}
	b := matrix{3, 4, true, []int{13, 9, 7, 15, 8, 7, 4, 6, 6, 4, 0, 3}}
	good := matrix{1, 4, true, []int{83, 63, 37, 75}}
	product := Product(a, b)
	if !reflect.DeepEqual(product, good) {
		t.Error("2x2dot2x4 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}

func TestConvert4(t *testing.T) {
	m_r := matrix{2, 2, true, []int{3, 4, 5, 6}}
	m_c := matrix{2, 2, false, []int{3, 5, 4, 6}}
	m2_c := Convert(m_r)
	m2_r := Convert(m_c)
	if !reflect.DeepEqual(m_r, m2_r) {
		t.Error("Converting column to row failed")
		t.Error("Expected:")
		PrintMatrix(m_r)
		t.Error("Got:")
		PrintMatrix(m2_r)
	}
	if !reflect.DeepEqual(m_c, m2_c) {
		t.Error("Converting row to column failed")
		t.Error("Expected:")
		PrintMatrix(m_c)
		t.Error("Got:")
		PrintMatrix(m2_c)
	}
}

func Test4x4dot2x2andInFiles(t *testing.T) {
	cmd := exec.Command("go", "run dot.go -f1 testdata/4x4_1.csv -f2 testdata/2x2.csv")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestOutFile(t *testing.T) {
	rand.Seed(42)
	someNum := rand.Int()
	filename := fmt.Sprintf("/tmp/%s.csv", strconv.Itoa(someNum))
	cmd := exec.Command("go", "run", "dot.go", "-a", "testdata/2x2.csv", "-b", "testdata/2x4.csv", "-o", filename)
	err := cmd.Run()
	if err != nil {
		log.Println("", err)
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// path/to/whatever does not exist
		t.Fatalf("File creation failed for %s", filename)
	}
}
