package group

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (p *Proposal) GetMsgs() ([]sdk.Msg, error) {
	return GetMsgs(p.Messages, "proposal")
}

func (p *Proposal) SetMsgs(msgs []sdk.Msg) error {
	anys, err := SetMsgs(msgs)
	if err != nil {
		return err
	}
	p.Messages = anys
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	return UnpackInterfaces(unpacker, p.Messages)
}
