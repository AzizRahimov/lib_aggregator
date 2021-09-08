package setting

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	DepCred ConfDepCred `json:"depcred"`
	HO      ConfWallet  `json:"humo"`
	Server  CondServer  `json:"server"`
}

type ConfDepCred struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Uri      string `json:"uri"`
}

type ConfWallet struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	SecretKey string `josn:"secretKey"`
	Uri       string `json:"uri"`
	Timeout   int    `json:"timeout"`
}

type CondServer struct {
	Token string `json:"token"`
	Port  string `json:"port"`
}

var AppConfig = &Config{}

//path будет стринг -> но вот Config Interface ( Главный вопрос) от куда я эту структуру возьму?
func Setup(path string, conf interface{}) (err error) {
//	mapTo("conf/config.json", AppConfig)
	if path == ""{
		log.Println("path can't be empty")
		return err

	}
	err =  mapTo(path, conf)
	if err != nil {
		return err
	}
	//log.Println(AppConfig)

	return nil
}

// mapTo map section
func mapTo(F string, v interface{}) (err error) {

	byteValue, err := ioutil.ReadFile(F)
	if err != nil {
		//log.Fatalf("%v", err)
		return err
	}

	err = json.Unmarshal(byteValue, &v)
	if err != nil {
		//log.Fatalf("%v", err)
		return err
	}
	return nil
}
