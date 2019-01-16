package config

import (
    "../../common/utils/pathUtils"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "path/filepath"
)

type DbJson struct {
    Mysqljson MysqlJson `json:"mysql"`
    Redisjson RedisJson `json:"redis"`
}

type MysqlJson struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
    Db string `json:"db"`
}

type RedisJson struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Auth     string    `json:"auth"`
    Db     int    `json:"db"`
}

func GetMysqlConfig() (host string, port int, username string, password string, db string) {
    mysqlConfig, _ := databaseConfig()
    return mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Db
}

func GetRedisConfig() (host string, port int, auth string, db int) {
    _, redisConfig := databaseConfig()
    return redisConfig.Host, redisConfig.Port, redisConfig.Auth, redisConfig.Db
}

func databaseConfig() (mysqlJson MysqlJson, redisJson RedisJson) {
    b, readErr := ioutil.ReadFile(getConfigFile())
    if readErr != nil {
        fmt.Println("load database config error: ", readErr.Error())
    }
    //fmt.Println(string(b))
    var config DbJson
    if err := json.Unmarshal(b, &config); err != nil {
        fmt.Println("unmarshal database json error: ", err.Error())
    }

    return config.Mysqljson, config.Redisjson
}

func getConfigFile() string {
    rootPath := pathUtils.GetParentPath(pathUtils.GetParentPath(pathUtils.GetParentPath(pathUtils.GetFilePath())))
    configFile := filepath.Join(rootPath, "resources/conf/database.json")
    return configFile
}
