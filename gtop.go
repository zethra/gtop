package main

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/zethra/termui"
)


func main() {
	//v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	//fmt.Println(v)
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	//termui.UseTheme("helloworld")


	lc0 := termui.NewLineChart()
	lc0.BorderLabel = "Memory Usage"
	lc0.Data = []float64{0, 100}
	lc0.Width = 50
	lc0.Height = 12
	lc0.X = 0
	lc0.Y = 0
	lc0.AxesColor = termui.ColorWhite
	lc0.LineColor = termui.ColorGreen | termui.AttrBold


	termui.Render(lc0)
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Handle("/timer/1s", func(e termui.Event) {
		t := e.Data.(termui.EvtTimer)
		// t is a EvtTimer
		if t.Count%2 ==0 {
			//lc0.Data = append(lc0.Data, float64(t.Count))
			v, _ := mem.VirtualMemory()
			lc0.Data = append(lc0.Data, v.UsedPercent)
			termui.Render(lc0)
		}
	})
	termui.Loop()
}