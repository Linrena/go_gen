# go_gen
go_gen is a code generator for golang

## Functions
### generate rds db model
- generate with rds ddl sql
```bash
cd xxx_repo(e.g. cd go_gen)
go_gen model -ddlPath internal/testdata/test.sql -output internal/testdata -override
```

- generate with db dns address (and use initialisms)
```bash
cd xxx_repo(e.g. cd go_gen)
# gen all 
go_gen model -ddlPath "user:password@tcp(127.0.0.1:3306)/test" -output internal/testdata/dbmodels -override -initialisms
# specify some tables
go_gen model -ddlPath "user:password@tcp(127.0.0.1:3306)/test" -output internal/testdata/dbmodels -override -initialisms -tables test1,a
```