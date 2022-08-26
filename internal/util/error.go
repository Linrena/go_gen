package util

import "errors"

func InvalidCreateTableSQL(sql string) error {
	return errors.New("invalid create table sql: " + sql)
}
