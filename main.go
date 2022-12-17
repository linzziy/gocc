package main

import (
	"bufio"
	"flag"
	"fmt"
	"gocc/goutil"
	"log"
	"os"

	"github.com/longbridgeapp/opencc"
)

var (
	input  = flag.String("i", "", "file of original text to read")
	output = flag.String("o", "", "file of converted text to write")
	config = flag.String("c", "", "convert config, s2t, t2s, etc")
)

func main() {
	flag.Parse()
	var err error
	var in, out *os.File //io.Reader
	if *input == "" {
		in = os.Stdin
	} else {
		in, err = os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		defer in.Close()
	}
	br := bufio.NewReader(in)

	if *output == "" {
		out = os.Stdout
	} else {
		out, err = os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
	}

	if *config == "" {
		*config = "s2t"
	}

	s2t, err := opencc.New(*config)
	if err != nil {
		log.Fatal(err)
	}

	err = goutil.ForEachLine(br, func(line string) error {
		str, e := s2t.Convert(line)
		if e != nil {
			return e
		}
		fmt.Fprint(out, str+"\n")
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	////in := `自然语言处理是人工智能领域中的一个重要方向。`
	//out, err := s2t.Convert(*input)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s\n%s\n", in, out)
	////自然语言处理是人工智能领域中的一个重要方向。
	////自然語言處理是人工智能領域中的一個重要方向。
}
