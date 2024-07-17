package types

const (
    ModuleName = "erc20"
    StoreKey   = ModuleName
    RouterKey  = ModuleName
)

var (
    TokenPairKeyPrefix = []byte{0x01}
)

func TokenPairKey(denom string) []byte {
    return append(TokenPairKeyPrefix, []byte(denom)...)
}
