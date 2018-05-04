# Cli log handler

For short running cli and human read

## Usage

````go
package main

import (
	"os"
	"time"
	
	dlog "github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/log/handlers/cli"
)

var log = dlog.NewApplicationLogger()

func main() {
	dlog.SetHandlerRecursive(log, cli.New(os.Stderr, true))
	log.Info("hi")
	log.Infof("open file %s", "foo.yml")
	log.InfoF("open", dlog.Fields{
		dlog.Str("file", "foo.yml"),
		dlog.Int("mode", 0666),
	})
	log.Warn("I am yellow")
	func() {
		defer func() {
			if r := recover(); r != nil {
				log.Info("recovered", r)
			}
		}()
		log.Panic("Panic reason")
	}()
	dlog.SetLevelRecursive(log, dlog.DebugLevel)
	log.Debug("I will sleep for a while")
	time.Sleep(1 * time.Second)
	log.Fatal("I am red")
}
````

Result in

````text
INFO 0000 hi
INFO 0000 open file foo.yml
INFO 0000 open file=foo.yml mode=438
WARN 0000 I am yellow
PANI 0000 Panic reason
INFO 0000 recoveredPanic reason
DEBU 0000 I will sleep for a while
FATA 0001 I am red
````

It has color (only tested on Linux (Ubuntu)) and can't be disabled

![gommon_log_cli_handler](gommon_log_cli_handler.png)