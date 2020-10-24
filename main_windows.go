//go:generate goversioninfo -64 -icon=Windows/icon_512x512.ico -manifest=Windows/windows.manifest
package main

import (
	"github.com/getlantern/systray"
	"os"
	"encoding/base64"
)

func main() {
	go Run()
	systray.Run(onReady, onExit)
}

func onReady() {
	icon := "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AAEAAAAAAAAAAQAAAAAAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAAAhSlL//8pSlKAAAAAAIAACQAAEAAAQAAAgAAAAAAAAAAAAAAAAAAAAQAAAAAAIAAAgAAgAAAAAAEBAAAAAAAAAIAAAAAAAAAAAgAAAAAAAAGAAAAAAAAAAIAAIAAAAAABAQAAAABAAACAAAAAAAAAAAAAAAAAQAABgAAAAAAAAACAAAAAAAAAAQAAEAAAAAAAgIAAAAAAAAAAABAAAEAAAIAAAAAAAAAAgAAAAAAAAAEAAAAAAAAAAIAAAAAAAAAAAAAAAABAAACAIAAAAAEAAEAQAAAAEAQBAAAAAAAAAgCAAAAAAAABAAAAAAAAAAAAgAAAEAAAAABAAAAAQAAAAQAAABCAAAAAgAAAAEAAAAEAAAAAIAAAAIAAAAAAAABAAAAAAEAAAAEAAAAQACgAAIAAAQgCgAABAAAAACAIAACAAAAAAAAAAAAAAAACAAABAAAAAAEAAgCAAAAAAIAAAQAAEAAAAAAAgAAAAAAAAAAAAAAAAEIAAYAAAAAAAAAAgAAgAAAAAAEAAAAAAAAAAIAAIAAAAAAAAAAAAAAAgAGAAAAAACAAAIAAIAAAAIABAAAAAAAAAACAAAAAAAAAAAAAEAAAQgAAgAAIAAAAAAAAAAAAAAAAAQAAAAABAAAAgAABAAAAAAAAAACACAAAAIQAAFCEAAAA1rWtb9a1rWs="
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(icon)))
	base64.StdEncoding.Decode(base64Text, []byte(icon))
	systray.SetIcon(base64Text)
	systray.SetTitle("Sexy DKP Sync")
	systray.SetTooltip("Sexy DKP Sync")
	update := systray.AddMenuItem("Обнови сейчас", "Обновляет данные прямо сейчас")
	quit := systray.AddMenuItem("Завершить", "Полностью отключает приложение")
	go func() {
		select {
		case <-quit.ClickedCh:
			systray.Quit()
			return
		case <-update.ClickedCh:
			if err := rewrite(); err != nil {
				panic(err)
			}
		}
	}()
}

func onExit() {
	os.Exit(0)
}