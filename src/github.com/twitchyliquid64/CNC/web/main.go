package web

import (
  "github.com/twitchyliquid64/CNC/web/httpgateway"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "github.com/hoisie/web"
)

func Run() {
  httpgateway.Init()
  logging.Info("web", "Initialised server on ", config.All().Web.Listener)
  web.RunTLS(config.All().Web.Listener, config.TLS())
  //web.Run(config.All().Web.Listener)
}
