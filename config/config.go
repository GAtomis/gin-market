package config

import "fmt"

const (
	SQL_URL = "43.143.244.182"
	// SQL_URL    = "127.0.0.1"
	SQL_NAME   = "vue_artemis_admin"
	SQL_CONFIG = "?charset=utf8mb4&parseTime=True&loc=Local"
	USERNAME   = "root"
	PASSWORD   = "QSCesz2rfx@"
	SQL_PORT   = "3306"
	PROTOCOL   = "tcp"
	TIME_OUT   = "30s"
)

/**
 * @description: 获取sql连接
 * @return {*}
 * @Date: 2022-07-20 14:06:48
 */
func GetDsn() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@%s(%s:%s)/%s%s&%s", USERNAME, PASSWORD, PROTOCOL, SQL_URL, SQL_PORT, SQL_NAME, SQL_CONFIG, TIME_OUT)
	fmt.Println(dsn)
	return
}
