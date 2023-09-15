package cli

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	tmos "github.com/reapchain/reapchain-core/libs/os"
	tmtypes "github.com/reapchain/reapchain-core/types"
	"github.com/spf13/cobra"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/reapchain/cosmos-sdk/client/tx"
	"github.com/reapchain/cosmos-sdk/crypto/keyring"
	"github.com/reapchain/cosmos-sdk/server"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/module"
	"github.com/reapchain/cosmos-sdk/version"
	authclient "github.com/reapchain/cosmos-sdk/x/auth/client"
	"github.com/reapchain/cosmos-sdk/x/genutil"
	"github.com/reapchain/cosmos-sdk/x/genutil/types"
	"github.com/reapchain/cosmos-sdk/x/staking/client/cli"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
)

// GenTxCmd builds the application's gentx command.
func GenTxCmd(mbm module.BasicManager, txEncCfg client.TxEncodingConfig, genBalIterator types.GenesisBalancesIterator, defaultNodeHome string) *cobra.Command {
	ipDefault, _ := server.ExternalIP()
	fsCreateValidator, defaultsDesc := cli.CreateValidatorMsgFlagSet(ipDefault)

	cmd := &cobra.Command{
		Use:   "gentx [key_name] [amount] [validator type - standing | steering]",
		Short: "Generate a genesis tx carrying a self delegation",
		Args:  cobra.ExactArgs(3),
		Long: fmt.Sprintf(`Generate a genesis transaction that creates a validator with a self-delegation,
that is signed by the key in the Keyring referenced by a given name. A node ID and Bech32 consensus
pubkey may optionally be provided. If they are omitted, they will be retrieved from the priv_validator.json
file. The following default parameters are included:
    %s

Example:
$ %s gentx my-key-name 100000stake steering --home=/path/to/home/dir --keyring-backend=os --chain-id=test-chain-1 \
    --moniker="myvalidator" \
    --commission-max-change-rate=0.01 \
    --commission-max-rate=1.0 \
    --commission-rate=0.07 \
    --details="..." \
    --security-contact="..." \
    --website="..."
`, defaultsDesc, version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			cdc := clientCtx.Codec

			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(serverCtx.Config)
			if err != nil {
				return errors.Wrap(err, "failed to initialize node validator files")
			}

			// read --nodeID, if empty take it from priv_validator.json
			if nodeIDString, _ := cmd.Flags().GetString(cli.FlagNodeID); nodeIDString != "" {
				nodeID = nodeIDString
			}

			// read --pubkey, if empty take it from priv_validator.json
			if pkStr, _ := cmd.Flags().GetString(cli.FlagPubKey); pkStr != "" {
				if err := clientCtx.Codec.UnmarshalInterfaceJSON([]byte(pkStr), &valPubKey); err != nil {
					return errors.Wrap(err, "failed to unmarshal validator public key")
				}
			}

			genDoc, err := tmtypes.GenesisDocFromFile(config.GenesisFile())
			if err != nil {
				return errors.Wrapf(err, "failed to read genesis doc file %s", config.GenesisFile())
			}

			var genesisState map[string]json.RawMessage
			if err = json.Unmarshal(genDoc.AppState, &genesisState); err != nil {
				return errors.Wrap(err, "failed to unmarshal genesis state")
			}

			if err = mbm.ValidateGenesis(cdc, txEncCfg, genesisState); err != nil {
				return errors.Wrap(err, "failed to validate genesis state")
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())

			name := args[0]
			key, err := clientCtx.Keyring.Key(name)
			if err != nil {
				return errors.Wrapf(err, "failed to fetch '%s' from the keyring", name)
			}

			moniker := config.Moniker
			if m, _ := cmd.Flags().GetString(cli.FlagMoniker); m != "" {
				moniker = m
			}

			// validator type [standing | steering]
			valType := args[2]

			// check validator type
			if valType != stakingtypes.ValidatorTypeStanding && valType != stakingtypes.ValidatorTypeSteering {
				return errors.New("The validator type must be one of standing or steering.")
			}

			// set flags for creating a gentx
			createValCfg, err := cli.PrepareConfigForTxCreateValidator(cmd.Flags(), moniker, nodeID, genDoc.ChainID, valPubKey, valType)
			if err != nil {
				return errors.Wrap(err, "error creating configuration to create validator msg")
			}

			amount := args[1]
			coins, err := sdk.ParseCoinsNormalized(amount)
			if err != nil {
				return errors.Wrap(err, "failed to parse coins")
			}

			// validate validator's conditions(Min Staking)
			if err := validateValidatorConditions(genesisState, valType, coins); err != nil {
				return errors.Wrap(err, "failed to validate validator's condition")
			}

			err = genutil.ValidateAccountInGenesis(genesisState, genBalIterator, key.GetAddress(), coins, cdc)
			if err != nil {
				return errors.Wrap(err, "failed to validate account in genesis")
			}

			txFactory := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			if err != nil {
				return errors.Wrap(err, "error creating tx builder")
			}

			clientCtx = clientCtx.WithInput(inBuf).WithFromAddress(key.GetAddress())

			// The following line comes from a discrepancy between the `gentx`
			// and `create-validator` commands:
			// - `gentx` expects amount as an arg,
			// - `create-validator` expects amount as a required flag.
			// ref: https://github.com/reapchain/cosmos-sdk/issues/8251
			// Since gentx doesn't set the amount flag (which `create-validator`
			// reads from), we copy the amount arg into the valCfg directly.
			//
			// Ideally, the `create-validator` command should take a validator
			// config file instead of so many flags.
			// ref: https://github.com/reapchain/cosmos-sdk/issues/8177
			createValCfg.Amount = amount

			// create a 'create-validator' message
			txBldr, msg, err := cli.BuildCreateValidatorMsg(clientCtx, createValCfg, txFactory, true)
			if err != nil {
				return errors.Wrap(err, "failed to build create-validator message")
			}

			if key.GetType() == keyring.TypeOffline || key.GetType() == keyring.TypeMulti {
				cmd.PrintErrln("Offline key passed in. Use `tx sign` command to sign.")
				return authclient.PrintUnsignedStdTx(txBldr, clientCtx, []sdk.Msg{msg})
			}

			// write the unsigned transaction to the buffer
			w := bytes.NewBuffer([]byte{})
			clientCtx = clientCtx.WithOutput(w)

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			if err = authclient.PrintUnsignedStdTx(txBldr, clientCtx, []sdk.Msg{msg}); err != nil {
				return errors.Wrap(err, "failed to print unsigned std tx")
			}

			// read the transaction
			stdTx, err := readUnsignedGenTxFile(clientCtx, w)
			if err != nil {
				return errors.Wrap(err, "failed to read unsigned gen tx file")
			}

			// sign the transaction and write it to the output file
			txBuilder, err := clientCtx.TxConfig.WrapTxBuilder(stdTx)
			if err != nil {
				return fmt.Errorf("error creating tx builder: %w", err)
			}

			err = authclient.SignTx(txFactory, clientCtx, name, txBuilder, true, true)
			if err != nil {
				return errors.Wrap(err, "failed to sign std tx")
			}

			outputDocument, _ := cmd.Flags().GetString(flags.FlagOutputDocument)
			if outputDocument == "" {
				outputDocument, err = makeOutputFilepath(config.RootDir, nodeID)
				if err != nil {
					return errors.Wrap(err, "failed to create output file path")
				}
			}

			if err := writeSignedGenTx(clientCtx, outputDocument, stdTx); err != nil {
				return errors.Wrap(err, "failed to write signed gen tx")
			}

			cmd.PrintErrf("Genesis transaction written to %q\n", outputDocument)
			return nil
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().String(flags.FlagOutputDocument, "", "Write the genesis transaction JSON document to the given file instead of the default location")
	cmd.Flags().String(flags.FlagChainID, "", "The network chain ID")
	cmd.Flags().AddFlagSet(fsCreateValidator)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// validate validator's init conditions
//	amount: standing member >= 2%(MUST)
//	        steering member >= 100,000(MUST)
func validateValidatorConditions(appGenesisState map[string]json.RawMessage, valType string, amount sdk.Coins) error {

	var genesisStaking map[string]json.RawMessage
	if err := json.Unmarshal(appGenesisState["staking"], &genesisStaking); err != nil {
		return errors.Wrap(err, "failed to unmarshal genesis staking state")
	}

	var genesisStakingParams map[string]json.RawMessage
	if err := json.Unmarshal(genesisStaking["params"], &genesisStakingParams); err != nil {
		return errors.Wrap(err, "failed to unmarshal genesis staking params state")
	}

	var minStandingMemberStaking string
	if err := json.Unmarshal(genesisStakingParams["min_standing_member_staking_quantity"], &minStandingMemberStaking); err != nil {
		return errors.Wrap(err, "failed to unmarshal genesis staking params [minStandingMemberStakingCoin] state")
	}

	var minSteeringMemberStakingCoin string
	if err := json.Unmarshal(genesisStakingParams["min_steering_member_staking_quantity"], &minSteeringMemberStakingCoin); err != nil {
		return errors.Wrap(err, "failed to unmarshal genesis staking params [minSteeringMemberStakingCoin] state")
	}

	minStandingCoin, _ := sdk.ParseCoinsNormalized(minStandingMemberStaking + sdk.DefaultBondDenom)
	minSteeringCoin, _ := sdk.ParseCoinsNormalized(minSteeringMemberStakingCoin + sdk.DefaultBondDenom)

	fmt.Printf("Genesis info - Min Standing Member Staking: %s, Min Steering Member Staking: %s\n", minStandingCoin, minSteeringCoin)

	if valType == "standing" && amount.IsAllLT(minStandingCoin) {
		return errors.New(fmt.Sprintf("Standing members must have more than %s. Got: %s", minStandingCoin, amount))
	}

	if valType == "steering" && amount.IsAllLT(minSteeringCoin) {
		return errors.New(fmt.Sprintf("Steering members must have more than %s. Got: %s", minSteeringCoin, amount))
	}

	return nil
}

func makeOutputFilepath(rootDir, nodeID string) (string, error) {
	writePath := filepath.Join(rootDir, "config", "gentx")
	if err := tmos.EnsureDir(writePath, 0o700); err != nil {
		return "", err
	}

	return filepath.Join(writePath, fmt.Sprintf("gentx-%v.json", nodeID)), nil
}

func readUnsignedGenTxFile(clientCtx client.Context, r io.Reader) (sdk.Tx, error) {
	bz, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	aTx, err := clientCtx.TxConfig.TxJSONDecoder()(bz)
	if err != nil {
		return nil, err
	}

	return aTx, err
}

func writeSignedGenTx(clientCtx client.Context, outputDocument string, tx sdk.Tx) error {
	outputFile, err := os.OpenFile(outputDocument, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	json, err := clientCtx.TxConfig.TxJSONEncoder()(tx)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(outputFile, "%s\n", json)

	return err
}
