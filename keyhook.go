package main

import (
	"log"

	hook "github.com/robotn/gohook"
)

func setupKeyboardHook() {
	defer hook.End()

	hook.Register(hook.KeyDown, []string{""}, func(e hook.Event) {
		increment := func() int {
			if (e.Mask & 2) != 0 {
				return -1
			} else {
				return 1
			}
		}()

		switch e.Rawcode {
		case hook.KeychartoRawcode(cfg.Keymaps.IncSupWins):
			log.Printf("Add %d win to sup role\n", increment)
			addWin("sup", increment)
		case hook.KeychartoRawcode(cfg.Keymaps.IncSupTies):
			log.Printf("Add %d tie to sup role\n", increment)
			addTie("sup", increment)
		case hook.KeychartoRawcode(cfg.Keymaps.IncSupLosses):
			log.Printf("Add %d lose to sup role\n", increment)
			addLose("sup", increment)
		case hook.KeychartoRawcode(cfg.Keymaps.IncDpsWins):
			log.Printf("Add %d win to dps role\n", increment)
			addWin("dps", increment)
		case hook.KeychartoRawcode(cfg.Keymaps.IncDpsTies):
			log.Printf("Add %d tie to dps role\n", increment)
			addTie("dps", increment)
		case hook.KeychartoRawcode(cfg.Keymaps.IncDpsLosses):
			log.Printf("Add %d lose to dps role\n", increment)
			addLose("dps", increment)
		case hook.KeychartoRawcode(cfg.Keymaps.IncTankWins):
			log.Printf("Add %d win to tank role\n", increment)
			addWin("tank", increment)
		case hook.KeychartoRawcode(cfg.Keymaps.IncTankTies):
			log.Printf("Add %d tie to tank role\n", increment)
			addTie("tank", increment)
		case hook.KeychartoRawcode(cfg.Keymaps.IncTankLosses):
			log.Printf("Add %d lose to tank role\n", increment)
			addLose("tank", increment)
		case hook.KeychartoRawcode(cfg.Keymaps.ResetStats):
			log.Println("Reset stats")
			resetStats()
		}
	})

	log.Println("Setup global keyboard hook")
	s := hook.Start()
	<-hook.Process(s)
}
