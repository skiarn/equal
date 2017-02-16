package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	flags := flags()
	fv1 := flags[0]
	fv2 := flags[1]

	var v1 interface{}
	var v2 interface{}

	//try unmarshal json
	if err := json.Unmarshal([]byte(fv1), &v1); err == nil {
		if err := json.Unmarshal([]byte(fv2), &v2); err == nil {
			if !reflect.DeepEqual(v1, v2) {
				fail()
			}
			success()
		}
	}

	//try unmarshal xml
	if err := xml.Unmarshal([]byte(fv1), &v1); err == nil {
		if err := xml.Unmarshal([]byte(fv2), &v2); err == nil {
			if !reflect.DeepEqual(v1, v2) {
				fail()
			}
			success()
		}
	}

	//try string compare
	if string(fv1) != string(fv2) {
		fail()
	}
	success()

}

func flags() []string {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: ./equal ´[ {name: 'Sweden', code: 'SE'}, {name: 'Trump', code: 'WALL'}, {name: 'Mexico', code: 'MX'} ]´ ´[ {name: 'Sweden', code: 'SE'}, {name: 'Trump', code: 'WALL'}, {name: 'Mexico', code: 'MX'} ]´\n")
	}
	flag.Parse()
	flags := strings.Split(strings.Join(os.Args[1:], ""), "´")
	flags = nRemove(flags, "")

	if len(flags) != 2 {
		fmt.Printf("Got %v values %s, expected two.\n", len(flags), flags)
		fail()
	}
	return flags
}

func fail() {
	if _, err := fmt.Fprintln(os.Stdout, "false"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func success() {
	os.Exit(0)
}
func nRemove(slice []string, s string) []string {
	var r []string
	for _, i := range slice {
		if i == "" {
			continue
		}
		r = append(r, i)
	}
	return r
}
