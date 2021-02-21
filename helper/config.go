package helper

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var resp = Response{}

type conf struct {
	ServerPort int `yaml:"server_port" json:"server_port" `

	LogFilePath string `yaml:"log_file_path" `

	MysqlHost string `yaml:"mysql_host" json:"mysql_host"`
	MysqlUsr  string `yaml:"mysql_usr" json:"mysql_usr"`
	MysqlPwd  string `yaml:"mysql_pwd" json:"mysql_pwd"`
	MysqlDB   string `yaml:"mysql_db" json:"mysql_db"`

	CROSEnabled bool `yaml:"cros_enabled" json:"cros_enabled"`

	ManageRouter string `yaml:"manage_router" json:"manage_router"`
	AdminRoot    string `yaml:"admin_root" json:"admin_root"`
	AdminPass    string `yaml:"admin_pass" json:"admin_pass"`

	SMTPEnabled  bool   `yaml:"smtp_enabled" json:"smtp_enabled"`
	SMTPHost     string `yaml:"smtp_host" json:"smtp_host"`
	SMTPPort     int    `yaml:"smtp_port" json:"smtp_port"`
	SMTPUsername string `yaml:"smtp_username" json:"smtp_username"`
	SMTPPassword string `yaml:"smtp_password" json:"smtp_password"`
	SMTPTo       string `yaml:"smtp_to" json:"smtp_to"`

	SensitivePath string `yaml:"sensitive_path"`
	IPBlockPath   string `yaml:"ip_block_path"`
}

// Config 配置内容
var Config conf

var configPath string = "./config/config.yaml"

// Setup 装载配置
func init() {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, &Config)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

// SaveConfigFile is save config
func SaveConfigFile(c *gin.Context) {
	var config conf
	err := c.BindJSON(&config)
	if err != nil {
		fmt.Println(err.Error())
		resp.Error(c, ResponseParamError, "入参错误")
		return
	}

	config.LogFilePath = "./logs/"
	config.SensitivePath = "./config/sensitive.txt"
	config.IPBlockPath = "./config/block_ip.txt"
	yamlFile, err := yaml.Marshal(config)
	if err != nil {
		resp.Error(c, ResponseServerError, "生成错误")
		return
	}
	err = ioutil.WriteFile(configPath, yamlFile, 0666)
	if err != nil {
		resp.Error(c, ResponseServerError, "保存错误")
		return
	}
	resp.Success(c, true)
}
