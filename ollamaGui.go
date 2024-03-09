package main

import (
	"cogentcore.org/core/coredom"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/grr"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/units"
	_ "embed"
	"github.com/ddkwork/golibrary/widget"
)

//go:embed example.md
var content string

func main() {
	b := gi.NewBody("ollamaGui")
	b.AddAppBar(func(toolbar *gi.Toolbar) {
		gi.NewButton(toolbar).SetText("install") //todo set icon
		gi.NewButton(toolbar).SetText("start server")
		gi.NewButton(toolbar).SetText("stop server")
		gi.NewButton(toolbar).SetText("logView")
		gi.NewButton(toolbar).SetText("about")
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
	grr.Log(coredom.ReadMDString(coredom.NewContext(), frame, content)) //todo

	//answer := gi.NewFrame(frame)
	//go func() {
	//	file := stream.NewReadFile("D:\\workspace\\workspace\\branch\\gpt4\\ollama\\ollama\\log.log.md")
	//	lines, ok := file.ToLines()
	//	if !ok {
	//		return
	//	}
	//	for _, line := range lines {
	//		println(line)
	//		b.AsyncLock()
	//		//answer.DeleteChildren(false)
	//		grr.Log(coredom.ReadMDString(coredom.NewContext(), answer, line))
	//		answer.Update()
	//		b.AsyncUnlock()
	//		time.Sleep(time.Second)
	//	}
	//}()

	downframe := gi.NewFrame(frame)
	downframe.Style(func(s *styles.Style) { s.Direction = styles.Row })
	topic := gi.NewButton(downframe).SetText("new topic").SetIcon(icons.ClearAll)
	topic.Style(func(s *styles.Style) {
		s.Min.Set(units.Dp(33))
	})
	//gi.NewTextField(downframe).SetText("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").Style(func(s *styles.Style) { s.SetTextWrap(true) })
	gi.NewTextField(downframe).SetText("Multiline textfield with a  long initial text").Style(func(s *styles.Style) {
		//s.Max.X.Em(10) //todo height not working
		s.Max.Zero()
	})
	gi.NewButton(downframe).SetText("send").Style(func(s *styles.Style) {
		s.Min.Set(units.Dp(33))
	})

	rightSplits.SetSplits(.6, .4) //todo not working

	b.RunMainWindow()
}
