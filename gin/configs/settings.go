package configs

import (
	"myHome/gin/utils"
	"time"
)

// 服务器配置
type serverSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//数据库配置
type databaseSettings struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// Redis setting
type redisSettings struct {
	Addr     string
	Password string
	DB       int
}

// Postgresql Setting
type postgreSettings struct {
	DBUrl string
}

// Postgresql ORM
type postgreOrms struct {
	DSN string
}

// MiniProgram
type miniProSettings struct {
	AppID     string
	AppSecret string
}

//定义全局变量
var (
	Server          *serverSettings
	databaseSetting *databaseSettings
	redisSetting    *redisSettings
	postgreSetting  *postgreSettings
	postgreOrm      *postgreOrms
	miniProSetting  *miniProSettings
)

// SetupSetting 读取配置到全局变量
func setupSetting() error {
	s, err := NewSetting()
	utils.CheckErr(err)

	/*err = s.ReadSection("Database", &databaseSetting)
	utils.CheckErr(err)*/

	err = s.ReadSection("Server", &Server)
	utils.CheckErr(err)

	err = s.ReadSection("Redis", &redisSetting)
	utils.CheckErr(err)

	err = s.ReadSection("PostgreSQL", &postgreSetting)
	utils.CheckErr(err)

	err = s.ReadSection("PostgreORM", &postgreOrm)
	utils.CheckErr(err)

	err = s.ReadSection("MiniProgramTest", &miniProSetting) // Test
	// err = s.ReadSection("MiniProgram", &miniSetting)  // Prod
	utils.CheckErr(err)
	/*fmt.Println("setting:")
	fmt.Println(ServerSetting)
	fmt.Println(DatabaseSetting)
	fmt.Println(RedisSetting)*/

	return err
}
