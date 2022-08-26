package tpl

var ModelTpl = `package {{.ModuleName}}

func ({{.ModelName}}) TableName() string {
	return "{{.ModelSnakeName}}"
}

type {{.ModelName}} struct { // nolint:maligned
{{.Fields}}
}

func New{{.ModelName}}() *{{.ModelName}} {
	return &{{.ModelName}}{}
}
`

type ModelTplData struct {
	ModuleName     string
	ModelName      string
	ModelSnakeName string
	Fields         string
}
