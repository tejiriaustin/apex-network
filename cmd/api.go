package cmd

import (
	"context"
	"github.com/tejiriaustin/apex-network/server"

	"github.com/spf13/cobra"

	"github.com/tejiriaustin/apex-network/database"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/repository"
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

	rc := repository.NewRepo(config, dbConn)

	sc := service.NewService(config)

	server.Start(ctx, sc, rc)
}

func SetEnvironmentConfigs() env.Env {
	config := env.NewEnv()

	config.SetEnv("servie_name", "apex-network-api").
		SetEnv("PORT", "8080").
		SetEnv(env.DbUsername, env.MustGetEnv(env.DbUsername)).
		SetEnv(env.DbHost, env.MustGetEnv(env.DbHost)).
		SetEnv(env.DbPassword, env.MustGetEnv(env.DbPassword)).
		SetEnv(env.DbTimeZone, env.MustGetEnv(env.DbTimeZone)).
		SetEnv(env.DbDatabase, env.MustGetEnv(env.DbDatabase))

	return config
}
