package configs

import (
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"github.com/silenceper/wechat/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"myHome/gin/configs/nacos"
	"myHome/gin/utils/logs"
	"os"
	"os/signal"

	// "github.com/jackc/pgx/v4/pgxpool"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"myHome/gin/utils"
)

var (
	RedisStore *redis.Client
	PGStore    *gorm.DB
	MiniPro    *miniprogram.MiniProgram
)

func InitSettings() {

	// init yaml config must be first
	err := setupSetting()
	logs.Error(err).Msg("Setup yaml config settings")

	// init log system
	logs.InitLogger(zerolog.TraceLevel)

	// handler ctrl+c sign
	SetupCLearHandler()

	// nacos
	nacos.RegisterNacosService()

	// init miniprogram
	initMiniProgram()

	if RedisStore == nil {
		RedisStore = redis.NewClient(&redis.Options{
			Addr:     redisSetting.Addr,
			Password: redisSetting.Password,
			DB:       redisSetting.DB,
		})
	}

	/*if PGCon == nil {
		PGCon, err = pgx.Connect(ctx, postgreSetting.DBUrl)
		utils.CheckErr(err)
	}*/

	if PGStore == nil {
		PGStore, err = gorm.Open(postgres.Open(postgreOrm.DSN), &gorm.Config{})
		utils.CheckErr(err)
	}

}

func SetupCLearHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		sig := <-c
		logs.Info().Msgf("system been interrupted by %s", sig)
		ReleaseSettings()
		os.Exit(0)
	}()
}

func ReleaseSettings() {

	nacos.DeregisterNacosService()

	if RedisStore != nil {
		err := RedisStore.Close()
		utils.CheckErr(err)
	}
	/*if PGCon != nil && !PGCon.IsClosed()  {
		err := PGCon.Close(ctx)
		utils.CheckErr(err)
	}*/
	if PGStore != nil {
		db, err := PGStore.DB()
		utils.CheckErr(err)
		err = db.Close()
		utils.CheckErr(err)
	}
}

func initMiniProgram() {
	if MiniPro == nil {
		wc := wechat.NewWechat()

		miniCfg := getMiniConfig()
		MiniPro = wc.GetMiniProgram(miniCfg)
	}
}

func getMiniConfig() *miniConfig.Config {

	memory := cache.NewMemory()

	miniCfg := &miniConfig.Config{
		AppID:     miniProSetting.AppID,
		AppSecret: miniProSetting.AppSecret,
		Cache:     memory,
	}

	return miniCfg
}
