package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type App struct {
	Port      string  `json:"port"`
	JWTSecret string  `json:"jwtSecret"`
	Admins    []Admin `json:"admins"`
}

type Firebase struct {
	TopicName                    string `json:"topicName"`
	NotificationIconURL          string `json:"notificationIconUrl"`
	NotificationClickedTargetURL string `json:"notificationClickedTargetUrl"`
}

type Twitter struct {
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	TextToTrack       string `json:"textToTrack"`
}

type MySQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Config struct {
	App      App      `json:"App"`
	Firebase Firebase `json:"Firebase"`
	Twitter  Twitter  `json:"Twitter"`
	MySQL    MySQL    `json:"mysql"`
}

func NewConfig(confFilePath string) (*Config, error) {
	buf, err := ioutil.ReadFile(confFilePath)
	if err != nil {
		return nil, fmt.Errorf("reading configuration file failed: %v", err)
	}

	config := &Config{}
	if err = json.Unmarshal(buf, config); err != nil {
		return nil, fmt.Errorf("parsing configuration file failed: %v", err)
	}

	return config, nil
}
