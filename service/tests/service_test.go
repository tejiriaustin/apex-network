package tests

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/tejiriaustin/apex-network/models"
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

func (suite *ServiceTestSuite) TestFundWallet() {
	suite.NotPanics(func() {
		ctx := context.Background()

		type testCase struct {
			name              string
			input             func() service.FundWalletInput
			player            func() models.Player
			walletTransaction func(id uuid.UUID) models.WalletTransaction
			output            func()
		}

		testCases := []testCase{
			{
				input: func() service.FundWalletInput {
					return service.FundWalletInput{PlayerId: "12345"}
				},
				player: func() models.Player {
					return models.Player{
						FirstName: "Tejiri",
						LastName:  "Dev",
					}
				},
				walletTransaction: func(id uuid.UUID) models.WalletTransaction {
					return models.WalletTransaction{
						PlayerId: id,
					}
				},
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {

				player := tc.player()
				walletTransaction := tc.walletTransaction(player.ID)

				playerRepo := new(repomocks.PlayerRepositoryInterface)
				playerRepo.On("GetPlayerbyID", mock.Anything, mock.Anything).Return(&player, nil)

				playerRepo.On("UpdatePlayer", mock.Anything, mock.Anything, mock.Anything).Return(&player, nil)

				walletRepo := new(repomocks.WalletRepositoryInterface)
				walletRepo.On("CreateTransaction", mock.Anything, mock.Anything).Return(&walletTransaction, nil)

				_, err := suite.service.FundWallet(ctx, tc.input(), playerRepo, walletRepo)
				if err != nil {
					suite.NotNil(err)
				}
				suite.Nil(err)
			})
		}
	})
}

func (suite *ServiceTestSuite) TestCreatePlayer() {
	suite.NotPanics(func() {
		ctx := context.Background()

		type testCase struct {
			name   string
			input  func() service.CreatePlayerInput
			output func() models.Player
		}

		testCases := []testCase{
			{
				name: "successfully create account",
				input: func() service.CreatePlayerInput {
					return service.CreatePlayerInput{
						FirstName: "Tejiri",
						LastName:  "Dev",
					}
				},
				output: func() models.Player {
					return models.Player{
						FirstName: "Tejiri",
						LastName:  "Dev",
						FullName:  " Tejiri Dev",
					}
				},
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {

				input := tc.input()
				output := tc.output()

				playerRepo := new(repomocks.PlayerRepositoryInterface)
				playerRepo.On("CreatePlayer", mock.Anything, mock.MatchedBy(func(i models.Player) bool {
					return i.FirstName == input.FirstName && i.LastName == input.LastName
				})).Return(&output, nil)

				_, err := suite.service.CreatePlayer(ctx, input, playerRepo)
				if err != nil {
					suite.NotNil(err)
				}
				suite.Nil(err)
			})
		}

	})
}
