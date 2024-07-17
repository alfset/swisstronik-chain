package keeper

import (
    "github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/username/evmos-hub/x/erc20/types"
)

type Keeper struct {
    storeKey sdk.StoreKey
    cdc      codec.BinaryCodec
}

func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey) Keeper {
    return Keeper{
        storeKey: storeKey,
        cdc:      cdc,
    }
}

func (k Keeper) RegisterTokenPair(ctx sdk.Context, pair types.TokenPair) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TokenPairKeyPrefix)
    bz := k.cdc.MustMarshal(&pair)
    store.Set(types.TokenPairKey(pair.Denom), bz)
}

func (k Keeper) GetTokenPair(ctx sdk.Context, denom string) (types.TokenPair, bool) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TokenPairKeyPrefix)
    bz := store.Get(types.TokenPairKey(denom))
    if bz == nil {
        return types.TokenPair{}, false
    }
    var pair types.TokenPair
    k.cdc.MustUnmarshal(bz, &pair)
    return pair, true
}

func (k Keeper) ConvertCoinToERC20(ctx sdk.Context, creator sdk.AccAddress, amount sdk.Coin) error {
    pair, found := k.GetTokenPair(ctx, amount.Denom)
    if !found || !pair.Enabled {
        return types.ErrTokenPairNotFound
    }

    // Implement the conversion logic from Coin to ERC20
    // ...
    return nil
}

func (k Keeper) ConvertERC20ToCoin(ctx sdk.Context, creator sdk.AccAddress, amount sdk.Int, denom string) error {
    pair, found := k.GetTokenPair(ctx, denom)
    if !found || !pair.Enabled {
        return types.ErrTokenPairNotFound
    }

    // Implement the conversion logic from ERC20 to Coin
    // ...
    return nil
}
