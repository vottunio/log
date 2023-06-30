# log

[![Go Reference](https://pkg.go.dev/badge/github.com/vottunio/log.svg)](https://pkg.go.dev/github.com/vottunio/log)
[![Build Status](https://travis-ci.org/vottunio/log.svg?branch=main)](https://travis-ci.org/vottunio/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/vottunio/log)](https://goreportcard.com/report/github.com/vottunio/log)

Vottun Log is a lightweight and flexible logging library for GoLang. It provides a simple interface for logging messages with different levels of severity. The library offers various features, including log level filtering, structured logging, and log output customization.

## Installation

To use Vottun Log in your Go project, you need to have Go installed and set up. Once you have Go ready, you can install the package by running the following command:


```go
go get github.com/vottunio/log
```

## Usage

Here's a basic example of using Vottun Log in your Go code:

```go
package main

import (
	"github.com/vottunio/log"
)

func main() {
	log.Debugln("This is a debug message")
	log.Infoln("This is an informational message")
	log.Warnln("This is a warning message")
	log.Errorln("This is an error message")
	log.Traceln("This is an trace message")
}
```

And the result is:
<div style="background: rgba(110, 118, 129, 0.4); padding:16px">
08:01:18.750460 <span style="color:magenta">DEBUG [main] This is a debug message</span><br/>
08:01:18.750460 <span style="color:green">INFO [main] This is an info message</span><br/>
08:01:18.750460 <span style="color:yellow">WARN [main] This is a warn message</span><br/>
08:01:18.750460 <span style="color:red">ERROR [main] This is an error message</span><br/>
08:01:18.750460 <span style="color:cyan">TRACE [main] This is a trace message</span><br/>
</div>

## License

Vottun Log is licensed under the [MIT License](https://github.com/vottunio/log/blob/main/LICENSE).

Visit us at [https://vottun.com](https://vottun.com)