package main

import (
	// "github.com/twitchyliquid64/CNC/signaller"
	"github.com/twitchyliquid64/CNC/messenger"
	"github.com/twitchyliquid64/CNC/registry"
	"github.com/twitchyliquid64/CNC/logging"
	"github.com/twitchyliquid64/CNC/plugin"
	"github.com/twitchyliquid64/CNC/config"
	"github.com/twitchyliquid64/CNC/data"
	"github.com/twitchyliquid64/CNC/web"
	"os/signal"
	"syscall"
	"time"
	"os"
)


func run(stopSignal chan bool) {
	logging.Info("main", "Starting server")
	if config.Load(getConfigPath()) != nil {
		logging.Fatal("main", "Configuration error")
	}
	logging.Info("main", "Configuration read, now starting '", config.GetServerName(), "'")
	registry.Initialise()
	data.Initialise()

	messenger.Initialise()
	plugin.Initialise()

	web.Initialise()
	go web.Run()

	for {
		select {
		case <- stopSignal:
			logging.Info("main", "Got stop signal, finalizing now")
			//signaller.Stop()
			return
		default:
			time.Sleep(time.Millisecond * 400)
		}
	}
}


func getConfigPath()string{
	if len(os.Args) < 2{
		return "config.json"
	}
	return os.Args[1]
}

func main() {
	processSignal := make(chan os.Signal, 1)
	signal.Notify(processSignal, os.Interrupt, os.Kill, syscall.SIGHUP)
	chanStop := make(chan bool)
	shouldRun := true

	go func(){//goroutine to monitor OS signals
		for{
			s := <- processSignal //wait for signal from OS
			logging.Info("main", "Got OS signal: ", s)
			if s != syscall.SIGHUP{
				shouldRun = false
			}
			chanStop <- true
		}
	}()

	run(chanStop)//will run until signalled to stop from above goroutine
	time.Sleep(time.Millisecond * 100)
}
