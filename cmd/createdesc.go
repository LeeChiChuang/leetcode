package main

import (
	"flag"
	"fmt"
	"github.com/leechichuang/leetcodetools/generater"
	"os"
)

func getPathName(q string) string {
	dir, _ := os.Getwd()
	dir = fmt.Sprintf("%s/../algorithm", dir)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}

	return dir
}

func main() {
	q := flag.String("name", "foo", "question name")
	flag.Parse()

	l, err := generater.NewLeetCode(*q)
	if err != nil {
		panic(err)
	}

	err = l.WriteCode(getPathName(*q))
	if err != nil {
		panic(err)
	}

	err = l.WriteDesc(getPathName(*q))
	if err != nil {
		panic(err)
	}

	fmt.Println("generate success")
}
