package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	es "github.com/olivere/elastic/v7"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"my/config"
)

var (
	GVA_DB            *gorm.DB
	GVA_REDIS         *redis.Client
	GVA_CONFIG        config.Server
	GVA_VP            *viper.Viper
	GVA_LOG           *oplogging.Logger
	GVA_ES            *es.Client
	GVA_BulkProcessor *es.BulkProcessor
)
