package model

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/Linrena/go_gen/internal/tpl"

	"github.com/Linrena/go_gen/internal/util/consts"

	"github.com/Linrena/go_gen/internal/util"
	"github.com/fatih/color"

	"github.com/iancoleman/strcase"
	"github.com/xwb1989/sqlparser"
)

const (
	fieldTag      = `gorm:"column:%s" json:"%s"`
	fieldPrefix   = "    "
	commentPrefix = " // "
)

type Generator struct {
	ddl         *sqlparser.DDL
	output      string
	output2Cmd  bool
	initialisms bool
	override    bool
}

func NewGenerator(ddl *sqlparser.DDL, output string, initialisms bool, override bool) *Generator {
	return &Generator{
		ddl:         ddl,
		output:      output,
		output2Cmd:  output == consts.CmdKey,
		initialisms: initialisms,
		override:    override,
	}
}

func (g *Generator) Generate() error {
	var moduleName string
	var err error
	if !g.output2Cmd { // output to file
		if err = g.checkAndCreateDir(); err != nil {
			return err
		}

		if util.IsFileExists(g.getFilePath()) && !g.override {
			// file exists and override is false
			color.Yellow("%s already exists, skip", g.getFileName())
			return nil
		}
	}

	moduleName, err = util.GetModuleNameByPath(g.output)
	if err != nil {
		return err
	}

	fields := g.getTableFields()

	tpl_ := template.Must(template.New(g.getFileName()).Parse(tpl.ModelTpl))
	buf := bytes.NewBuffer(nil)

	err = tpl_.Execute(buf, tpl.ModelTplData{
		ModuleName:     moduleName,
		ModelName:      g.getModelName(),
		ModelSnakeName: g.ddl.NewName.Name.String(),
		Fields:         fields,
	})
	if err != nil {
		return err
	}
	content := buf.Bytes()

	if !g.output2Cmd {
		if err = util.WriteFile(g.getFilePath(), content, true); err != nil {
			return err
		}

		if err = util.GoImportFile(g.getFilePath()); err != nil {
			return err
		}
	} else {
		log.Println(g.getModelName() + "\n" + string(content))
	}

	return nil
}

func (g *Generator) checkAndCreateDir() error {
	if !util.IsDirExists(g.output) {
		// create dir
		return util.CreateDir(g.output)
	}
	return nil
}

func (g *Generator) getFilePath() string {
	return strings.Join([]string{g.output, g.getFileName()}, "/")
}

func (g *Generator) getFileName() string {
	return g.ddl.NewName.Name.String() + "_model.go"
}

func (g *Generator) getModelName() string {
	return strcase.ToCamel(g.ddl.NewName.Name.String())
}

func (g *Generator) getTableFields() string {
	fieldList := make([]string, 0)
	for _, col := range g.ddl.TableSpec.Columns {
		fieldName := strcase.ToCamel(col.Name.String())
		if g.initialisms {
			fieldName = util.InitialismsWithCamel(fieldName)
		}

		typeName := consts.TypeFromMysqlToGo[col.Type.Type]

		tag := "`" + fmt.Sprintf(fieldTag, col.Name.String(), col.Name.String()) + "`"

		comment := ""
		if col.Type.Comment != nil {
			comment = commentPrefix + string(col.Type.Comment.Val)
		}

		field := fieldPrefix + fieldName + " " + typeName + " " + tag + comment
		fieldList = append(fieldList, field)
	}
	return strings.Join(fieldList, "\n")
}
