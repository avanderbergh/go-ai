package cmd

import (
	"go-ai/internal"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func Run() {

	// Register the full hotkey combination directly with gohook
	hook.Register(hook.KeyDown, []string{"space", "cmd", "shift"}, func(ev hook.Event) {

		if ev.Rawcode != 49 {
			return
		}

		robotgo.KeyTap("c", "cmd")
		time.Sleep(time.Millisecond * 200) // Allow clipboard to settle
		text, err := robotgo.ReadAll()

		if err != nil {
			log.Printf("Error reading from clipboard: %v", err)
			return
		}

		if text == "" {
			log.Println("Clipboard is empty, nothing to translate.")
			return
		}

		log.Printf("Text to translate: [%s]", text)

		translated := internal.Translate(text)
		log.Printf("Translated text: [%s]", translated)

		if translated == "" {
			log.Println("Translation resulted in empty string, not pasting.")
			return
		}

		robotgo.WriteAll(translated)
		time.Sleep(time.Millisecond * 200) // Allow clipboard to settle
		robotgo.KeyTap("v", "cmd")
		log.Println("Translation pasted.")
	})

	log.Println("Hotkey listener started. Press CMD+SHIFT+SPACE to trigger.")
	events := hook.Start()
	<-hook.Process(events)
}
