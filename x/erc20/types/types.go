package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/ethereum/go-ethereum/common"
)

type TokenPair struct {
    ERC20Address common.Address `json:"erc20_address"`
    Denom        string         `json:"denom"`
    Enabled      bool           `json:"enabled"`
}

type MsgRegisterTokenPair struct {
    Creator      sdk.AccAddress `json:"creator"`
    ERC20Address common.Address `json:"erc20_address"`
    Denom        string         `json:"denom"`
}

func NewMsgRegisterTokenPair(creator sdk.AccAddress, erc20Address common.Address, denom string) MsgRegisterTokenPair {
    return MsgRegisterTokenPair{
        Creator:      creator,
        ERC20Address: erc20Address,
        Denom:        denom,
    }
}

// Implement ValidateBasic and GetSignBytes for MsgRegisterTokenPair

type MsgConvertCoinToERC20 struct {
    Creator sdk.AccAddress `json:"creator"`
    Amount  sdk.Coin       `json:"amount"`
}

func NewMsgConvertCoinToERC20(creator sdk.AccAddress, amount sdk.Coin) MsgConvertCoinToERC20 {
    return MsgConvertCoinToERC20{
        Creator: creator,
        Amount:  amount,
    }
}

// Implement ValidateBasic and GetSignBytes for MsgConvertCoinToERC20

type MsgConvertERC20ToCoin struct {
    Creator sdk.AccAddress `json:"creator"`
    Amount  sdk.Int        `json:"amount"`
}

func NewMsgConvertERC20ToCoin(creator sdk.AccAddress, amount sdk.Int) MsgConvertERC20ToCoin {
    return MsgConvertERC20ToCoin{
        Creator: creator,
        Amount:  amount,
    }
}

// Implement ValidateBasic and GetSignBytes for MsgConvertERC20ToCoin
