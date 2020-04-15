package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/DataDog/datadog-go/statsd"
	"io/ioutil"
	"os"
	"regexp"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Patterns []*Pattern `yaml:"patterns"`
}

type Pattern struct {
	Pattern string `yaml:"pattern"`
	Regex *regexp.Regexp
	Metric string `yaml:"metric"`
	Tags []string `yaml:"tags"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	// parse flags
	config_path := flag.String("config_path", "logs_to_datadog_metrics.yaml", "path to config_path file")
	flag.Parse()

	// read config_path
	var config Config

	content, err := ioutil.ReadFile(*config_path)
	check(err)

	err = yaml.Unmarshal(content, &config)
	check(err)

	for _, c := range config.Patterns {
		c.Regex = regexp.MustCompile(c.Pattern)
	}

	// connect to statsd
	statsd, err := statsd.New(os.Getenv("STATSD_HOST") + ":" + os.Getenv("STATSD_PORT"))
	check(err)
	defer func() {
		statsd.Flush() // https://github.com/DataDog/datadog-go/issues/138
		statsd.Close()
	}()

	// read stdin and send metrics
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		fmt.Println(string(line))
		for _, c := range config.Patterns {
			if c.Regex.Match(line) {
				statsd.Incr(c.Metric, c.Tags, 1)
			}
		}
	}
}
