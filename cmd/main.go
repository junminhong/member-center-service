package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	deliverGrpc "github.com/junminhong/member-center-service/api/v1/delivery/grpc"
	deliver "github.com/junminhong/member-center-service/api/v1/delivery/http"
	"github.com/junminhong/member-center-service/api/v1/delivery/http/middleware"
	repo "github.com/junminhong/member-center-service/api/v1/repository"
	"github.com/junminhong/member-center-service/api/v1/usecase"
	_ "github.com/junminhong/member-center-service/docs"
	"github.com/junminhong/member-center-service/domain"
	sugarLogger "github.com/junminhong/member-center-service/pkg/logger"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net"
	"os"
)

var zapLogger = sugarLogger.Setup()

func init() {
	// get now work dir
	path, err := os.Getwd()
	if err != nil {
		zapLogger.Error(err.Error())
	}
	// setting viper get config yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		zapLogger.Error(err.Error())
	}
	if viper.GetString("APP.GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}

type postgresDB struct {
	db *gorm.DB
}

func setUpDB() *postgresDB {
	// sslmode=disable
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Taipei",
		viper.GetString("APP.DB_HOST"),
		viper.GetString("APP.DB_USERNAME"),
		viper.GetString("APP.DB_PASSWORD"),
		viper.GetString("APP.DB_DATABASE"),
		viper.GetString("APP.DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		zapLogger.Error("Failed to connect DB")
	}
	return &postgresDB{db: db}
}

func setUpRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("APP.REDIS_HOST") + ":" + viper.GetString("APP.REDIS_PORT"),
		Password: viper.GetString("APP.REDIS_PASSWORD"), // no password set
		DB:       0,                                     // use default DB
	})
	return client
}

func setUpDomain(router *gin.Engine, server *grpc.Server, lis net.Listener, db *gorm.DB, redis *redis.Client) {
	memberRepo := repo.NewMemberRepo(db, redis, zapLogger)
	memberUseCase := usecase.NewMemberUseCase(memberRepo, zapLogger)
	authRepo := repo.NewAuthRepo(db, redis, zapLogger)
	authUseCase := usecase.NewAuthUseCase(authRepo, memberRepo, zapLogger)
	deliver.NewMemberHandler(router, memberUseCase, memberRepo)
	deliver.NewAuthRepo(router, authUseCase, authRepo)
	deliverGrpc.NewMemberGrpc(server, lis, memberRepo)
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Middleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)))
	go router.Run(viper.GetString("HOST") + ":" + viper.GetString("APP.PORT"))
	return router
}

func setupGRPC() (*grpc.Server, net.Listener) {
	zapLogger.Info("starting gRPC server...")
	zapLogger.Info("Listening and serving HTTP on :" + viper.GetString("APP.HOST") + ":" + viper.GetString("APP.GRPC_PORT"))
	gRpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", viper.GetString("APP.HOST")+":"+viper.GetString("APP.GRPC_PORT"))
	if err != nil {
		zapLogger.Info("failed to listen: %v \n", err)
	}
	return gRpcServer, lis
}

func (postgresDB *postgresDB) migrationDB() {
	err := postgresDB.db.AutoMigrate(&domain.Member{}, &domain.MemberInfo{})
	if err != nil {
		zapLogger.Error(err.Error())
	}
}

// @title           Member Center Service API
// @version         1.0
// @description     This is a base golang develop member center service

// @contact.name   junmin.hong
// @contact.url    https://github.com/junminhong
// @contact.email  junminhong1110@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:9200
// @BasePath  /api/v1
func main() {
	db := setUpDB()
	go db.migrationDB()
	redis := setUpRedis()
	router := setUpRouter()
	grpcServer, lis := setupGRPC()
	setUpDomain(router, grpcServer, lis, db.db, redis)
}
