package builtin

import (
  "github.com/twitchyliquid64/CNC/plugin/exec"
  "github.com/twitchyliquid64/CNC/registry"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/util"
  "github.com/robertkrimen/otto"
  "github.com/robfig/cron"
)

var cronObj *cron.Cron
const CRON_ID_LENGTH = 12
const CRON_HOOK_PREFIX = "cron_"

func init() {
  cronObj = cron.New()
  cronObj.Start()
}

// Called when JS code executes cron.schedule()
// cronString format defined @: https://godoc.org/github.com/robfig/cron
//
func function_cron_schedule(plugin *exec.Plugin, call otto.FunctionCall)otto.Value{
  cronString := call.Argument(0).String()
  method := util.GetFunc(call.Argument(1), plugin.VM)

  cronID := util.RandAlphaKey(CRON_ID_LENGTH)
  cronObj.AddFunc(cronString, func(){
    registry.DispatchEvent(CRON_HOOK_PREFIX + cronID, nil)
  })

  hook := CronHook{P: plugin, Callback: &method, CronID: cronID}
  plugin.RegisterHook(&hook)
  return otto.Value{}
}

type CronHook struct {
  CronID string
  P *exec.Plugin
  Callback *otto.Value
}

func (h *CronHook)Destroy(){
  logging.Info(h.Name(), "hook.Destroy() called")
}
func (h *CronHook)Name()string{
  return CRON_HOOK_PREFIX + h.CronID
}
func (h *CronHook)Dispatch(data interface{}){
  val, err := h.P.VM.ToValue(struct{CronID string}{CronID: h.Name()})
  if err != nil {
    logging.Error("builtin-cron", err.Error())
  }

  h.P.PendingInvocations <- &exec.JSInvocation{Callback: h.Callback, Parameters: []interface{} { val }}
}
