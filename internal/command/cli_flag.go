package command

import "github.com/urfave/cli/v2"

const (
	ddlFlagName         = "ddlPath"
	outputFlagName      = "output"
	initialismsFlagName = "initialisms"
	overrideFlagName    = "override"
	tablesFlagName      = "tables"
)

var (
	ddlFlag = &cli.StringFlag{
		Name:     ddlFlagName,
		Usage:    "ddl sql file path(absolute or relative), end with .sql or dns address(connect to rds to generate)",
		Required: true,
	}
	outputFlag = &cli.StringFlag{
		Name:        outputFlagName,
		Usage:       "output dir(absolute or relative or output to cmd if input CMD)",
		Required:    false,
		DefaultText: "default: .",
	}
	initialismsFlag = &cli.BoolFlag{
		Name:        initialismsFlagName,
		Usage:       "use common initialisms like Id -> ID(referer: https://github.com/golang/lint/blob/master/lint.go#L770)",
		Required:    false,
		DefaultText: "default: false",
	}
	overrideFlag = &cli.BoolFlag{
		Name:        overrideFlagName,
		Usage:       "override when output file exists",
		Required:    false,
		DefaultText: "default: false",
	}
	tablesFlag = &cli.StringSliceFlag{
		Name:        tablesFlagName,
		Usage:       "(ddlPath -> dns addr)specific tables(split by comma) to generate, if not set, all tables in db will be generated",
		Required:    false,
		DefaultText: "default: all tables",
	}
)
