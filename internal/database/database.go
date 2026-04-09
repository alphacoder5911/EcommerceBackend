package database

import (
	"fmt"

	"github.com/alphacoder5911/EcommerceBackend/internal/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func New(cfg config.DatabaseConfig)(*gorm.DB,error){

	var dsn string

	if cfg.URL!=""{
		log.Info().Msg("Url retrieved ")
		dsn=cfg.URL
	}else{

	dsn=fmt.Sprintf("host=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC User=%s",cfg.User,
cfg.Host,cfg.Password,cfg.Name,cfg.Port,cfg.SSLMode)
	}


	db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err!=nil{
		return nil,fmt.Errorf("Failed to connect to the database %w",err)
	}

	return db,nil

}