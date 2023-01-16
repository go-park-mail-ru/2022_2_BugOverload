package main

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"

	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

type performanceTestConfig struct {
	// Global
	Rate           int    `toml:"rate"`
	Duration       int    `toml:"duration"`
	NameTesting    string `toml:"name_testing"`
	TimeoutTesting int    `toml:"timeout_testing"`

	// Request params
	RequestMaxFilmID     int `toml:"request_max_film_id"`
	RequestMaxFilmImages int `toml:"request_max_film_images"`

	RequestMaxFilmReviews     int `toml:"request_max_film_reviews"`
	RequestMaxFilmReviewsPage int `toml:"request_max_film_reviews_page"`

	RequestMaxFilmSimilar int `toml:"request_max_film_similar"`

	RequestMaxFilmInTag  int     `toml:"request_max_film_in_tag"`
	RequestMaxFilmRating float64 `toml:"request_max_film_in_tag_delimiter"`

	// Targets
	Domain         string   `toml:"domain"`
	MethodsTargets []string `toml:"methods_targets"`
	CountTargets   []int    `toml:"count_targets"`
	URLTargets     []string `toml:"url_targets"`
}

func BasePerformanceTest() {
	var config performanceTestConfig

	_, err := toml.DecodeFile("./perf_test/base.toml", &config)
	if err != nil {
		logrus.Fatal("No config file")
	}

	targets := GetTestingData(&config)

	rate := vegeta.Rate{Freq: config.Rate, Per: time.Second}
	duration := time.Duration(config.Duration) * time.Second

	targeter := vegeta.NewStaticTargeter(targets...)
	attacker := vegeta.NewAttacker()

	go func() {
		time.Sleep(time.Duration(config.TimeoutTesting) * time.Second)

		attacker.Stop()
	}()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, config.NameTesting) {
		metrics.Add(res)
	}
	metrics.Close()

	GetReport(metrics)
}

func GetTestingData(config *performanceTestConfig) []vegeta.Target {
	globalCountRequest := 0

	for _, value := range config.CountTargets {
		globalCountRequest += value
	}

	targets := make([]vegeta.Target, globalCountRequest)

	pos := 0

	curTypeTarget := 0

	curURL := make([]string, len(config.URLTargets))

	for i := 0; i < config.CountTargets[curTypeTarget]; i++ {
		curURL[curTypeTarget] = fmt.Sprintf(
			config.Domain+config.URLTargets[curTypeTarget],
			pkg.RandMaxInt(config.RequestMaxFilmID)+1,
			pkg.RandMaxInt(config.RequestMaxFilmImages),
		)

		targets[pos] = vegeta.Target{
			Method: config.MethodsTargets[curTypeTarget],
			URL:    curURL[curTypeTarget],
		}

		curURL[curTypeTarget] = config.MethodsTargets[curTypeTarget] + " " + curURL[curTypeTarget]

		pos++
	}

	curTypeTarget++

	for i := 0; i < config.CountTargets[curTypeTarget]; i++ {
		curURL[curTypeTarget] = fmt.Sprintf(
			config.Domain+config.URLTargets[curTypeTarget],
			pkg.RandMaxInt(config.RequestMaxFilmID)+1,
		)

		targets[pos] = vegeta.Target{
			Method: config.MethodsTargets[curTypeTarget],
			URL:    curURL[curTypeTarget],
		}

		curURL[curTypeTarget] = config.MethodsTargets[curTypeTarget] + " " + curURL[curTypeTarget]

		pos++
	}

	curTypeTarget++

	for i := 0; i < config.CountTargets[curTypeTarget]; i++ {
		curURL[curTypeTarget] = fmt.Sprintf(
			config.Domain+config.URLTargets[curTypeTarget],
			pkg.RandMaxInt(config.RequestMaxFilmID)+1,
			pkg.RandMaxInt(config.RequestMaxFilmReviews)+1,
			pkg.RandMaxInt(config.RequestMaxFilmReviewsPage),
		)

		targets[pos] = vegeta.Target{
			Method: config.MethodsTargets[curTypeTarget],
			URL:    curURL[curTypeTarget],
		}

		curURL[curTypeTarget] = config.MethodsTargets[curTypeTarget] + " " + curURL[curTypeTarget]

		pos++
	}

	curTypeTarget++

	for i := 0; i < config.CountTargets[curTypeTarget]; i++ {
		curURL[curTypeTarget] = fmt.Sprintf(
			config.Domain+config.URLTargets[curTypeTarget],
			pkg.RandMaxInt(config.RequestMaxFilmInTag)+1,
			pkg.RandMaxFloat64(config.RequestMaxFilmRating, 1),
		)

		targets[pos] = vegeta.Target{
			Method: config.MethodsTargets[curTypeTarget],
			URL:    curURL[curTypeTarget],
		}

		curURL[curTypeTarget] = config.MethodsTargets[curTypeTarget] + " " + curURL[curTypeTarget]

		pos++
	}

	logrus.Info("----Testing setup----")
	for i := 0; i < len(curURL); i++ {
		logrus.Infof("Request: %s", curURL[i])
	}

	return targets
}

const (
	percent     = 100.0
	scaleMemory = 1024.0
)

func GetReport(metrics vegeta.Metrics) {
	loadKB := float64(metrics.BytesIn.Total) / scaleMemory

	logrus.Info("\n----Testing results----")
	logrus.Infof("Count request:   %d", metrics.Requests)
	logrus.Infof("99th percentile: %s", metrics.Latencies.P99)
	logrus.Infof("Duration:        %s", metrics.Duration)
	logrus.Infof("Rate:            %g", metrics.Rate)
	logrus.Infof("BiteIn:          Mean: %g | Total: %dB %3.2fKB", metrics.BytesIn.Mean, metrics.BytesIn.Total, loadKB)
	logrus.Infof("BiteOut:         Mean: %g | Total: %d", metrics.BytesOut.Mean, metrics.BytesOut.Total)

	countAnswers := 0.0

	logrus.Info("Status Codes\n")
	statusCodes := metrics.StatusCodes
	logrus.Info("Code | Count")
	for key, value := range statusCodes {
		logrus.Infof("%s  : %d", key, value)
		if key != "0" && key != "500" {
			countAnswers += float64(value)
		}
	}

	logrus.Info("\n----Total----")
	logrus.Infof("Count success (200):              %3.2f %%", metrics.Success*percent)
	logrus.Infof("Count Answers (!500, !0):         %3.2f %%", countAnswers/float64(metrics.Requests)*percent)
	logrus.Infof("RPS (Count Answers / Duration):   %3.2f", countAnswers/metrics.Duration.Seconds())
	logrus.Infof("Read flow (BiteIn / Duration):    %3.2f KB/s", loadKB/metrics.Duration.Seconds())
}

func main() {
	BasePerformanceTest()
}
