package xml

import (
	"github.com/beevik/etree"
)

func createSignature() *etree.Element {
	signature := etree.NewElement("Signature")
	signature.CreateAttr("xmlns", "http://www.w3.org/2000/09/xmldsig#")

	signedInfo := createSignedInfo()
	signature.AddChild(signedInfo)

	return signature
}

func createSignedInfo() *etree.Element {
	signedInfo := etree.NewElement("SignedInfo")

	signedInfo.AddChild(createCanonicalizationMethod())
	signedInfo.AddChild(createSignatureMethod())
	signedInfo.AddChild(createReference())

	return signedInfo
}

func createCanonicalizationMethod() *etree.Element {
	canonMethod := etree.NewElement("CanonicalizationMethod")
	canonMethod.CreateAttr("Algorithm", "http://www.w3.org/2001/10/xml-exc-c14n#")

	return canonMethod
}

func createSignatureMethod() *etree.Element {
	canonMethod := etree.NewElement("SignatureMethod")
	canonMethod.CreateAttr("Algorithm", "http://www.w3.org/2000/09/xmldsig#rsa-sha1")

	return canonMethod
}

func createReference() *etree.Element {
	reference := etree.NewElement("Reference")
	reference.CreateAttr("URI", "")

	reference.AddChild(createTransforms())

	reference.AddChild(createDigestMethod())
	reference.AddChild(etree.NewElement("DigestValue"))

	return reference
}

func createTransforms() *etree.Element {
	transforms := etree.NewElement("Transforms")

	transAlgo := etree.NewElement("Transform")
	transAlgo.CreateAttr("Algorithm", "http://www.w3.org/2000/09/xmldsig#enveloped-signature")
	transforms.AddChild(transAlgo)

	return transforms
}

func createDigestMethod() *etree.Element {
	digestMethod := etree.NewElement("DigestMethod")
	digestMethod.CreateAttr("Algorithm", "http://www.w3.org/2000/09/xmldsig#sha1")

	return digestMethod
}
