package web

import (
  "github.com/twitchyliquid64/CNC/web/devclienthandler"
  "github.com/twitchyliquid64/CNC/web/pluginsockets"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "golang.org/x/net/websocket"
  "github.com/hoisie/web"
)

// ### THIS FILE SHOULD CONTAIN ALL INITIALISATION CODE FOR BOTH TEMPLATES AND URL HANDLERS ###

func Initialise() {
  logging.Info("web", "Registering page handlers")
  registerCoreHandlers()
  registerUserHandlers()
  registerSummaryHandlers()
  registerEntityHandlers()
  registerPluginHandlers()
  registerWebSockets()
  registerTemplateViews()

  logging.Info("web", "Registering templates")
  registerCoreTemplates()
  registerUserTemplates()
  registerSummaryTemplates()
  registerEntityTemplates()
  registerPluginTemplates()
}

func registerCoreHandlers() {
  web.Get("/login", loginMainPage, config.All().Web.Domain)
  web.Get("/dev/reload", templateReloadHandler, config.All().Web.Domain)
  web.Get("/sys-status", getSysComponentsStatusAPIHandler, config.All().Web.Domain)
  web.Get("/p/(.*)", pluginGeneralHandler, config.All().Web.Domain)
  web.Post("/p/(.*)", pluginGeneralHandler, config.All().Web.Domain)
  web.Get("/ref/api", apiRefView, config.All().Web.Domain)
}

func registerUserHandlers() {
  web.Post("/login", loginHandler, config.All().Web.Domain)
  web.Get("/users", getUsersHandlerAPI, config.All().Web.Domain)
  web.Get("/user", getUserHandlerAPI, config.All().Web.Domain)
  web.Get("/user/delete", deleteUserHandlerAPI, config.All().Web.Domain)
  web.Get("/logout", logoutHandler, config.All().Web.Domain)
  web.Post("/users/new", newUserHandlerAPI, config.All().Web.Domain)
  web.Post("/users/edit", updateUserHandlerAPI, config.All().Web.Domain)
  web.Get("/user/permission/add", addPermissionUserHandlerAPI, config.All().Web.Domain)
  web.Get("/user/permission/delete", deletePermissionUserHandlerAPI, config.All().Web.Domain)
  web.Get("/user/updatepass", resetPasswordHandlerAPI, config.All().Web.Domain)
}

func registerSummaryHandlers(){ //main page - dashboard at '/'
  web.Get("/", dashboardMainPage, config.All().Web.Domain)
}

func registerEntityHandlers(){
  web.Get("/entities", getAllEntitiesHandlerAPI, config.All().Web.Domain)
  web.Post("/entities/new", newEntityHandlerAPI, config.All().Web.Domain)
  web.Post("/entities/edit", updateEntityHandlerAPI, config.All().Web.Domain)
  web.Get("/entity", getEntityHandlerAPI, config.All().Web.Domain)
  web.Get("/entity/events/count", getNumEntityEventsQueued, config.All().Web.Domain)
  web.Get("/e/status", apiHandler(updateEntityStatusHandlerAPI), config.All().Web.Domain)
  web.Post("/e/status", apiHandler(updateEntityStatusHandlerAPI), config.All().Web.Domain)
  web.Get("/e/loc", apiHandler(updateEntityLocationHandlerAPI), config.All().Web.Domain)
  web.Post("/e/loc", apiHandler(updateEntityLocationHandlerAPI), config.All().Web.Domain)
  web.Get("/entityLocations", apiHandler(getEntityLocationsHandlerAPI), config.All().Web.Domain)
  web.Get("/e/events", apiHandler(getEntityQueueAPI), config.All().Web.Domain)
  web.Get("/e/newevent", apiHandler(insertEntityEventAPI), config.All().Web.Domain)
  web.Get("/e/l_events", apiHandler(getEntityQueueAPI_longpoll), config.All().Web.Domain)
}

func registerPluginHandlers(){
  web.Get("/plugins", apiHandler(getAllPluginsHandlerAPI), config.All().Web.Domain)
  web.Post("/plugins/new", apiHandler(newPluginHandlerAPI), config.All().Web.Domain)
  web.Post("/plugins/edit", apiHandler(editPluginHandlerAPI), config.All().Web.Domain)
  web.Get("/plugins/changestate", apiHandler(changePluginStateAPI), config.All().Web.Domain)
  web.Get("/plugin", apiHandler(getPluginHandlerAPI), config.All().Web.Domain)
  web.Post("/plugins/newresource", apiHandler(newResourceHandlerAPI), config.All().Web.Domain)
  web.Post("/plugins/saveresource", apiHandler(editResourceHandlerAPI), config.All().Web.Domain)
  web.Get("/resource", apiHandler(getResourceHandlerAPI), config.All().Web.Domain)
  web.Get("/plugins/deleteresource", apiHandler(deleteResourceHandlerAPI), config.All().Web.Domain)
  web.Get("/plugins/deleteplugin", apiHandler(deletePluginHandlerAPI), config.All().Web.Domain)
}

func registerWebSockets() {
  web.Get("/ws/echotest", websocket.Handler(ws_EchoServer), config.All().Web.Domain)
  web.Get("/ws/logging", websocket.Handler(ws_LogServer), config.All().Web.Domain)
  web.Get("/ws/p/(.*)", pluginsockets.Handle, config.All().Web.Domain)
  web.Get("/ws/entityUpdates", websocket.Handler(ws_EntityUpdateServer), config.All().Web.Domain)
  web.Get("/ws/devclient", websocket.Handler(devclienthandler.Main_ws), config.All().Web.Domain)
  web.Get("/ws/sql", websocket.Handler(ws_SqlQueryServer), config.All().Web.Domain)
}


func registerTemplateViews() {
  web.Get("/view/users", usersAdminMainPage_view, config.All().Web.Domain)
  web.Get("/view/entities", entityAdminViewerPage_view, config.All().Web.Domain)
  web.Get("/view/entity", entityViewerPage_view, config.All().Web.Domain)
  web.Get("/view/entities/form", entityAdminForm_view, config.All().Web.Domain)
  web.Get("/view/entitymap", entityMapPage_view, config.All().Web.Domain)
  web.Get("/view/dashboard/summary", dashboardSummary_view, config.All().Web.Domain)
  web.Get("/view/plugins", pluginAdminListPage_view, config.All().Web.Domain)
  web.Get("/view/plugins/newform", pluginAdminNewPage_view, config.All().Web.Domain)
  web.Get("/view/plugins/editform", pluginAdminEditPage_view, config.All().Web.Domain)
  web.Get("/view/plugins/resourceform", pluginAdminResourcePage_view, config.All().Web.Domain)
  web.Get("/view/data", dataMainPage_view, config.All().Web.Domain)
}


func registerCoreTemplates(){
  logError(registerTemplate("bannertop.tpl", "bannertop"), "Template load error: ")
  logError(registerTemplate("headcontent.tpl", "headcontent"), "Template load error: ")
  logError(registerTemplate("tailcontent.tpl", "tailcontent"), "Template load error: ")
  logError(registerTemplate("apiref.tpl", "apiref"), "Template load error: ")
  logError(registerTemplate("data.tpl", "data"), "Template load error: ")
}

func registerUserTemplates(){
  logError(registerTemplate("login.tpl", "login"), "Template load error: ")
  logError(registerTemplate("user/userpage.tpl", "userpage"), "Template load error: ")
  logError(registerTemplate("user/usercreateeditpage.tpl", "usercreateeditpage"), "Template load error: ")
  logError(registerTemplate("user/userpermissions.tpl", "userpermissions"), "Template load error: ")
}

func registerSummaryTemplates(){
  logError(registerTemplate("dashboardindex.tpl", "dashboardindex"), "Template load error: ")
  logError(registerTemplate("dashboardsummary.tpl", "dashboardsummary"), "Template load error: ")
}

func registerEntityTemplates(){
  logError(registerTemplate("entity/adminentityviewer.tpl", "adminentityviewer"), "Template load error: ")
  logError(registerTemplate("entity/adminentitycreateedit.tpl", "adminentityform"), "Template load error: ")
  logError(registerTemplate("entity/entityviewer.tpl", "entityviewer"), "Template load error: ")
  logError(registerTemplate("entity/entitymap.tpl", "entitymap"), "Template load error: ")
}

func registerPluginTemplates(){
  logError(registerTemplate("plugin/pluginlist.tpl", "pluginlist"), "Template load error: ")
  logError(registerTemplate("plugin/newplugin.tpl", "newplugin"), "Template load error: ")
  logError(registerTemplate("plugin/pluginedit.tpl", "pluginedit"), "Template load error: ")
  logError(registerTemplate("plugin/resourcecreateedit.tpl", "resourcecreateedit"), "Template load error: ")
}

func logError(e error, prefix string){
  if e != nil{
    logging.Error("web", prefix, e.Error())
  }
}
