package main

import (
	"cogentcore.org/core/coredom"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/grr"
	"cogentcore.org/core/styles"
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

	lestSplits := gi.NewSplits(splits)
	lestSplits.Style(func(s *styles.Style) { s.Direction = styles.Column })
	widget.NewSeparatorWithLabel("module choose", lestSplits)
	giv.NewFileView(lestSplits)
	gi.NewButton(lestSplits).SetText("run module")
	lestSplits.SetSplits(.1, .6, .3) //todo not working

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
	//		updt := answer.UpdateStartAsync()
	//		//answer.DeleteChildren(false)
	//		grr.Log(coredom.ReadMDString(coredom.NewContext(), answer, line))
	//		answer.Update()
	//		answer.UpdateEndAsyncLayout(updt)
	//		time.Sleep(time.Second)
	//	}
	//}()

	downframe := gi.NewFrame(frame)
	downframe.Style(func(s *styles.Style) { s.Direction = styles.Row })
	gi.NewButton(downframe).SetText("new topic")
	//gi.NewTextField(downframe).SetText("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").Style(func(s *styles.Style) { s.SetTextWrap(true) })
	gi.NewTextField(downframe).SetText("Multiline textfield with a relativelyrelativelyrelativelyrelativelyrelativelyrelativelyrelatively long initial text").Style(func(s *styles.Style) { s.SetTextWrap(true) })
	gi.NewButton(downframe).SetText("send")

	rightSplits.SetSplits(.6, .4) //todo not working

	b.RunMainWindow()
}
