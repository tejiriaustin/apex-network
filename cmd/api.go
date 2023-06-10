package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/tejiriaustin/apex-network/database"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/repository"
	"github.com/tejiriaustin/apex-network/server"
	"github.com/tejiriaustin/apex-network/service"
)

var apiServerCmd = &cobra.Command{
	Use:   "apex_network_api",
	Short: "Starts the apex-network API",
	Long:  ``,
	Run:   startApiServer,
}

func init() {
	rootCmd.AddCommand(apiServerCmd)
}

func startApiServer(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	config := SetEnvironmentConfigs()

	dbConn := database.OpenDatabaseConnection(config)

	rc := repository.NewRepositoryContainer(config, dbConn)

	sc := service.NewService(config)

	server.Start(ctx, sc, rc)
}

func SetEnvironmentConfigs() env.Env {
	config := env.NewEnv()

	config.SetEnv("servie_name", "apex-network-api").
		SetEnv(env.Port, env.MustGetEnv(env.Port)).
		SetEnv(env.DbUrl, env.MustGetEnv(env.DbUrl))

	return config
}
