package model

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	db "gitlab.meitu.com/gocommons/database/mysql"
	"go-test/env"
	"reflect"
)

type SingleDb struct {
	TableName      string
	Template       interface{}
	configFileName string
}

/**
*warn:密码为空，会报未选择数据库。
*待优化
*todo : 使用链表来保存从库链接
*/
func (singleDb *SingleDb) getConn(isMaster int) db.DB {
	config := env.GetMysqlConfig((*singleDb).configFileName)
	var curConfig env.DBTpl
	if isMaster != 0 {
		curConfig = config.Master
	} else {
		curConfig = config.Slave_list[0]
	}
	dbConf := new(mysql.Config)
	dbConf.User = curConfig.Username
	dbConf.Passwd = curConfig.Password
	dbConf.DBName = curConfig.Dbname
	dbConf.Addr = curConfig.Host + ":" + curConfig.Port
	dbConf.Net = "tcp"
	dsn := dbConf.FormatDSN()
	if isMaster != 0 {
		masterDb, err := db.Open(dsn)
		if err != nil {
			fmt.Println(err)
		}
		return masterDb
	} else {
		slaveDb, err := db.Open(dsn)
		if err != nil {
			fmt.Println(err)
		}
		return slaveDb
	}
}

func (singleDb *SingleDb) Select() []interface{} {

	var values []interface{}

	query, err := (*singleDb).getConn(0).Query("select * from " + (*singleDb).TableName)
	if err != nil {
		fmt.Println(err)
		return values
	}

	for query.Next() {
		ttt := reflect.New(reflect.TypeOf((*singleDb).Template)).Interface()
		v := reflect.ValueOf(ttt).Elem()
		var scans []interface{}

		for i := 0; i < v.NumField(); i++ {
			addr := v.Field(i).Addr().Interface()
			scans = append(scans, addr)
		}
		query.Scan(scans...) //返回的是[]uint8,[]uint8,[]uint8
		values = append(values, ttt)
	}

	for _, v := range values {
		t := v.(*TestTemplate)
		fmt.Println(t)
	}
	return values
}

func (singleDb *SingleDb) Update(where map[string]interface{}) {

}
