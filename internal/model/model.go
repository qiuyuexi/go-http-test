package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"encoding/json"
)

type T struct {
	Id int
	T  int
}

type TTT struct {
	Id   int    `json:"id"`
	T    int    `json:"t"`
	Test string `json:"test"`
}

var Db sql.DB

func init() {

}

func GetConn() {
	dsn := "root:@/test?charset=utf8"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("db connect error")
		panic(err)
	}
	defer conn.Close()
	conn.Exec("use test")
	query, queryErr := conn.Query("SELECT * FROM t")

	if queryErr != nil {
		fmt.Println("db select error")
		panic(queryErr)
	}

	cols, err := query.Columns()
	fmt.Println(cols)

	for query.Next() {
		t := T{}
		mp := reflect.New(reflect.TypeOf(t)).Interface() //Elem() is not a pointer or an interface reflect.Value.Elem() will panic
		structV := reflect.ValueOf(mp).Elem()
		var addrList []interface{}
		for i := 0; i < structV.NumField(); i++ {
			addr := structV.Field(i).Addr().Interface()
			addrList = append(addrList, addr)
		}

		query.Scan(addrList...)
		fmt.Println(addrList)

		for k, v := range addrList {
			t := v.(*int)
			fmt.Println(cols[k], *t)
		}

	}

}

func Query() {

	dsn := "root:@/test?charset=utf8"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("db connect error")
		panic(err)
	}
	defer conn.Close()
	conn.Exec("use test")
	query, queryErr := conn.Query("SELECT * FROM t")

	if queryErr != nil {
		fmt.Println("db select error")
		panic(queryErr)
	}

	cols, err := query.Columns()
	fmt.Println(cols)
	var values []map[interface{}]interface{}

	vals := make([]byte, len(cols)) //存在字符串，需要二维数组
	scans := make([]interface{}, len(cols))
	//fmt.Println(vals)

	for i := range vals {
		//fmt.Println(i)
		scans[i] = &vals[i]
		fmt.Println(scans[i])
	}

	for query.Next() {

		query.Scan(scans...) //返回的是[]uint8,[]uint8,[]uint8
		fmt.Println(scans)

		row := make(map[interface{}]interface{})
		for k, v := range vals {
			fmt.Println(reflect.TypeOf(v).String(), v)
			row[cols[k]] = v
		}
		values = append(values, row)
	}
	fmt.Println(values)
	for kk, vv := range values {
		fmt.Println(kk, vv)
	}
}

func QueryReturnStruct() {
	dsn := "root:@/test?charset=utf8"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("db connect error")
		panic(err)
	}
	defer conn.Close()
	conn.Exec("use test")
	query, queryErr := conn.Query("SELECT * FROM t")

	if queryErr != nil {
		fmt.Println("db select error")
		panic(queryErr)
	}
	var values []interface{}

	for query.Next() {
		ttt := reflect.New(reflect.TypeOf(*new(TTT))).Interface()
		v := reflect.ValueOf(ttt).Elem()
		var scans []interface{}

		for i := 0; i < v.NumField(); i++ {
			addr := v.Field(i).Addr().Interface()
			scans = append(scans, addr)
		}
		query.Scan(scans...) //返回的是[]uint8,[]uint8,[]uint8
		values = append(values, ttt)
	}
	jsonByte, err := json.Marshal(values[0])
	fmt.Println(string(jsonByte))
}
