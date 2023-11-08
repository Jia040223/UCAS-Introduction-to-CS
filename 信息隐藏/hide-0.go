package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	S = 54 // standard size of bmp headers
	T = 32 // number of bytes needed to hide the text length
	C = 4  // number of bytes needed to hide a character
)

// modify hides an integer to a byte slice
func modify(value int, pix []byte, size int) {
	for i := 0; i < size; i++ {
		pix[i]=pix[i] & 0xFC
		var mid_value byte=byte(value & 0x03)
		pix[i]=pix[i] | mid_value
		value=value >> 2
		// TODO: write your code here
		// replace last 2 bits of pix[i] with the last 2 bits of value
		// the next iteration repeats with the next 2 bits of value
	}
}

var (
	srcImage  string // input image name
	srcTxt    string // input text name
	destImage string // output doctored image name
)

// init sets command line arguments
func init() {
	// DON'T modify this function!!!
	flag.StringVar(&srcImage, "i", "", "input image name")
	flag.StringVar(&srcTxt, "t", "", "input text name")
	flag.StringVar(&destImage, "d", "", "output doctored image name")
}

func main() {
	// parse command line arguments
	flag.Parse()
	if srcImage == "" || srcTxt == "" || destImage == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	// read input image to a byte slice p
	p, err := ioutil.ReadFile(srcImage)
	if err != nil {
		fmt.Printf("Read image file failed, err = %v\n", err)
		os.Exit(1)
	}
	// read input text to a byte slice t
	t, err := ioutil.ReadFile(srcTxt)
	if err != nil {
		fmt.Printf("Read text file failed, err = %v\n", err)
		os.Exit(1)
	}
	// check if the text is too big
	if T+len(t)*C > len(p[S:]) {
		fmt.Println("The text file is too big")
		os.Exit(1)
	}
	// save the text length to p
	modify(len(t), p[S:S+T], T)
	// save the content of text to p
	for i := 0; i < len(t); i++ {
		offset := S + T + C*i
		modify(int(t[i]), p[offset:offset+C], C)
	}

	// write the modified p to destImage
	err = ioutil.WriteFile(destImage, p, 0666)
	if err != nil {
		fmt.Printf("Write doctored image failed, err = %v\n", err)
		os.Exit(1)
	}
}
