# Logs to Datadog Metrics [![Build Status](https://travis-ci.org/eatwithforks/logs_to_datadog_metrics.svg?branch=master)](https://travis-ci.org/eatwithforks/logs_to_datadog_metrics)
Parse process logs for patterns and send metrics to Datadog 

Lots of programs generate logs but do not generate metrics. This is a simple program that reads from stdin, matches a pattern of your choice, and sends metrics to Datadog.
 ### Installation
 via curl:
 ```bash
curl -sfL <PICK URL FROM RELEASES PAGE> | \
tar -zx && \
chmod +x logs_to_datadog_metrics
```
 
 via go:
  ```bash
 go get github.com/eatwithforks/logs_to_datadog_metrics
 ```
 
 ### Setup
Create a config yaml file with specified patterns.

```
patterns:
- pattern: 'this is bad'
  metric: 'foo.bad'
  tags: ['foo:bar']
```

To use:
```
execute <your-program-here> | STATSD_HOST="localhost" STATSD_PORT=8125 logs_to_datadog_metrics -config_path /path/to/config/file"
```

Note: if important stuff comes out on stderr, you need to add `2>&1` before the pipe.

Todo:
1. Add a way to make pattern matches appear in tags.
