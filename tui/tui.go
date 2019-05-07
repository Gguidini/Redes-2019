package tui

import (
	// "fmt"
	"log"
	// "time"

	"github.com/marcusolsson/tui-go"
	"github.com/Redes-2019/userinterface"
	"github.com/Redes-2019/connection"
)

type post struct {
	username string
	message  string
	time     string
}

var posts = []post{
	{username: "john", message: "hi, what's up?", time: "14:41"},
	{username: "jane", message: "not much", time: "14:43"},
}

// func TuiFunc() {
// 	sidebar := tui.NewVBox(
// 		tui.NewLabel("CHANNELS"),
// 		tui.NewLabel("general"),
// 		tui.NewLabel("random"),
// 		tui.NewLabel(""),
// 		tui.NewLabel("DIRECT MESSAGES"),
// 		tui.NewLabel("slackbot"),
// 		tui.NewSpacer(),
// 	)
// 	sidebar.SetBorder(true)

// 	history := tui.NewVBox()

// 	for _, m := range posts {
// 		history.Append(tui.NewHBox(
// 			// tui.NewLabel(m.time),
// 			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", m.username))),
// 			tui.NewLabel(m.message),
// 			tui.NewSpacer(),
// 		))
// 	}

// 	historyScroll := tui.NewScrollArea(history)
// 	historyScroll.SetAutoscrollToBottom(true)

// 	historyBox := tui.NewVBox(historyScroll)
// 	historyBox.SetBorder(true)

// 	input := tui.NewEntry()
// 	input.SetFocused(true)
// 	input.SetSizePolicy(tui.Expanding, tui.Maximum)

// 	inputBox := tui.NewHBox(input)
// 	inputBox.SetBorder(true)
// 	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

// 	chat := tui.NewVBox(historyBox, inputBox)
// 	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

// 	input.OnSubmit(func(e *tui.Entry) {
// 		history.Append(tui.NewHBox(
// 			tui.NewLabel(time.Now().Format("15:04")),
// 			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", "john"))),
// 			tui.NewLabel(e.Text()),
// 			tui.NewSpacer(),
// 		))
// 		input.SetText("")
// 	})

// 	root := tui.NewHBox(sidebar, chat)

// 	ui, err := tui.New(root)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ui.SetKeybinding("Esc", func() { ui.Quit() })

// 	if err := ui.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// Mostra a tui
func Show(client *connection.IrcClient) {
	
}

func TuiHandler(client *connection.IrcClient) {
	history := tui.NewVBox()

	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	// userinterface.ReadCommand("> ")
	input.OnSubmit(func(e *tui.Entry) {
		out := ""
		command := userinterface.ReadCommand(e.Text())
		if command == nil {
			out = "[Input error] comando invalido"
		} else {
			validStructure, err := userinterface.VerifyStructure(command)
			if validStructure == false {
				// fmt.Println(err)
				out = err
			} else {
				client.DataFromUser <- command
			}
		}

		history.Append(tui.NewHBox(
			tui.NewLabel(out),
			tui.NewSpacer(),
		))
		input.SetText("")
	})

	root := tui.NewHBox(chat)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}