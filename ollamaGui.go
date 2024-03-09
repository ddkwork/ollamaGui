package main

import (
	"cogentcore.org/core/coredom"
	"cogentcore.org/core/events"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/grr"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/units"
	"embed"
	_ "embed"
	"github.com/ddkwork/golibrary/stream/cmd"
	"github.com/ddkwork/golibrary/widget"
	"io/fs"
	"time"
)

//go:embed tokenMock.md
var content string

//go:generate core generate
//go:generate core build -v -t android/arm64
//go:generate core build -v -t windows/amd64
//go:generate go build .
//go:generate go run  -race  .
//go:generate go install .
//go:generate svg embed-image svg/ollama.png

var tokens = []string{"**", "Generic", " type", " constraints", "**", " allow", " you", " to", " specify", " constraints", " on", " a", " type", " that", " can", " vary", " depending", " on", " the", " specific", " type", " being", " instantiated", ".", "\n\n", "**", "Syntax", ":**", "\n\n", "```", "go", "\n", "type", " Name", "[", "T", " any", "]", " string", "\n", "```", "\n\n", "**", "Parameters", ":**", "\n\n", "*", " `", "T", "`:", " The", " type", " variable", ".", " It", " can", " be", " any", " type", ",", " including", " primitive", " types", ",", " structures", ",", " and", " functions", ".", "\n\n", "**", "Examples", ":**", "\n\n", "*", " ", "Integer", " constraint", ":**", "\n", "```", "go", "\n", "type", " Age", "[", "T", " int", "]", " int", "\n", "```", "\n\n", "This", " constraint", " ensures", " that", " `", "T", "`", " is", " an", " integer", " type", ".", "\n\n", "*", " ", "String", " constraint", ":**", "\n", "```", "go", "\n", "type", " Name", "[", "T", " string", "]", "\n", "```", "\n\n", "This", " constraint", " ensures", " that", " `", "T", "`", " is", " a", " string", " type", ".", "\n\n", "*", " ", "Struct", " constraint", ":**", "\n", "```", "go", "\n", "type", " User", "[", "T", " struct", "]", " {", "\n", "  ", "Name", " string", "\n", "  ", "Age", "  ", "int", "\n", "}", "\n", "```", "\n\n", "This", " constraint", " ensures", " that", " `", "T", "`", " is", " a", " struct", " type", " with", " at", " least", " two", " fields", " named", " `", "Name", "`", " and", " `", "Age", "`.", "\n\n", "*", " ", "Function", " constraint", ":**", "\n", "```", "go", "\n", "type", " Calculator", "[", "T", " any", "]", " func", "(", "T", ",", " T", ")", " T", "\n", "```", "\n\n", "This", " constraint", " ensures", " that", " `", "T", "`", " is", " a", " type", " that", " implements", " the", " `", "Calculator", "`", " interface", ".", "\n\n", "**", "Benefits", " of", " using", " generic", " type", " constraints", ":**", "\n\n", "*", " ", "Code", " reus", "ability", ":**", " You", " can", " apply", " the", " same", " constraint", " to", " multiple", " types", ",", " reducing", " code", " duplication", ".", "\n", "*", " ", "Type", " safety", ":**", " Constraints", " ensure", " that", " only", " valid", " types", " are", " used", ",", " preventing", " runtime", " errors", ".", "\n", "*", " ", "Improved", " maintain", "ability", ":**", " By", " separating", " the", " constraint", " from", " the", " type", ",", " it", " becomes", " easier", " to", " understand", " and", " modify", ".", "\n\n", "**", "Note", ":**", "\n\n", "*", " Generic", " type", " constraints", " are", " not", " applicable", " to", " primitive", " types", " (", "e", ".", "g", ".,", " `", "int", "`,", " `", "string", "`).", "\n", "*", " Constraints", " can", " be", " applied", " to", " function", " types", " only", " if", " the", " function", " is", " generic", ".", "\n", "*", " Constraints", " can", " be", " used", " with", " type", " parameters", ",", " allowing", " you", " to", " specify", " different", " constraints", " for", " different", " types", "."}

//go:embed svg/ollama.svg
var windowIco []byte

//go:embed svg/*.svg
var myIcons embed.FS

func main() {
	icons.AddFS(grr.Log1(fs.Sub(myIcons, "svg")))
	gi.TheApp.SetIconBytes(windowIco)
	b := gi.NewBody("ollamaGui")
	b.AddAppBar(func(toolbar *gi.Toolbar) {
		gi.NewButton(toolbar).SetText("install") //todo set icon
		gi.NewButton(toolbar).SetText("start server").OnClick(func(e events.Event) { cmd.Run("ollama serve") })
		gi.NewButton(toolbar).SetText("stop server").OnClick(func(e events.Event) {
			//todo kill thread ?
			//netstat -aon|findstr 11434
		})
		gi.NewButton(toolbar).SetText("logView")
		gi.NewButton(toolbar).SetText("about").SetIcon("about")
	})

	splits := gi.NewSplits(b)

	leftFrame := gi.NewFrame(splits)
	leftFrame.Style(func(s *styles.Style) { s.Direction = styles.Column })
	widget.NewSeparatorWithLabel("module choose", leftFrame)
	giv.NewFileView(leftFrame)
	gi.NewButton(leftFrame).SetText("run module").Style(func(s *styles.Style) {
		s.Align.Self = styles.End
		s.Min.Set(units.Dp(33))
	})

	rightSplits := gi.NewSplits(splits)
	splits.SetSplits(.2, .8)

	frame := gi.NewFrame(rightSplits)
	frame.Style(func(s *styles.Style) { s.Direction = styles.Column })
	widget.NewSeparatorWithLabel("chat with ai", frame)

	answer := gi.NewFrame(frame)
	go func() {
		for _, token := range tokens {
			println(token)
			b.AsyncLock()

			//todo
			//We need to check the token's newlines to deal with it,
			//and secondly, we want to keep the previous token instead of deleting it
			//answer.DeleteChildren(false)

			grr.Log(coredom.ReadMDString(coredom.NewContext(), answer, token))
			answer.Update()
			b.AsyncUnlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	downframe := gi.NewFrame(frame)
	downframe.Style(func(s *styles.Style) { s.Direction = styles.Row })
	topic := gi.NewButton(downframe).SetText("new topic").SetIcon(icons.ClearAll)
	topic.Style(func(s *styles.Style) {
		s.Min.Set(units.Dp(33))
	})
	//gi.NewTextField(downframe).SetText("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").Style(func(s *styles.Style) { s.SetTextWrap(true) })
	gi.NewTextField(downframe).SetText("go1.22 Generic type constraints").Style(func(s *styles.Style) {
		//s.Max.X.Em(10) //todo height not working
		s.Max.Zero()
	})
	gi.NewButton(downframe).SetText("send").Style(func(s *styles.Style) {
		s.Min.Set(units.Dp(33))
	})

	rightSplits.SetSplits(.6, .4) //todo not working

	b.RunMainWindow()
}
