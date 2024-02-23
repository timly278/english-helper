package session

import "time"

const (
	MODE_MANUAL = "manual"
	MODE_RANDOM = "random"
)

type Story struct {
	Path     string        `mapstructure:"path"`
	Duration time.Duration `mapstructure:"duration"`
}

type DailyJournal struct {
	Path     string        `mapstructure:"path"`
	Duration time.Duration `mapstructure:"duration"`
}

type DescribeImage struct {
	Path     string        `mapstructure:"path"`
	Duration time.Duration `mapstructure:"duration"`
}

type AnswerQuestion struct {
	Path           string        `mapstructure:"path"`
	NumOfQuestions int           `mapstructure:"numofquestions"`
	Duration       time.Duration `mapstructure:"duration"`
}

type VocabSentence struct {
	Path     string        `mapstructure:"path"`
	Duration time.Duration `mapstructure:"duration"`
}

type AlphabetSentence struct {
	Path     string        `mapstructure:"path"`
	Duration time.Duration `mapstructure:"duration"`
}
