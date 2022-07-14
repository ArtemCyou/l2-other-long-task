package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Cfg struct {
	Port string
	Host string
}

func LoadAndStoreConfig() Cfg  {
	v := viper.New() //создаем экземпляр нашего ридера для Env
	v.SetEnvPrefix("SERV") // чтобы не создавать коллизий с переменными окружения которые уже есть
	v.SetDefault("HOST", "localhost:")
	v.SetDefault("PORT",  "8080")
	v.AutomaticEnv() //собираем наши переменные с системных

	var cfg Cfg
	err := v.Unmarshal(&cfg) //закидываем переменные в cfg после анмаршалинга
	if err != nil {
		log.Panic(err)
	}

	return cfg
}

func (c *Cfg)GetServConfig() string {
	return fmt.Sprintf("%s%s",c.Host,c.Port)
}