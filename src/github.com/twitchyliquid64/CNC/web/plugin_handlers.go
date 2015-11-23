package web

import (
  pluginController "github.com/twitchyliquid64/CNC/plugin"
  pluginExec "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/logging"
  pluginData "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/hoisie/web"
  "encoding/json"
)

// Passes back a JSON array of all plugins
//
//
func getAllPluginsHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "getAllPlugins() called unauthorized, aborting")
    return
  }

  plugins := pluginController.GetAll()
  if plugins == nil {
    plugins = []*pluginExec.Plugin{}
  }

  d, err := json.Marshal(plugins)
  if err != nil {
    logging.Error("web-plugin", err)
  }
  ctx.ResponseWriter.Write(d)
}



// API endpoint called to create a new plugin.
// Checks if the session's user is an admin.
//
func newPluginHandlerAPI(ctx *web.Context) {
  isLoggedIn, u, _ := getSessionByCookie(ctx)

  if (!isLoggedIn) || (!u.IsAdmin()){
    logging.Warning("web-plugin", "newPlugin() called unauthorized, aborting")
    return
  }

  decoder := json.NewDecoder(ctx.Request.Body)
  var pl pluginData.Plugin
  err := decoder.Decode(&pl)
  if err != nil {
      logging.Error("web-plugin", "newPluginHandlerAPI() failed to decode JSON:", err)
      ctx.Abort(500, "JSON error")
      return
  }

  err = pluginData.Create(pl, data.DB)
  if err == nil {
      ctx.ResponseWriter.Write([]byte("GOOD"))
  } else {
      ctx.ResponseWriter.Write([]byte(err.Error()))
      logging.Error("web-plugin", err)
  }
}
