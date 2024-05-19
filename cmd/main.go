package main

import (
	"goCourseService/internal/repository"
	"goCourseService/internal/service"
	"goCourseService/internal/handler"
	
	"goCourseService/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := initConfig(); err != nil {
		log.Fatal().Err(err).Msg("some error with initializiing")	
	}

	db, err := server.NewPostgresConnection(server.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})

	if err != nil {
		log.Fatal().Err(err).Msg("error with data base")
	}

	redisClient, err := server.NewRedisConnection(
		viper.GetString("redis.port"),
		viper.GetString("redis.password"),
	)

	repos := repository.NewRepository(db, redisClient)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(server.Server)

	if err := srv.RunServer(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal().Err(err).Msg("error with run server")
	}
}

func initConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.BindEnv("db.password", "DB_PASSWORD")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")
	return nil
}