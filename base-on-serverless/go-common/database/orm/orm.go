/*
   @Time : 2020/4/22 3:17 下午
   @Author : Rebeta
   @Email : master@rebeta.cn
   @File : orm
   @Software: GoLand
*/

package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"serverless/go-common/database"
	"serverless/go-common/logger"
)

var Version = "0.1.0"

type orm struct {
	PostgreSQL postgreSQL
}

type postgreSQL struct {
	Marketing *gorm.DB
}

func New() *orm {
	orm := &orm{}
	dsn := database.GetDSN()
	var err error
	if orm.PostgreSQL.Marketing, err = gorm.Open("postgres", dsn.PostgreSQL); err != nil {
		logger.Panic(err)
	}
	return orm
}

func (o *orm) Close() {
	o.PostgreSQL.Marketing.Close()
}
