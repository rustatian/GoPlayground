package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	NewCalculatorForm().Show()

	widgets.QApplication_Exec()
}

func NewCalculatorForm() *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2("/Users/0xdev/Projects/repo/GoPlayground/ui/1/qml/calculatorform.ui")

	file.Open(core.QIODevice__ReadOnly)
	var formWidget = loader.Load(file, widget)
	file.Close()

	var (
		uiInputspinbox1 = widgets.NewQSpinBoxFromPointer(widget.FindChild("inputSpinBox1", core.Qt__FindChildrenRecursively).Pointer())
		uiInputspinbox2 = widgets.NewQSpinBoxFromPointer(widget.FindChild("inputSpinBox2", core.Qt__FindChildrenRecursively).Pointer())
		uiOutputwidget  = widgets.NewQLabelFromPointer(widget.FindChild("outputWidget", core.Qt__FindChildrenRecursively).Pointer())
	)

	uiInputspinbox1.ConnectValueChanged(func(value int) {
		uiOutputwidget.SetText(fmt.Sprint(value + uiInputspinbox2.Value()))
	})

	uiInputspinbox2.ConnectValueChanged(func(value int) {
		uiOutputwidget.SetText(fmt.Sprint(value + uiInputspinbox1.Value()))
	})

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(formWidget, 0, 0)
	widget.SetLayout(layout)

	widget.SetWindowTitle("Calculator Builder")

	return widget
}
