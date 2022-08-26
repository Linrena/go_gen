package command

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strings"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"github.com/Linrena/go_gen/internal/model"
	"github.com/Linrena/go_gen/internal/util"

	"github.com/Linrena/go_gen/internal/util/consts"
	"github.com/xwb1989/sqlparser"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

// ModelGenerate is a cli.Command for generate model.
func ModelGenerate() *cli.Command {
	return &cli.Command{
		Name:        "model",
		Description: "generate model",
		Flags: []cli.Flag{
			ddlFlag,
			outputFlag,
			tablesFlag,
			initialismsFlag,
			overrideFlag,
		},
		Action: func(c *cli.Context) error {
			g := newGenModelBySQL(c)
			return g.Run()
		},
	}
}

func newGenModelBySQL(c *cli.Context) *genModelBySQL {
	ddl := c.String(ddlFlagName)
	output := c.String(outputFlagName)
	tables := c.StringSlice(tablesFlagName)
	initialisms := c.Bool(initialismsFlagName)
	override := c.Bool(overrideFlagName)
	g := &genModelBySQL{
		ddlPath:     ddl,
		output:      output,
		tables:      tables,
		initialisms: initialisms,
		override:    override,
	}
	return g
}

type genModelBySQL struct {
	ddlPath     string
	output      string
	tables      []string
	initialisms bool
	override    bool

	genBySQL bool
	ddlSQLs  []string
	ddlStmts []*sqlparser.DDL

	db *gorm.DB
}

func (g *genModelBySQL) Run() error {
	if err := g.checkAndInit(); err != nil {
		return err
	}

	if g.genBySQL {
		if err := g.extractDDL(); err != nil {
			return err
		}

		if err := g.parseSQL(); err != nil {
			return err
		}
	} else {
		if err := g.ping(); err != nil {
			return err
		}

		if err := g.processTables(); err != nil {
			return err
		}
	}

	if err := g.genModel(); err != nil {
		return err
	}

	return nil
}

func (g *genModelBySQL) checkAndInit() error {
	if g.ddlPath == "" {
		return errors.New("ddlPath file must be specified")
	}

	if strings.HasSuffix(g.ddlPath, ".sql") {
		g.genBySQL = true
	}

	if g.output == "" {
		g.output = "."
	}

	if g.override {
		color.Red("Now using override mode, please make sure you do not edit target files manually.")
	}

	return nil
}

func (g *genModelBySQL) extractDDL() error {
	content, err := ioutil.ReadFile(g.ddlPath)
	if err != nil {
		return err
	}

	splits := strings.Split(string(content), consts.DDLQLSeparator)
	reg := regexp.MustCompile(consts.CreateTablePrefixRegExp)

	for _, split := range splits {
		if split == "" {
			continue
		}

		split = strings.TrimSpace(split)
		split = strings.Replace(split, "\n", "", -1)

		// judge if it is a `create table sql`
		if reg.Match([]byte(split)) {
			g.ddlSQLs = append(g.ddlSQLs, split)
		}
	}

	return nil
}

func (g *genModelBySQL) parseSQL() error {
	for _, ddl := range g.ddlSQLs {
		// verify sql syntax
		stmt, err := sqlparser.ParseStrictDDL(ddl)
		if err != nil {
			return errors.New("invalid create table ddl sql: " + ddl)
		}

		ddlStmt, ok := stmt.(*sqlparser.DDL)
		if !ok {
			return errors.New("invalid create table ddl sql: " + ddl)
		}

		g.ddlStmts = append(g.ddlStmts, ddlStmt)
	}

	return nil
}

func (g *genModelBySQL) processTables() error {
	if len(g.tables) == 0 { // all tables
		var results [][]uint8
		if err := g.db.Raw("SHOW TABLES").Find(&results).Error; err != nil {
			return err
		}
		for _, res := range results {
			g.tables = append(g.tables, string(res))
		}
	}

	// check table exists and get ddl sql
	for _, table := range g.tables {
		result := struct {
			Table       string `gorm:"column:TABLE"`
			CreateTable string `gorm:"column:Create Table"`
		}{}
		if err := g.db.Raw("SHOW CREATE TABLE " + table).Scan(&result).Error; err != nil { // ignore_security_alert
			return err
		}
		if result.CreateTable != "" {
			g.ddlSQLs = append(g.ddlSQLs, result.CreateTable)
			stmt, err := sqlparser.ParseStrictDDL(result.CreateTable)
			if err != nil {
				return err
			}
			g.ddlStmts = append(g.ddlStmts, stmt.(*sqlparser.DDL)) // show create table must be ddl stmt
		}
	}

	return nil
}

func (g *genModelBySQL) ping() error {
	db, err := gorm.Open(mysql.Open(g.ddlPath))
	if err != nil {
		return err
	}
	g.db = db
	return nil
}

func (g *genModelBySQL) genModel() error {
	for _, stmt := range g.ddlStmts {
		generator := model.NewGenerator(stmt, g.output, g.initialisms, g.override)
		if err := generator.Generate(); err != nil {
			return err
		}
	}

	util.ModelGenFinished()
	return nil
}
