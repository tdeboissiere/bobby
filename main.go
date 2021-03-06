package main

import (
	commands4 "bobby/commands/cluster"
	commands3 "bobby/commands/deploy"
	environments "bobby/commands/environments"
	commands "bobby/commands/init"
	commands2 "bobby/commands/services"
	"github.com/urfave/cli"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Init a new gcp project",
			Action:  commands.Init,
		},
		{
			Name:      "deploy",
			Aliases:   []string{"d"},
			Usage:     "Deploy ",
			ArgsUsage: "[SERVICE_NAME] [DOCKER_IMAGE]",
			Action:    commands3.Deploy,
		}, {
			Name:    "cluster",
			Aliases: []string{"s"},
			Usage:   "Manage cluster",
			Subcommands: []cli.Command{
				{
					Name:    "info",
					Aliases: []string{"l"},
					Usage:   "Info about the cluster",
					Action:  commands4.ClusterInfo,
				},
			},
		},
		{
			Name:    "services",
			Aliases: []string{"s"},
			Usage:   "Manage services",
			Subcommands: []cli.Command{
				{
					Name:    "create",
					Aliases: []string{"c"},

					Usage:     "Create a new services",
					ArgsUsage: "[NAME] [URL]",
					Action:    commands2.ServicesCreate,
				},
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "List all services",
					Action:  commands2.ServicesList,
				},
				{
					Name:      "delete",
					Aliases:   []string{"d"},
					ArgsUsage: "[NAME]",
					Usage:     "Delete service",
					Action:    commands2.ServicesDelete,
				},
				{
					Name:      "scale",
					Aliases:   []string{"s"},
					ArgsUsage: "[NAME] [DYNO_COUNT]",
					Usage:     "Scale service",
					Action:    commands2.ServicesScale,
				},
				{
					Name:      "publish",
					Aliases:   []string{"p"},
					ArgsUsage: "[NAME] [URL]",
					Usage:     "Publish a service",
					Action:    commands2.ServicesPublish,
				},
			},
		},
		{
			Name:    "env",
			Aliases: []string{"s"},
			Usage:   "Manage environments",
			Subcommands: []cli.Command{
				{
					Name:    "create",
					Aliases: []string{"c"},

					Usage:     "Create a new environments",
					ArgsUsage: "[NAME]",
					Action:    environments.EnvironmentCreate,
				},
				{
					Name:    "list",
					Aliases: []string{"l"},

					Usage:  "List environments",
					Action: environments.EnvironmentList,
				},
				{
					Name:      "add",
					Aliases:   []string{"a"},
					ArgsUsage: "[ENVIRONMENT_NAME CONFIG_NAME CONFIG_VALUE]",
					Usage:     "Add variable in  environment",
					Action:    environments.EnvironmentAddConfig,
				},
				{
					Name:      "detail",
					Aliases:   []string{"a"},
					ArgsUsage: "[ENVIRONMENT_NAME]",
					Usage:     "Get detail of an environment",
					Action:    environments.EnvironmentDetail,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
