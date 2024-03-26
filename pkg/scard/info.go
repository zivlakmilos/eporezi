package scard

import "fmt"

type Info struct {
	Label        string
	SerialNumber string
}

func (i Info) String() string {
	return fmt.Sprintf("{ Label: %s, SerialNumber: %s, }\n", i.Label, i.SerialNumber)
}
