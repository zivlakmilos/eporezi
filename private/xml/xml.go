package xml

import (
	"crypto/sha1"
	"encoding/base64"
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

func (x *Xml) AddEnvelope(ref string) error {
	sigEl := x.doc.FindElement("//signatures")
	if sigEl == nil {
		return fmt.Errorf("error: cannot find signatures tag")
	}

	sigEl.AddChild(createSignature())

	err := x.SetReference("123")
	if err != nil {
		return err
	}

	err = x.SetDigest("123")
	if err != nil {
		return err
	}

	return nil
}

func (x *Xml) SetDigest(ref string) error {
	el := x.doc.FindElement(fmt.Sprintf("//*[@id='%s']", ref))
	if el == nil {
		return fmt.Errorf("error: cannot find reference tag")
	}

	digestEl := x.doc.FindElement("//Signature/SignedInfo/Reference/DigestValue")
	if digestEl == nil {
		return fmt.Errorf("error: cannot find DigestValue tag")
	}

	doc := etree.NewDocumentWithRoot(el.Copy())
	doc.WriteSettings.CanonicalEndTags = true
	doc.WriteSettings.CanonicalText = true
	doc.WriteSettings.CanonicalAttrVal = true

	str, err := doc.WriteToString()
	if err != nil {
		return err
	}

	sha := sha1.Sum([]byte(str))
	digest := base64.StdEncoding.EncodeToString(sha[:])

	digestEl.SetText(digest)

	return nil
}

func (x *Xml) SetReference(uri string) error {
	el := x.doc.FindElement("//Signature/SignedInfo/Reference")
	if el == nil {
		return fmt.Errorf("error: Reference element not found")
	}

	el.CreateAttr("URI", uri)

	return nil
}

func (x *Xml) String() string {
	str, _ := x.doc.WriteToString()
	return str
}
