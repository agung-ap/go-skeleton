package main

import (
	"go-skeleton/cmd"
	"go-skeleton/cmd/app"
	"go-skeleton/config"
	"go-skeleton/pkg/database"
	"go-skeleton/pkg/logger"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	app.Init()
	defer app.ShutDown()

	cliApp := cli.NewApp()
	cliApp.Name = "skeleton: Template for fast bootstrapping"
	cliApp.Version = "1.0.0"

	cliApp.Commands = cli.Commands{
		{
			Name:  "server",
			Usage: "Start server",
			Action: func(c *cli.Context) error {
				logger.Info("Starting server command")
				cmd.StartServer(c.Context)
				return nil
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(_ *cli.Context) error {
				logger.Info("Running database migrations")
				migration, err := database.InitMigration(config.Database)
				if err != nil {
					logger.Error("Failed to initialize migration", zap.Error(err))
					return err
				}

				err = migration.ApplyMigrations()
				if err != nil {
					logger.Error("Failed to apply migrations", zap.Error(err))
				} else {
					logger.Info("Migrations applied successfully")
				}
				return err
			},
			Subcommands: []*cli.Command{
				{
					Name:  "create",
					Usage: "create a new migration file",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "name",
							Aliases:  []string{"n"},
							Usage:    "name of the migration",
							Required: true,
						},
					},
					Action: func(c *cli.Context) error {
						migrationName := c.String("name")
						logger.Info("Creating new migration", zap.String("name", migrationName))

						migration, err := database.InitMigration(config.Database)
						if err != nil {
							logger.Error("Failed to initialize migration", zap.Error(err))
							return err
						}

						err = migration.CreateMigration(migrationName)
						if err != nil {
							logger.Error("Failed to create migration", zap.Error(err))
						} else {
							logger.Info("Migration created successfully", zap.String("name", migrationName))
						}
						return err
					},
				},
				{
					Name:  "up",
					Usage: "apply all migrations",
					Action: func(_ *cli.Context) error {
						logger.Info("Applying migrations up")
						migration, err := database.InitMigration(config.Database)
						if err != nil {
							logger.Error("Failed to initialize migration", zap.Error(err))
							return err
						}

						err = migration.ApplyMigrations()
						if err != nil {
							logger.Error("Failed to apply migrations", zap.Error(err))
						} else {
							logger.Info("Migrations applied successfully")
						}
						return err
					},
				},
				{
					Name:  "down",
					Usage: "apply all migrations",
					Action: func(_ *cli.Context) error {
						logger.Info("Rolling back migrations")
						migration, err := database.InitMigration(config.Database)
						if err != nil {
							logger.Error("Failed to initialize migration", zap.Error(err))
							return err
						}

						err = migration.RollbackMigration()
						if err != nil {
							logger.Error("Failed to rollback migration", zap.Error(err))
						} else {
							logger.Info("Migration rolled back successfully")
						}
						return err
					},
				},
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		logger.Fatal("CLI application error", zap.Error(err))
	}
}
