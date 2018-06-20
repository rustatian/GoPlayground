package main

import (
	tm "github.com/buger/goterm"
	ui "github.com/gizak/termui"
	"runtime"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	g := ui.NewGauge()
	g.Percent = 50
	g.Width = 50
	g.BorderLabel = "Gauge"

	ui.Render(g)

	ui.Loop()

	//for {
	//	// By moving cursor to top-left position we ensure that console output
	//	// will be overwritten each time, instead of adding new.
	//	tm.MoveCursor(1, 1)
	//
	//	tm.Println("Current Time:", time.Now().Format(time.RFC1123))
	//
	//	tm.Flush() // Call it every time at the end of rendering
	//
	//	time.Sleep(time.Second)
	//}

	chart := tm.NewLineChart(100, 20)

	data := new(tm.DataTable)
	data.AddColumn("Time")
	data.AddColumn("Sin(x)")
	data.AddColumn("Cos(x+1)")

	//for i := 0.1; i < 10; i += 0.1 {
	//	data.AddRow(i, math.Sin(i), math.Cos(i+1))
	//}

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	a := make(map[int]string)
	for i := 0; i < 100000; i++ {

		//tm.MoveCursor(1, 1)

		data.AddRow(float64(i), float64(mem.HeapInuse))

		a[i] = "Lorem input"
		//tm.Printf("Memory alloc: %d | Total alloc: %d | Heap alloc: %d | HeapSys: %d, Number of GCs: %d", mem.HeapInuse, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys, mem.NumGC)
		//tm.Flush()

		runtime.ReadMemStats(&mem)
	}

	tm.Println(chart.Draw(data))
	tm.Flush()

}
