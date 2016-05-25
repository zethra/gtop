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

	graphs := make(map[string]*termui.FixedLineChart)

	memory := termui.NewFixedLineChart()
	memory.BorderLabel = "Memory Usage"
	memory.Data = []float64{}
	memory.Width = 50
	memory.Height = 12
	memory.X = 0
	memory.Y = 0
	memory.TopValue = 100
	memory.BottomValue = 0
	memory.AxesColor = termui.ColorWhite
	memory.LineColor = termui.ColorGreen | termui.AttrBold
	graphs["memory"] = memory


	termui.Render(memory)
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Handle("/timer/1s", func(e termui.Event) {
		t := e.Data.(termui.EvtTimer)
		// t is a EvtTimer
		if t.Count%2 ==0 {
			update(graphs)
		}
	})
	termui.Loop()
}


func update(graphs map[string]*termui.FixedLineChart) {
	//lc0.Data = append(lc0.Data, float64(t.Count))
	memory := graphs["memory"]
	v, _ := mem.VirtualMemory()
	memory.Data = append(memory.Data, v.UsedPercent)
	termui.Render(memory)
}