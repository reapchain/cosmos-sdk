package server

// DONTCOVER

import (
	"fmt"
	"io/ioutil"
	"os"

	tmjson "github.com/reapchain/reapchain-core/libs/json"
	tmproto "github.com/reapchain/reapchain-core/proto/podc/types"
	tmtypes "github.com/reapchain/reapchain-core/types"
	"github.com/spf13/cobra"

	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/reapchain/cosmos-sdk/server/types"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sm "github.com/reapchain/reapchain-core/state"
)

const (
	FlagHeight           = "height"
	FlagForZeroHeight    = "for-zero-height"
	FlagJailAllowedAddrs = "jail-allowed-addrs"
)

// ExportCmd dumps app state to JSON.
func ExportCmd(appExporter types.AppExporter, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export state to JSON",
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			homeDir, _ := cmd.Flags().GetString(flags.FlagHome)
			config.SetRoot(homeDir)

			if _, err := os.Stat(config.GenesisFile()); os.IsNotExist(err) {
				return err
			}

			db, err := openDB(config.RootDir)
			if err != nil {
				return err
			}

			if appExporter == nil {
				if _, err := fmt.Fprintln(os.Stderr, "WARNING: App exporter not defined. Returning genesis file."); err != nil {
					return err
				}

				genesis, err := ioutil.ReadFile(config.GenesisFile())
				if err != nil {
					return err
				}

				fmt.Println(string(genesis))
				return nil
			}

			traceWriterFile, _ := cmd.Flags().GetString(flagTraceStore)
			traceWriter, err := openTraceWriter(traceWriterFile)
			if err != nil {
				return err
			}

			height, _ := cmd.Flags().GetInt64(FlagHeight)
			forZeroHeight, _ := cmd.Flags().GetBool(FlagForZeroHeight)
			jailAllowedAddrs, _ := cmd.Flags().GetStringSlice(FlagJailAllowedAddrs)

			exported, err := appExporter(serverCtx.Logger, db, traceWriter, height, forZeroHeight, jailAllowedAddrs, serverCtx.Viper)
			if err != nil {
				return fmt.Errorf("error exporting state: %v", err)
			}

			doc, err := tmtypes.GenesisDocFromFile(serverCtx.Config.GenesisFile())
			if err != nil {
				return err
			}

			exportedState, err := sm.ExportState(serverCtx.Config, exported.Height)
			if err != nil {
				return err
			}

			doc.StandingMembers = make([]tmtypes.GenesisMember, len(exportedState.StandingMemberSet.StandingMembers))

			for i, standingMember := range exportedState.StandingMemberSet.StandingMembers {
					if standingMember != nil {
							doc.StandingMembers[i] = tmtypes.GenesisMember{
									Address: standingMember.Address,
									PubKey: standingMember.PubKey,
									Power: standingMember.VotingPower,
									Name: "",
							}
					}
			}
			
			doc.SteeringMemberCandidates = make([]tmtypes.GenesisMember, len(exportedState.SteeringMemberCandidateSet.SteeringMemberCandidates))
			for i, steeringMemberCandidate := range exportedState.SteeringMemberCandidateSet.SteeringMemberCandidates {
					if steeringMemberCandidate != nil {
							doc.SteeringMemberCandidates[i] = tmtypes.GenesisMember{
									Address: steeringMemberCandidate.Address,
									PubKey: steeringMemberCandidate.PubKey,
									Power: steeringMemberCandidate.VotingPower,
									Name: "",
							}
					}
			}
			
			doc.ConsensusRound = tmtypes.ConsensusRound {
				ConsensusStartBlockHeight: exportedState.ConsensusRound.ConsensusStartBlockHeight,
				Period: exportedState.ConsensusRound.Period,
				QrnPeriod: exportedState.ConsensusRound.QrnPeriod,
				VrfPeriod: exportedState.ConsensusRound.VrfPeriod,
				ValidatorPeriod: exportedState.ConsensusRound.ValidatorPeriod,
			}
			
			
			doc.Qrns = make([]tmtypes.Qrn, len(exportedState.QrnSet.Qrns))
			for i, qrn := range exportedState.QrnSet.Qrns {
				doc.Qrns[i] = *qrn.Copy()
			}

			doc.NextQrns = make([]tmtypes.Qrn, len(exportedState.NextQrnSet.Qrns))
			for i, nextQrn := range exportedState.NextQrnSet.Qrns {
					doc.NextQrns[i] = *nextQrn.Copy()
			}
			
			doc.Vrfs = make([]tmtypes.Vrf, len(exportedState.VrfSet.Vrfs))
			for i, vrf := range exportedState.VrfSet.Vrfs {
					doc.Vrfs[i] = *vrf.Copy()
			}

			doc.NextVrfs = make([]tmtypes.Vrf, len(exportedState.NextVrfSet.Vrfs))
			for i, nextVrf := range exportedState.NextVrfSet.Vrfs {
					doc.NextVrfs[i] = *nextVrf.Copy()
			}

			doc.AppState = exported.AppState
			doc.Validators = exported.Validators
			doc.InitialHeight = exported.Height
			doc.ConsensusParams = &tmproto.ConsensusParams{
				Block: tmproto.BlockParams{
					MaxBytes:   exported.ConsensusParams.Block.MaxBytes,
					MaxGas:     exported.ConsensusParams.Block.MaxGas,
					TimeIotaMs: doc.ConsensusParams.Block.TimeIotaMs,
				},
				Evidence: tmproto.EvidenceParams{
					MaxAgeNumBlocks: exported.ConsensusParams.Evidence.MaxAgeNumBlocks,
					MaxAgeDuration:  exported.ConsensusParams.Evidence.MaxAgeDuration,
					MaxBytes:        exported.ConsensusParams.Evidence.MaxBytes,
				},
				Validator: tmproto.ValidatorParams{
					PubKeyTypes: exported.ConsensusParams.Validator.PubKeyTypes,
				},
			}

			

			// NOTE: Tendermint uses a custom JSON decoder for GenesisDoc
			// (except for stuff inside AppState). Inside AppState, we're free
			// to encode as protobuf or amino.
			encoded, err := tmjson.Marshal(doc)
			if err != nil {
				return err
			}

			cmd.Println(string(sdk.MustSortJSON(encoded)))
			return nil
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().Int64(FlagHeight, -1, "Export state from a particular height (-1 means latest height)")
	cmd.Flags().Bool(FlagForZeroHeight, false, "Export state to start at height zero (perform preproccessing)")
	cmd.Flags().StringSlice(FlagJailAllowedAddrs, []string{}, "Comma-separated list of operator addresses of jailed validators to unjail")

	return cmd
}
