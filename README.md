# Milli
Traffic Generator using Alexa's top 1 Million Websites

## Overview

This is my first GoLang project.  I needed to generate traffic to exercise a NetFlow collector with diverse destinations.  I also wanted to start playig with goroutines, hence this project.   

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
$ ./milli.go
```

### To-Do List

* Add CLI arguments
* Better logging (with timestamps)
* Add total bandwidth consumed

### License

Milli is MIT License.
