package testutil

import (
	"fmt"
	"strings"

	tmcli "github.com/reapchain/reapchain-core/libs/cli"
	"github.com/stretchr/testify/suite"

	clitestutil "github.com/reapchain/cosmos-sdk/testutil/cli"
	"github.com/reapchain/cosmos-sdk/testutil/network"
	"github.com/reapchain/cosmos-sdk/x/params/client/cli"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	s.network = network.New(s.T(), s.cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestNewQuerySubspaceParamsCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectedOutput string
	}{
		{
			"json output",
			[]string{
				"staking", "MaxValidators",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			`{"subspace":"staking","key":"MaxValidators","value":"100"}`,
		},
		{
			"text output",
			[]string{
				"staking", "MaxValidators",
				fmt.Sprintf("--%s=text", tmcli.OutputFlag),
			},
			`key: MaxValidators
subspace: staking
value: "100"`,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewQuerySubspaceParamsCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			s.Require().Equal(tc.expectedOutput, strings.TrimSpace(out.String()))
		})
	}
}
