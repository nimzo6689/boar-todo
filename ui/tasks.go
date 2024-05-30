package ui

import (
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/nimzo6689/boar-todo/config"
	"github.com/nimzo6689/boar-todo/storage/models"
	"github.com/nimzo6689/boar-todo/ui/widget"
)

type TaskTable struct {
	table        *widget.Table
	items        []*models.Task
	hasFocus     bool
	metadataFunc func(Task *models.Task)
	deleteFunc   func(Task *models.Task)
	sortFunc     func(column string, sort widget.Sort)
}

func (t *TaskTable) Draw(screen tcell.Screen) {
	t.table.Draw(screen)
}

func (t *TaskTable) GetRect() (int, int, int, int) {
	return t.table.GetRect()
}

func (t *TaskTable) SetRect(x, y, width, height int) {
	t.table.SetRect(x, y, width, height)
}

func (t *TaskTable) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		key := event.Key()
		if key == tcell.KeyCtrlSpace {
			selected, _ := t.table.GetSelection()
			Task := t.items[selected]
			t.metadataFunc(Task)
		} else if key == tcell.KeyDelete {
			if t.deleteFunc != nil {
				index, _ := t.table.GetSelection()
				Task := t.items[index-1]
				t.deleteFunc(Task)
			}
		} else if event.Rune() == 'n' {
			t.moveCursor(10)
		} else if event.Rune() == 'm' {
			t.moveCursor(-10)
		} else {
			t.table.InputHandler()(event, setFocus)
		}
	}
}

func (t *TaskTable) moveCursor(n int) {
	index, col := t.table.GetSelection()
	result := index + n
	if result >= len(t.items) {
		result = len(t.items) - 1
	} else if result <= 0 {
		result = 1
	}
	t.table.Select(result, col)
}

func (t *TaskTable) Focus(delegate func(p tview.Primitive)) {
	t.hasFocus = true
	t.table.SetBorderColor(config.Configuration.Colors.BorderFocus)
	t.table.Focus(delegate)
}

func (t *TaskTable) HasFocus() bool {
	return t.table.HasFocus()
}

func (t *TaskTable) Blur() {
	t.hasFocus = false
	t.table.SetBorderColor(config.Configuration.Colors.Border)
	t.table.Blur()
}

func (t *TaskTable) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return t.table.MouseHandler()
}

func (t *TaskTable) PasteHandler() func(text string, setFocus func(p tview.Primitive)) {
	return t.table.PasteHandler()
}

func (t *TaskTable) SetData(data []*models.Task) {
	if data == nil {
		return
	}
	t.items = data

	t.table.Clear(false)
	for i, v := range data {
		row := []string{
			v.Subject,
			v.Description,
			ShortTimeSince(v.CreatedAt),
		}

		t.table.AddRow(i, row...)
	}
	if len(t.items) > 0 {
		t.table.Select(1, 0)
	}
}

func (t *TaskTable) ResetCursor() {
	t.table.Select(1, 0)

}

func (t *TaskTable) SetDeleteFunc(delete func(Task *models.Task)) {
	t.deleteFunc = delete
}

func (t *TaskTable) SetSortFunc(sort func(column string, sort widget.Sort)) {
	t.sortFunc = sort
}

func NewTaskTable() *TaskTable {
	t := &TaskTable{
		table: widget.NewTable(),
		items: []*models.Task{},
	}

	colors := config.Configuration.Colors.Tasks
	t.table.SetBackgroundColor(colors.Background)
	t.table.SetBorder(true)
	t.table.SetBorders(false)
	t.table.SetBorderColor(config.Configuration.Colors.Border)
	t.table.SetSelectedStyle(tcell.StyleDefault.Foreground(colors.TextSelected).Background(colors.BackgroundSelected))
	t.table.SetSelectable(true, false)
	t.table.SetFixed(1, 10)

	t.table.SetAddCellFunc(t.addCell)
	t.table.SetShowIndex(true)
	t.table.SetColumns([]string{"Subject", "Description"})
	t.table.SetColumnWidths([]int{3, 25, 35, 20, 10, 15, 10})
	t.table.SetColumnExpansions([]int{0, 1, 3, 1, 1, 1, 1})
	t.table.SetSort(0, widget.SortAsc)
	t.table.SetSortFunc(t.sort)
	return t
}

func (t *TaskTable) GetSelection() *models.Task {
	index, _ := t.table.GetSelection()
	if t.items == nil {
		return nil
	}
	if index > len(t.items) {
		return nil
	}
	return t.items[index-1]
}

func (t *TaskTable) addCell(cell *tview.TableCell, header bool, row int) {
	if header {
		cell.SetTextColor(config.Configuration.Colors.Tasks.HeaderText)
		cell.SetAlign(tview.AlignLeft)
	} else {
		cell.SetTextColor(config.Configuration.Colors.Tasks.Text)
		if row%2 == 1 {
			cell.SetBackgroundColor(config.Configuration.Colors.Tasks.Background2nd)
		}
		cell.SetAlign(tview.AlignLeft)
	}
}

func (t *TaskTable) sort(column string, sort widget.Sort) {
	if t.sortFunc != nil {
		t.sortFunc(column, sort)
	}
}
