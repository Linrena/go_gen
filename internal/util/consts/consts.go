package consts

var (
	// CommonInitialisms refer to referer: https://github.com/golang/lint/blob/master/lint.go#L770
	CommonInitialisms = map[string]bool{
		"ACL":   true,
		"API":   true,
		"ASCII": true,
		"CPU":   true,
		"CSS":   true,
		"DNS":   true,
		"EOF":   true,
		"GUID":  true,
		"HTML":  true,
		"HTTP":  true,
		"HTTPS": true,
		"ID":    true,
		"IP":    true,
		"JSON":  true,
		"LHS":   true,
		"QPS":   true,
		"RAM":   true,
		"RHS":   true,
		"RPC":   true,
		"SLA":   true,
		"SMTP":  true,
		"SQL":   true,
		"SSH":   true,
		"TCP":   true,
		"TLS":   true,
		"TTL":   true,
		"UDP":   true,
		"UI":    true,
		"UID":   true,
		"UUID":  true,
		"URI":   true,
		"URL":   true,
		"UTF8":  true,
		"VM":    true,
		"XML":   true,
		"XMPP":  true,
		"XSRF":  true,
		"XSS":   true,
	}

	// TypeFromMysqlToGo refer to https://github.com/gohouse/converter/blob/master/table2struct.go
	TypeFromMysqlToGo = map[string]string{
		"int":                "int32",
		"integer":            "int32",
		"tinyint":            "int8",
		"smallint":           "int16",
		"mediumint":          "int32",
		"bigint":             "int64",
		"int unsigned":       "int64",
		"integer unsigned":   "int64",
		"tinyint unsigned":   "int64",
		"smallint unsigned":  "int64",
		"mediumint unsigned": "int64",
		"bigint unsigned":    "int64",
		"bit":                "int64",
		"bool":               "bool",
		"enum":               "string",
		"set":                "string",
		"varchar":            "string",
		"char":               "string",
		"tinytext":           "string",
		"mediumtext":         "string",
		"text":               "string",
		"longtext":           "string",
		"blob":               "string",
		"tinyblob":           "string",
		"mediumblob":         "string",
		"longblob":           "string",
		"date":               "time.Time", // time.Time or string
		"datetime":           "time.Time", // time.Time or string
		"timestamp":          "time.Time", // time.Time or string
		"time":               "time.Time", // time.Time or string
		"float":              "float32",
		"double":             "float64",
		"decimal":            "float64",
		"binary":             "string",
		"varbinary":          "string",
		"json":               "string",
	}
)

const (
	DDLQLSeparator          = ";\n"
	CreateTablePrefixRegExp = `(?i)^(\s)*(create)(\s)+table.*`
	CmdKey                  = "CMD"
)
