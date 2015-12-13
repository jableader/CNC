package builtin

import (
  "github.com/twitchyliquid64/CNC/web/pluginsockets"
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/util"
  "github.com/robertkrimen/otto"
)


const WEBSOCK_PREFIX = "ws_event_"
const WSHANDLER_ID_LENGTH = 12


// Called when JS code executes websocket.handle()
//
//
func function_websockets_register(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  patternRegex  := call.Argument(0).String()
  onOpenMethod  := util.GetFunc(call.Argument(1), plugin.VM)

  hookID := util.RandAlphaKey(WSHANDLER_ID_LENGTH)
  hook := WebsocketHook{P: plugin,
                        onOpen: &onOpenMethod,
                        HookID: hookID,
                        Pattern: patternRegex}
  plugin.RegisterHook(&hook)

  if pluginsockets.AddHook(hook.Name(), hook.Name(), hook.Name(), patternRegex) { //use the same hook dispatch for all three event types
    return otto.TrueValue()
  } else {
    return otto.FalseValue()
  }
}

type WebsocketHook struct {
  Pattern string
  HookID string
  P *exec.Plugin

  onOpen *otto.Value
  webObject *otto.Object
}

type WebSock interface{
  Write(string)
  URL()string
  Parameter(string)string

  LoggedIn()bool
  User()*user.User
  Session()*session.Session
  GetID()uint64
  Addr()string
  Close()
}

type SocketEvent interface {
  Event() string
  GetData() string
  Sock() interface{}
}





func (h *WebsocketHook)Destroy(){
  logging.Info(h.Name(), "hook.Destroy() called")
  pluginsockets.RemoveHook(h.Name())
}

func (h *WebsocketHook)Name()string{
  return WEBSOCK_PREFIX + h.HookID
}


func (h *WebsocketHook)Dispatch(data interface{}){
  event := data.(SocketEvent)
  sock := event.Sock().(WebSock)

  if h.webObject == nil {
    h.webObject = newSocketObject(h.P, sock)
  }

  if event.Event() == "OPEN" {
    logging.Info(h.Name(), "Dispatch() OPEN")
    h.P.PendingInvocations <- &exec.JSInvocation{Callback: h.onOpen, Parameters: []interface{} { h.webObject }}
  } else if event.Event() == "MSG" {
    logging.Info(h.Name(), "Dispatch() MSG")
    h.queueMethodInvocation("onmessage", event.GetData())
  } else if event.Event() == "CLOSE" {
    logging.Info(h.Name(), "Dispatch() CLOSE")
    h.queueMethodInvocation("onclose")
  }
}

func (h *WebsocketHook) queueMethodInvocation(functionName string, parameters ...interface{}) {
  method, err := h.webObject.Get(functionName);
  if (err != nil) {
    logging.Error("builtin-ws", err.Error())
  }

  if (!method.IsUndefined()) {
    if (!method.IsFunction()) {
      panic("Expected field " + functionName + " to be a function")
    }

    h.P.PendingInvocations <- &exec.JSInvocation{Callback: &method, Parameters: parameters}
  }
}

func newSocketObject(p *exec.Plugin, s WebSock) *otto.Object {
  obj, err := p.VM.Object("new Object()")
  if err != nil {
    logging.Error("builtin-ws", err.Error())
  }

  obj.Set("write", func(in otto.FunctionCall)otto.Value{
    s.Write(in.Argument(0).String())
    return otto.UndefinedValue()
  })

  obj.Set("url", s.URL())
  obj.Set("id", s.GetID())
  obj.Set("addr", s.Addr())
  obj.Set("close", func(in otto.FunctionCall)otto.Value{
    s.Close()
    return otto.UndefinedValue()
  })
  obj.Set("parameter", func(in otto.FunctionCall)otto.Value{
    ret, _ := otto.ToValue(s.Parameter(in.Argument(0).String()))
    return ret
  })
  obj.Set("isLoggedIn", func(in otto.FunctionCall)otto.Value{
    ret, _ := otto.ToValue(s.LoggedIn())
    return ret;
  })
  obj.Set("user", func(in otto.FunctionCall)otto.Value{
    ret, err := p.VM.ToValue(s.User())
    if err != nil {
      logging.Error("builtin-ws", err.Error())
    }
    return ret
  })
  obj.Set("session", func(in otto.FunctionCall)otto.Value{
    ret, _ := p.VM.ToValue(s.Session())
    return ret
  })
  obj.Set("onmessage", func(in otto.FunctionCall)otto.Value {
    logging.Warning(p.Name, "Message was recieved for socket but no handler exists")
    return otto.UndefinedValue()
  })

  return obj
}
