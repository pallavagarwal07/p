package p

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type File struct {
	f *os.File
}

type String string
type Strings []string

func Open(name string) File {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return File{f}
}

func (f File) ReadLines() Strings {
	return f.Read().Split("\n")
}

func (f File) Read() String {
	data, err := ioutil.ReadAll(f.f)
	if err != nil {
		panic(err)
	}
	return String(data)
}

func (s String) Split(sep string) Strings {
	return Strings(strings.Split(string(s), sep))
}

func (s String) Strip() String {
	return String(strings.TrimSpace(string(s)))
}

func Int(s String) int {
	n, err := strconv.Atoi(string(s))
	if err != nil {
		panic(err)
	}
	return n
}

func (s Strings) Map(f func(String) String) Strings {
	arr := make(Strings, len(s))
	for i, str := range s {
		arr[i] = string(f(String(str)))
	}
	return arr
}

func (s Strings) MapInt(f func(String) int) []int {
	arr := make([]int, len(s))
	for i, str := range s {
		arr[i] = f(String(str))
	}
	return arr
}

func (s Strings) MapFloat64(f func(String) float64) []float64 {
	arr := make([]float64, len(s))
	for i, str := range s {
		arr[i] = f(String(str))
	}
	return arr
}

func (s Strings) Filter(f func(String) bool) Strings {
	var out Strings
	for _, str := range s {
		if f(String(str)) {
			out = append(out, str)
		}
	}
	return out
}
