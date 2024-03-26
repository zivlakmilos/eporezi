package scard

import (
	"fmt"

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

func (s *SCard) Open(pin string) error {
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

	ids, err := s.module.SlotIDs()
	for _, id := range ids {
		info, _ := s.module.SlotInfo(id)
		if info.Serial != "" {
			s.info = info
			s.slot, err = s.module.Slot(id, pkcs11.Options{
				PIN: pin,
			})
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}

func (s *SCard) Close() error {
	if s.slot != nil {
		err := s.slot.Close()
		if err != nil {
			return err
		}
	}

	if s.module != nil {
		err := s.module.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SCard) Info() (*Info, error) {
	if s.info == nil {
		return nil, fmt.Errorf("error: slot not present")
	}

	return &Info{
		Label:        s.info.Label,
		SerialNumber: s.info.Serial,
	}, nil
}
