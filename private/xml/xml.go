package xml

import (
	"fmt"

	"github.com/beevik/etree"
)

type Xml struct {
	doc *etree.Document
}

func NewXml() *Xml {
	return &Xml{}
}

func (x *Xml) Parse(xml string) error {
	doc := etree.NewDocument()
	err := doc.ReadFromString(xml)
	if err != nil {
		return err
	}

	x.doc = doc

	return nil
}

func (x *Xml) AddEnvelope() error {
	sigEl := x.doc.FindElement("//signatures")
	if sigEl == nil {
		return fmt.Errorf("error: cannot find signatures tag")
	}

	sigEl.AddChild(createSignature())

	return nil
}

func (x *Xml) Canonicalize() (string, error) {
	xml, err := x.doc.WriteToString()
	if err != nil {
		return "", err
	}

	return xml, nil
}
