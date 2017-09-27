# Milli
Traffic Generator using Alexa's top 1 Million Websites

## Overview

This was my first GoLang project when I wanted to start playing with goroutines.  When I wrote this in 2015, I needed to generate traffic to exercise a NetFlow collector with diverse destinations.  Today I use Milli to generate web requests for creating [Nanolog Streaming Service](https://www.zscaler.com/resources/data-sheets/zscaler-nanolog-streaming-service.pdf) entries, which then feed into various SIEM facilities. 

### Install Dependencies 

```bash
$ go get github.com/parnurzeal/gorequest
```

### Install Milli

```bash
$ git clone https://github.com/eparra/milli.git
$ cd milli/
$ go build milli.go
```

### Download Alexa Data (Top 1 Million Websites)
```bash
$ curl -O http://s3.amazonaws.com/alexa-static/top-1m.csv.zip
$ unzip top-1m.csv.zip 
$ ./milli
```

### To-Do List

* Add CLI arguments
* Better logging (with timestamps)
* Add total bandwidth consumed

### License

Milli is MIT License.
