package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"lotterychain/x/lottery/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ParticipantList: []types.Participant{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ParticipantCount: 2,
				Lottery: &types.Lottery{
					TxCounter: 47,
					TxDataAll: "14",
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated participant",
			genState: &types.GenesisState{
				ParticipantList: []types.Participant{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid participant count",
			genState: &types.GenesisState{
				ParticipantList: []types.Participant{
					{
						Id: 1,
					},
				},
				ParticipantCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
