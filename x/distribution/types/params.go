package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/reapchain/cosmos-sdk/types"
	paramtypes "github.com/reapchain/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	ParamStoreKeyCommunityTax             = []byte("communitytax")
	ParamStoreKeyBaseProposerReward       = []byte("baseproposerreward")
	ParamStoreKeyBonusProposerReward      = []byte("bonusproposerreward")
	ParamStoreKeyWithdrawAddrEnabled      = []byte("withdrawaddrenabled")
	ParamStoreKeyStandingMemberRewardRate = []byte("standingmemberrewardrate")
	ParamStoreKeySteeringMemberRewardRate = []byte("steeringmemberrewardrate")
	ParamStoreKeyAllMemberRewardRate      = []byte("allmemberrewardrate")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return Params{
		CommunityTax:        sdk.NewDecWithPrec(0, 2), // 0%
		BaseProposerReward:  sdk.NewDecWithPrec(1, 2), // 1%
		BonusProposerReward: sdk.NewDecWithPrec(4, 2), // 4%
		WithdrawAddrEnabled: true,
	}
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyCommunityTax, &p.CommunityTax, validateCommunityTax),
		paramtypes.NewParamSetPair(ParamStoreKeyBaseProposerReward, &p.BaseProposerReward, validateBaseProposerReward),
		paramtypes.NewParamSetPair(ParamStoreKeyBonusProposerReward, &p.BonusProposerReward, validateBonusProposerReward),
		paramtypes.NewParamSetPair(ParamStoreKeyWithdrawAddrEnabled, &p.WithdrawAddrEnabled, validateWithdrawAddrEnabled),
		paramtypes.NewParamSetPair(ParamStoreKeyStandingMemberRewardRate, &p.StandingMemberRewardRate, validateStandingMemberRewardRate),
		paramtypes.NewParamSetPair(ParamStoreKeySteeringMemberRewardRate, &p.SteeringMemberRewardRate, validateSteeringMemberRewardRate),
		paramtypes.NewParamSetPair(ParamStoreKeyAllMemberRewardRate, &p.AllMemberRewardRate, validateAllMemberRewardRate),
	}
}

// ValidateBasic performs basic validation on distribution parameters.
func (p Params) ValidateBasic() error {
	if p.CommunityTax.IsNegative() || p.CommunityTax.GT(sdk.OneDec()) {
		return fmt.Errorf(
			"community tax should be non-negative and less than one: %s", p.CommunityTax,
		)
	}
	if p.BaseProposerReward.IsNegative() {
		return fmt.Errorf(
			"base proposer reward should be positive: %s", p.BaseProposerReward,
		)
	}
	if p.BonusProposerReward.IsNegative() {
		return fmt.Errorf(
			"bonus proposer reward should be positive: %s", p.BonusProposerReward,
		)
	}
	if v := p.BaseProposerReward.Add(p.BonusProposerReward).Add(p.CommunityTax); v.GT(sdk.OneDec()) {
		return fmt.Errorf(
			"sum of base, bonus proposer rewards, and community tax cannot be greater than one: %s", v,
		)
	}

	if v := p.StandingMemberRewardRate.Add(p.SteeringMemberRewardRate).Add(p.AllMemberRewardRate); v.GT(sdk.OneDec()) {
		return fmt.Errorf(
			"sum of standing member, steering member and All member reward cannot be greater than one: %s", v,
		)
	}

	return nil
}

func validateCommunityTax(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("community tax must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("community tax must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("community tax too large: %s", v)
	}

	return nil
}

func validateBaseProposerReward(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("base proposer reward must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("base proposer reward must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("base proposer reward too large: %s", v)
	}

	return nil
}

func validateBonusProposerReward(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("bonus proposer reward must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("bonus proposer reward must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("bonus proposer reward too large: %s", v)
	}

	return nil
}

func validateWithdrawAddrEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateStandingMemberRewardRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("StandingMemberRewardRate must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("StandingMemberRewardRate must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("StandingMemberRewardRate too large: %s", v)
	}

	return nil
}

func validateSteeringMemberRewardRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("SteeringMemberRewardRate must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("SteeringMemberRewardRate must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("SteeringMemberRewardRate too large: %s", v)
	}

	return nil
}

func validateAllMemberRewardRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("AllMemberRewardRate must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("AllMemberRewardRate must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("AllMemberRewardRate too large: %s", v)
	}

	return nil
}
