package widget

import (
	"fmt"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//Modal interface creates a modal that overlaps other views and get's destroyed when it's ready

// Modal interface is primitive that is drawn on top of other views and get's destroyed when it's ready
type Modal interface {
	//Primitive
	tview.Primitive
	//SetDoneFunc sets function that get's called when modal wants to close itself
	SetDoneFunc(doneFunc func())
	//Setvisible tells modal to show or hide itself
	SetVisible(visible bool)
}

type ModalSize int

const (
	//ModalSizeSmall creates modal of 1/5 size of layout, or 2 columns
	ModalSizeSmall ModalSize = 2
	//ModalSizeMedium creates modal of 2/5 size of layout, or 4 columns
	ModalSizeMedium ModalSize = 3
	//ModalSizeMedium creates modal of 3/5 size of layout, or 6 columns
	ModalSizeLarge ModalSize = 4
)

/*
ModalLayout is a grid layout that draws modal on center of grid.
To add ordinary items to layout, get grid with Grid() function. Layout consists of 10 columns and rows,
each one of size -1. This can be changed with SetGridSize, but size must still 10. Modal is drawn on middle 4 cells
Use AddModal and RemoveModal to manage modals. Only single modal can be shown at a time.
*/
type ModalLayout struct {
	grid       *tview.Grid
	hasModal   bool
	customGrid bool
	modal      Modal

	//Default grid col/row weights
	gridAxisX []int
	gridAxisY []int
}

// NewModalLayout creates new modal layout and returns it
func NewModalLayout() *ModalLayout {
	m := &ModalLayout{
		grid:       tview.NewGrid(),
		hasModal:   false,
		customGrid: false,
		modal:      nil,
		gridAxisX:  nil,
		gridAxisY:  nil,
	}

	/*
		Put modal to rows/cols 3-4
		Changing these requires also changing AddColumn()-> grid.AddItem indices.
	*/
	m.gridAxisX = []int{-1, -1, -1, -1, -1}
	m.gridAxisX = append(m.gridAxisX, m.gridAxisX...)
	m.gridAxisY = m.gridAxisX

	m.grid.SetRows(m.gridAxisX...)
	m.grid.SetColumns(m.gridAxisX...)
	m.grid.SetMinSize(2, 2)

	return m
}

func (m *ModalLayout) Draw(screen tcell.Screen) {
	m.grid.Draw(screen)
}

func (m *ModalLayout) GetRect() (int, int, int, int) {
	return m.grid.GetRect()
}

func (m *ModalLayout) SetRect(x, y, width, height int) {
	m.grid.SetRect(x, y, width, height)
}

func (m *ModalLayout) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return m.grid.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {

	})
}

func (m *ModalLayout) Focus(delegate func(p tview.Primitive)) {
	m.grid.Focus(delegate)
}

func (w *ModalLayout) HasFocus() bool {
	return w.grid.HasFocus()
}

func (m *ModalLayout) Blur() {
	m.grid.Blur()
}

func (w *ModalLayout) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return w.grid.MouseHandler()
}

func (w *ModalLayout) PasteHandler() func(text string, setFocus func(p tview.Primitive)) {
	return w.grid.PasteHandler()
}

// GetGrid returns underlying grid that items are added to
func (m *ModalLayout) Grid() *tview.Grid {
	return m.grid
}

// GetGridSize returns grid that's in use
func (m *ModalLayout) GetGridSize() []int {
	return m.gridAxisX
}

func (m *ModalLayout) SetGridXSize(grid []int) error {
	if len(grid) != 10 {
		return fmt.Errorf("invalid size")
	}
	m.gridAxisX = grid

	m.grid.SetColumns(m.gridAxisX...)
	return nil
}

func (m *ModalLayout) SetGridYSize(grid []int) error {
	if len(grid) != 10 {
		return fmt.Errorf("invalid size")
	}
	m.gridAxisY = grid
	m.grid.SetRows(m.gridAxisY...)
	return nil
}

// AddDynamicModal adds modal of dynamic size
func (m *ModalLayout) AddDynamicModal(modal Modal, size ModalSize) {
	m.addModal(modal, 0, 0, false, size)
}

// AddFixedModal adds modal of fixed size.
// Size parameter controls how many rows and columns are used for modal
func (m *ModalLayout) AddFixedModal(modal Modal, height, width uint, size ModalSize) {
	m.addModal(modal, height, width, true, size)
}

// AddModal adds modal to center of layout
// lockSize flag defines whether modal size should be locked or dynamic.
func (m *ModalLayout) addModal(modal Modal, height, width uint, lockSize bool, size ModalSize) {
	r, w := 0, 0
	span := 0
	switch size {
	case ModalSizeSmall:
		r, w = 4, 4
		span = 2
	case ModalSizeMedium:
		r, w = 3, 3
		span = 4
	case ModalSizeLarge:
		r, w = 2, 2
		span = 6
	default:
		return
	}

	if m.hasModal {
		return
	}
	if !lockSize {
		m.customGrid = false
		m.grid.AddItem(modal, r, w, span, span, 8, 30, true)
	} else {
		m.customGrid = true
		x := make([]int, len(m.gridAxisX))
		y := make([]int, len(m.gridAxisY))
		copy(x, m.gridAxisX)
		copy(y, m.gridAxisY)

		// single col / row size
		row := width / uint(span)
		col := height / uint(span)

		for i := r; i < r+span; i++ {
			x[i] = int(row)
			y[i] = int(col)
		}
		m.grid.SetRows(y...)
		m.grid.SetColumns(x...)
		m.grid.AddItem(modal, r, w, span, span, int(height), int(width), true)
	}
	m.hasModal = true
	m.modal = modal
}

// RemoveModal removes modal
func (m *ModalLayout) RemoveModal(modal Modal) {
	if m.hasModal {
		modal.SetVisible(false)
		m.grid.RemoveItem(modal)
		m.hasModal = false
		m.modal = nil
		if m.customGrid {
			m.grid.SetRows(m.gridAxisY...)
			m.grid.SetColumns(m.gridAxisX...)
			m.customGrid = false
		}
	}
}