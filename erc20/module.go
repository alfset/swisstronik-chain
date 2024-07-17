package erc20

import (
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/types/module"
    "github.com/cosmos/cosmos-sdk/x/erc20/keeper"
    "github.com/cosmos/cosmos-sdk/x/erc20/types"
    "github.com/spf13/cobra"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
    _ module.AppModule      = AppModule{}
    _ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string { return types.ModuleName }

func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
    types.RegisterCodec(cdc)
}

func (AppModuleBasic) DefaultGenesis() json.RawMessage {
    return types.ModuleCdc.MustMarshalJSON(types.DefaultGenesis())
}

func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
    var data types.GenesisState
    if err := types.ModuleCdc.UnmarshalJSON(bz, &data); err != nil {
        return err
    }
    return types.ValidateGenesis(data)
}

type AppModule struct {
    AppModuleBasic
    keeper keeper.Keeper
}

func NewAppModule(k keeper.Keeper) AppModule {
    return AppModule{
        AppModuleBasic: AppModuleBasic{},
        keeper:         k,
    }
}

func (AppModule) Name() string { return types.ModuleName }

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() string { return types.RouterKey }

func (am AppModule) NewHandler() sdk.Handler {
    return NewHandler(am.keeper)
}

func (am AppModule) QuerierRoute() string { return types.QuerierRoute }

func (am AppModule) NewQuerierHandler() sdk.Querier {
    return keeper.NewQuerier(am.keeper)
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
    types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
}

func (am AppModule) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    types.RegisterInterfaces(registry)
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
    var genesisState types.GenesisState
    cdc.MustUnmarshalJSON(data, &genesisState)
    return am.keeper.InitGenesis(ctx, genesisState)
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
    gs := am.keeper.ExportGenesis(ctx)
    return cdc.MustMarshalJSON(gs)
}

func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {}

func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
    return nil
}
