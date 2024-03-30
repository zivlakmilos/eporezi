package scard

import (
	"crypto/x509"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/go-pkcs11/pkcs11"
)

type SCard struct {
	modulePath string
	module     *pkcs11.Module
	slot       *pkcs11.Slot
	info       *pkcs11.SlotInfo
}

func NewSCard(module string) *SCard {
	return &SCard{
		modulePath: module,
	}
}

func (s *SCard) Open() error {
	var err error
	defer func() {
		if err != nil {
			s.Close()
		}
	}()

	s.module, err = pkcs11.Open(s.modulePath)
	if err != nil {
		return err
	}

	return nil
}

func (s *SCard) Close() error {
	s.Disconnect()

	if s.module != nil {
		err := s.module.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SCard) Connect(id uint32, pin string) error {
	var err error
	s.slot, err = s.module.Slot(id, pkcs11.Options{
		PIN: pin,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *SCard) Disconnect() error {
	if s.slot != nil {
		err := s.slot.Close()
		if err != nil {
			return err
		}

		s.slot = nil
	}

	return nil
}

func (s *SCard) ListCards() ([]*Info, error) {
	cards := []*Info{}

	ids, err := s.module.SlotIDs()
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		info, err := s.module.SlotInfo(id)
		if err != nil {
			return nil, err
		}

		if info.Serial != "" {
			cards = append(cards, &Info{
				Id:           id,
				Label:        info.Label,
				SerialNumber: info.Serial,
			})
		}
	}

	return cards, nil
}

func (s *SCard) SubjectInfo() (*SubjectInfo, error) {
	cert, err := s.getCertificate()
	if err != nil {
		return nil, err
	}

	name := strings.Split(cert.Subject.CommonName, " ")
	if len(name) < 2 {
		return nil, fmt.Errorf("error: cen't parse name from certificate")
	}

	r, err := regexp.Compile("[0-9]{13}")
	if err != nil {
		panic("error: compile regexp")
	}
	personalId := r.FindString(cert.Subject.SerialNumber)
	if personalId == "" {
		return nil, fmt.Errorf("error: can't parse personal id")
	}

	subject := &SubjectInfo{
		Name:       name[0],
		Surname:    name[1],
		PersonalId: personalId,
	}
	return subject, nil
}

func (s *SCard) getCertificate() (*x509.Certificate, error) {
	objs, err := s.slot.Objects(pkcs11.Filter{
		Class: pkcs11.ClassCertificate,
	})
	if err != nil {
		return nil, err
	}

	if len(objs) == 0 {
		return nil, fmt.Errorf("error: certifiace not found")
	}

	cert, err := objs[0].Certificate()
	if err != nil {
		return nil, err
	}

	x509Cert, err := cert.X509()
	if err != nil {
		return nil, err
	}

	return x509Cert, nil
}
