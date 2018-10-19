package env

import (
	"encoding/json"
	"io/ioutil"
)

type Conf struct {
	Master     MysqlConf   `json:"master"`
	Slave_list []MysqlConf `json:"slave_list"`
}

type MysqlConf struct {
	Master     DBTpl   `json:"master"`
	Slave_list []DBTpl `json:"slave_list"`
}

type DBTpl struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Charset  string `json:"charset"`
	Timeout  int    `json:"timeout"`
}

func (mysqlConf MysqlConf) GetConfig(filePath string) MysqlConf {
	fileContent, err := ioutil.ReadFile(filePath)
	mysqlCnf := MysqlConf{}
	if err != nil {
		return mysqlCnf;
	}
	json.Unmarshal(fileContent, &mysqlCnf)
	return mysqlCnf
}
