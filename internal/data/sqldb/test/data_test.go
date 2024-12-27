package test

import (
	"appix/internal/data/sqldb"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var logger = log.With(log.NewStdLogger(os.Stdout))

var dataMem = DataMem()

func DataMem() *sqldb.DataGorm {
	_db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	return &sqldb.DataGorm{DB: _db}
}
