package session

import "github.com/spf13/viper"

type Config struct {
	Mode             string           `mapstructure:"mode"`
	Story            Story            `mapstructure:"story"`
	DailyJournal     DailyJournal     `mapstructure:"daily-journal"`
	DescribeImage    DescribeImage    `mapstructure:"describe-image"`
	AnswerQuestion   AnswerQuestion   `mapstructure:"answer-question"`
	VocabSentence    VocabSentence    `mapstructure:"vocab-sentence"`
	AlphabetSentence AlphabetSentence `mapstructure:"alphabet-sentence"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") // you can use json/xml here if you want so as long it has correct format

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
