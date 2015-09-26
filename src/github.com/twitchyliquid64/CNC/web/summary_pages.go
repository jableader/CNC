package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/hoisie/web"
)

func dashboardMainPage(ctx *web.Context) {
  isLoggedIn, user, session := getSessionByCookie(ctx)
  logging.Info("web", isLoggedIn, user, session)
  if !isLoggedIn {
    ctx.Redirect(302, "/login")
    return
  }

  t := templates.Lookup("dashboardindex")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, nil)
}
