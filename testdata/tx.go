package testdata

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// KeyTestPubAddr generates a new secp256k1 keypair.
func KeyTestPubAddr() (cryptotypes.PrivKey, cryptotypes.PubKey, sdk.AccAddress) {
	key := secp256k1.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}

// NewTestMsg creates a message for testing with the given signers.
func NewTestMsg(addrs ...sdk.AccAddress) *TestMsg {
	var accAddresses []string

	for _, addr := range addrs {
		accAddresses = append(accAddresses, addr.String())
	}

	return &TestMsg{
		Signers: accAddresses,
	}
}

var _ sdk.Msg = (*TestMsg)(nil)

func (msg *TestMsg) Route() string { return "TestMsg" }
func (msg *TestMsg) Type() string  { return "Test message" }
func (msg *TestMsg) GetSignBytes() []byte {
	bz, err := json.Marshal(msg.Signers)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (msg *TestMsg) GetSigners() []sdk.AccAddress {
	signers := make([]sdk.AccAddress, 0, len(msg.Signers))
	for _, addr := range msg.Signers {
		a, _ := sdk.AccAddressFromBech32(addr)
		signers = append(signers, a)
	}
	return signers
}

func (msg *TestMsg) ValidateBasic() error {
	for _, addr := range msg.Signers {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return sdkerrors.ErrInvalidAddress.Wrapf("invalid signer address: %s", err)
		}
	}
	return nil
}
