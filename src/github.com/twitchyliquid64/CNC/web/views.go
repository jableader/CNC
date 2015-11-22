package web

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/hoisie/web"
)

func loginMainPage(ctx *web.Context) {
  t := templates.Lookup("login")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, nil)
}

func pluginAdminListPage_view(ctx *web.Context) {
  t := templates.Lookup("pluginlist")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, nil)
}

func usersAdminMainPage_view(ctx *web.Context) {
  t := templates.Lookup("userpage")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, nil)
}

func entityAdminViewerPage_view(ctx *web.Context) {
  t := templates.Lookup("adminentityviewer")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, nil)
}

func entityAdminForm_view(ctx *web.Context) {
  t := templates.Lookup("adminentityform")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, nil)
}

func entityViewerPage_view(ctx *web.Context) {
  t := templates.Lookup("entityviewer")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, nil)
}

func dashboardSummary_view(ctx *web.Context) {
  t := templates.Lookup("dashboardsummary")
	if t == nil {
		logging.Error("web", "No template found.")
	}
	t.Execute(ctx.ResponseWriter, nil)
}

func templateReloadHandler(ctx *web.Context) {
  templateReInit()
}
