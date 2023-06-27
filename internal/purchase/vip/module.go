package vip

import (
	"github.com/phuhao00/greatestworks-proto/module"
	"greatestworks/internal"
	"sync"
)

const (
	ModuleName = "vip"
)

var (
	Mod         *Module
	onceInitMod sync.Once
)

func init() {
	internal.MManager.RegisterModule(ModuleName, GetMod())
}

type Module struct {
	*internal.BaseModule
}

func GetMod() *Module {
	onceInitMod.Do(func() {
		Mod = &Module{
			BaseModule: internal.NewBaseModule(),
		}
	})
	return Mod
}

func (m *Module) SetName(name string) {
	m.BaseModule.SetName(name)
}

func (m *Module) OnDailyReset() {
	//
}

func (m *Module) OnRecharge() {
	//add exp
}

func (m *Module) GetName() string {
	return module.Module_Vip.String()
}
