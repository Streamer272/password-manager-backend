package database

func Init() {
	ConnectMysql()
	ConnectRedis()
}
