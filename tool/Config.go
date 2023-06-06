package tool

import (
	"bufio"
	"github.com/goccy/go-json"
	"os"
)

// 总配置
type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Sms      SmsConfig      `json:"sms"`
	DataBase DataBaseConfig `json:"database"`
	Redis    RedisConfig    `json:"redis"`
}

// 短信验证码配置
type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	RegionId     string `json:"region_id"`
}

// 数据库配置
type DataBaseConfig struct {
	Driver    string `json:"driver"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	DbName    string `json:"db_name"`
	Charset   string `json:"charset"`
	ParseTime string `json:"parse_time"`
	Loc       string `json:"loc"`
}

// redis配置
type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

var _cfg *Config = &Config{}

// 获取配置信息
func GetConfig() *Config {
	return _cfg
}

// 通过PareConfig方法读取对应路径的JSON文件，转化为结构体
func PareConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	newDecoder := json.NewDecoder(reader)
	err = newDecoder.Decode(_cfg)
	if err != nil {
		return nil, err
	}
	return _cfg, nil
}
