package scard

import (
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
				Label:        info.Label,
				SerialNumber: info.Serial,
			})
		}
	}

	return cards, nil
}
