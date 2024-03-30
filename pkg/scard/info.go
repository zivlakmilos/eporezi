package scard

import "fmt"

type Info struct {
	Id           uint32
	Label        string
	SerialNumber string
}

func (i Info) String() string {
	return fmt.Sprintf("{ Id: %d, Label: %s, SerialNumber: %s, }", i.Id, i.Label, i.SerialNumber)
}
