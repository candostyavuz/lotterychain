package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ParticipantList: []Participant{},
		Lottery: &Lottery{
			TxCounter:     0,
			TotalFees:     sdk.NewCoin("token", sdk.ZeroInt()),
			TotalBets:     sdk.NewCoin("token", sdk.ZeroInt()),
			CurrentMinBet: sdk.NewCoin("token", sdk.NewInt(9223372036854775807)),
			CurrentMaxBet: sdk.NewCoin("token", sdk.ZeroInt()),
			TxDataAll:     "", // TBD
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in participant
	participantIdMap := make(map[uint64]bool)
	participantCount := gs.GetParticipantCount()
	for _, elem := range gs.ParticipantList {
		if _, ok := participantIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for participant")
		}
		if elem.Id >= participantCount {
			return fmt.Errorf("participant id should be lower or equal than the last id")
		}
		participantIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
