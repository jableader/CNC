package web

import (
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/data"
  "github.com/hoisie/web"
)

func loginHandler(ctx *web.Context) {

  usrname := ctx.Params["user"]
  passwd := ctx.Params["pass"]

  isValidLogin, usr := user.CheckAuthByPassword(usrname, passwd, data.DB)

  if isValidLogin {
    logging.Info("web", "User '", usrname, "' has authenticated.")
    skey := session.CreateSession(int(usr.ID), "web", data.DB)
    ctx.SetCookie(web.NewCookie(COOKIE_KEY_NAME, skey, 60*60*24*20))
    ctx.ResponseWriter.Write([]byte("GOOD"))
  }else{
    ctx.ResponseWriter.Write([]byte("ERROR"))
  }
}


func logoutHandler(ctx *web.Context) {
  isLoggedIn, user, s := getSessionByCookie(ctx)
  if isLoggedIn {
    logging.Info("web", "Now logging out:", user.Username)
    session.Delete(s, data.DB)
    deleteSessionKey(ctx)
  } else {
    logging.Warning("web", "/logout called with an invalid session!")
  }
  ctx.Redirect(302, "/")
}
