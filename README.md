# dotmatrix
A simple go program by Mike Thompson

Package dot computes the dot product of two matrices. It can be run interactively on the console or scripted using csv files for input and output.

Specifying Matrices as csv files is quite easy and works exactly how you'd expect.

A matrix that looks like:

```
  [ 2, 3 ]
  [ 4, 5 ]
```
Should be represented in a csv file simply as:
```
2,3
4,5
```
###Usage
Usage is quite straight forward, you must specify matrices as files, but you can receive output to the console or to another csv file (using the -o flag).  Relevant usage:
```
[mt@1Q84 dotmatrix]$ go run dot.go -h
Computes and returns the dot product of two matrices
  -a string
    	csv file containing matrix a
  -b string
    	csv file containing matrix b
  -o string
    	output file (Optional)

```
There are a couple of fun easter eggs in there, because why not?
###Example files
There are numerous examples of good and bad matrix files in the testdata/ directory.  Generally bad files should have the word "bad" in the filename.