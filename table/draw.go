package table

import (
	"os"
	"task/config"
	"task/repository"

	"github.com/jedib0t/go-pretty/v6/table"
)

var t table.Writer

var tableStyleMap = map[string]table.Style{
	"default": table.StyleDefault,
	"light":   table.StyleLight,
	"bold":    table.StyleBold,
	"rounded": table.StyleRounded,
	"doubled": table.StyleDouble,
}

func init() {
	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
}

func Init() {
	t.SetStyle(tableStyleMap[config.GetConfig().TableStyle])
	t.Style().Options.SeparateColumns = config.GetConfig().SeparateColumns
	t.Style().Options.SeparateRows = config.GetConfig().SeparateRows
}

func PrintTaskTable(tasks []repository.Task) {
	t.AppendHeader(table.Row{"id", "task", "notes", "done"})
	for _, task := range tasks {
		t.AppendRow(table.Row{task.Id, task.Topic, task.Notes, task.Done})

	}

	t.AppendSeparator()
	t.Render()
}

func PrintCathegoryTable(cathegories []repository.Cathegory) {
	t.AppendHeader(table.Row{"id", "body"})
	for _, cathegory := range cathegories {

		t.AppendRow(table.Row{cathegory.Id, cathegory.Body})

	}
	t.AppendSeparator()
	t.Render()
}
