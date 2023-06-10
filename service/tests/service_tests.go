package tests

import (
	"context"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/service"
	"github.com/tejiriaustin/apex-network/testutils"

	repomocks "github.com/tejiriaustin/apex-network/testutils/mocks/repository"
)

type ServiceTestSuite struct {
	testutils.BaseSuite
	service service.ServiceInterface
}

func TestService(t *testing.T) {
	config := env.NewEnv()

	testService := &ServiceTestSuite{
		service: service.NewService(config),
	}
	suite.Run(t, testService)
}

func (suite *ServiceTestSuite) FundWallet() {
	suite.NotPanics(func() {
		ctx := context.Background()

		type testCase struct {
			name   string
			input  func() service.FundWalletInput
			output func()
		}

		testCases := []testCase{
			{
				input: func() service.FundWalletInput {
					return service.FundWalletInput{PlayerId: "12345"}
				},
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {

				playerRepo := new(repomocks.PlayerRepositoryInterface)
				playerRepo.On("", mock.Anything, mock.Anything)

				walletRepo := new(repomocks.WalletRepositoryInterface)
				walletRepo.On("", mock.Anything, mock.Anything)

				_, err := suite.service.FundWallet(ctx, tc.input(), playerRepo, walletRepo)
				if err != nil {
					suite.NotNil(err)
				}
				suite.Nil(err)
			})
		}
	})
}
