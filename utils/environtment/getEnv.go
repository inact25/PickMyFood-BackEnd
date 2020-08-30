package environtment

import (
	"log"

	"github.com/spf13/viper"
)

func Get(key, defVal string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Keyname : %v, not found, default key value : %v, has been loaded", key, defVal)
		//log.Println("Default Config has been loaded")

		return defVal
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
