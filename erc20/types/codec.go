package types

import (
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
    ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    cdc.RegisterConcrete(MsgRegisterTokenPair{}, "erc20/MsgRegisterTokenPair", nil)
    cdc.RegisterConcrete(MsgConvertCoinToERC20{}, "erc20/MsgConvertCoinToERC20", nil)
    cdc.RegisterConcrete(MsgConvertERC20ToCoin{}, "erc20/MsgConvertERC20ToCoin", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
    registry.RegisterImplementations(
        (*sdk.Msg)(nil),
        &MsgRegisterTokenPair{},
        &MsgConvertCoinToERC20{},
        &MsgConvertERC20ToCoin{},
    )
}
