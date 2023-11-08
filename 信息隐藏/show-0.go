package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	S = 54 // standard size of header
	T = 32 // number of bytes needed to hide the text length
	C = 4  // number of bytes needed to hide a character
)

var (
	image string // input doctor image name
	txt   string // output text name
)

// init sets command line arguments
func init() {
	// DON'T modify this function!!!
	flag.StringVar(&image, "i", "", "input image name")
	flag.StringVar(&txt, "t", "", "output text name")
}

func main() {
	// parse command line arguments
	flag.Parse()
	if image == "" || txt == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	// TODO: write your code here
	p, err := ioutil.ReadFile(image)
	if err != nil {
		fmt.Printf("Read image file failed, err = %v\n", err)
		os.Exit(1)
	}
	// read input text to a byte slice t
	
	var length int=0
	for i := 0; i < T-1; i++ {
		mid_value:=int(p[S+T-i-1]) & 0x03
		length=length | mid_value
		length=length<<2
	}
	mid_value:=int(p[S]) & 0x03
	length=length | mid_value

	fmt.Println(length)
	var t []byte
	for i:=0;i<length;i++{
		offset := S + T + C*i
		var value byte=0
		for j:=0;j<C-1;j++{
			value=value | (p[offset+C-j-1] & 0x03)
			value=value<<2
		}
		value=value | (p[offset] & 0x03)
		t=append(t,value)
	}

	err = ioutil.WriteFile(txt, t, 0666)
	if err != nil {
		fmt.Printf("Write doctored image failed, err = %v\n", err)
		os.Exit(1)
	}
}
