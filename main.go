package main

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

var exp = 0

var cases = []struct {
	name   string
	layout string
	value  string
}{
	{"invalid layout and value", "xxx", "xxx"},
	{"invalid layout", "xxx", "2013-Feb-03"},
	{"invalid value", time.RFC3339, "2013-Feb-03"},
}

func main() {

	for _, tt := range cases {
		fmt.Println("****************************")
		t, err := parseCause(tt.layout, tt.value)
		if err != nil {
			fmt.Printf("%+v", err)
		}
		fmt.Println(t)
	}
}

func parseFmt(layout, value string) (*time.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &t, nil
}

func parseWithStack(layout, value string) (*time.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &t, nil
}

func parseWrap(layout, value string) (*time.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse")
	}
	return &t, nil
}

func parseCause(layout, value string) (*time.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return nil, errors.Cause(err)
	}
	return &t, nil
}
