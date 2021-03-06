package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mycodeku/transtionhelper/codec"
	codectypes "github.com/mycodeku/transtionhelper/codec/types"
	sdk "github.com/mycodeku/transtionhelper/types"
	"github.com/mycodeku/transtionhelper/types/module"
	simtypes "github.com/mycodeku/transtionhelper/types/simulation"
	"github.com/mycodeku/transtionhelper/x/gov/simulation"
	"github.com/mycodeku/transtionhelper/x/gov/types"
	v1 "github.com/mycodeku/transtionhelper/x/gov/types/v1"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abnormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: 1000,
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var govGenesis v1.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &govGenesis)

	dec1, _ := sdk.NewDecFromStr("0.361000000000000000")
	dec2, _ := sdk.NewDecFromStr("0.512000000000000000")
	dec3, _ := sdk.NewDecFromStr("0.267000000000000000")

	require.Equal(t, "905stake", govGenesis.DepositParams.MinDeposit[0].String())
	require.Equal(t, "77h26m10s", govGenesis.DepositParams.MaxDepositPeriod.String())
	require.Equal(t, float64(148296), govGenesis.VotingParams.VotingPeriod.Seconds())
	require.Equal(t, dec1.String(), govGenesis.TallyParams.Quorum)
	require.Equal(t, dec2.String(), govGenesis.TallyParams.Threshold)
	require.Equal(t, dec3.String(), govGenesis.TallyParams.VetoThreshold)
	require.Equal(t, uint64(0x28), govGenesis.StartingProposalId)
	require.Equal(t, []*v1.Deposit{}, govGenesis.Deposits)
	require.Equal(t, []*v1.Vote{}, govGenesis.Votes)
	require.Equal(t, []*v1.Proposal{}, govGenesis.Proposals)
}

// TestRandomizedGenState tests abnormal scenarios of applying RandomizedGenState.
func TestRandomizedGenState1(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)
	// all these tests will panic
	tests := []struct {
		simState module.SimulationState
		panicMsg string
	}{
		{ // panic => reason: incomplete initialization of the simState
			module.SimulationState{}, "invalid memory address or nil pointer dereference"},
		{ // panic => reason: incomplete initialization of the simState
			module.SimulationState{
				AppParams: make(simtypes.AppParams),
				Cdc:       cdc,
				Rand:      r,
			}, "assignment to entry in nil map"},
	}

	for _, tt := range tests {
		require.Panicsf(t, func() { simulation.RandomizedGenState(&tt.simState) }, tt.panicMsg)
	}
}
