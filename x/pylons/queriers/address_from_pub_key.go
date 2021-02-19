package queriers

import (
	"encoding/hex"

	"github.com/Pylons-tech/pylons/x/pylons/keep"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
	crypto "github.com/tendermint/tendermint/crypto/secp256k1"
)

// query endpoints supported by the nameservice Querier
const (
	KeyAddrFromPubKey = "addr_from_pub_key"
)

// AddrFromPubKey returns a bech32 public address from the public key
func AddrFromPubKey(ctx sdk.Context, path []string, req abci.RequestQuery, keeper keep.Keeper) ([]byte, error) {

	if len(path) < 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The hex pub key not provided")
	}
	hexPubKey := path[0]

	pubKeyBytes, err := hex.DecodeString(hexPubKey)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	var pubKeyBytes33 [33]byte
	copy(pubKeyBytes33[:], pubKeyBytes)
	var pubKey = crypto.PubKey(pubKeyBytes33[:])

	addrResp := AddrResp{
		Bech32Addr: sdk.AccAddress(pubKey.Address().Bytes()).String(),
	}
	// if we cannot find the value then it should return an error
	bz, err := keeper.Cdc.MarshalJSON(addrResp)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return bz, nil

}

// AddrResp holds the bech32 encoded address
type AddrResp struct {
	Bech32Addr string
}
