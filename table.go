package table

import (
	"bytes"
	"strconv"
)

type col int

type row []col

type Table struct {
	rows []row
}

func (t *Table) Height() int {
	return len(t.rows)
}

func (t *Table) Width() int {
	return len(t.rows[0])
}

// Clears the table. The size remains the same. All the pixels became white (O)
func (t *Table) Clear() {
	t.rows = make([]row, t.Height())
}

func (t *Table) GetPixel(x int, y int) col {
	return t.rows[y][x]
}

// Colors the pixel with coordinates (X, Y ) in colour C
func (t *Table) PaintPixel(x int, y int, color col) {
	t.rows[y][x] = color
}

// Paints the vertical segment in the column X between the rows Y1 and Y2
// inclusive in colour C
func (t *Table) PaintVertical(x int, y1 int, y2 int, color col) {
	for y := y1; y <= y2; y++ {
		t.PaintPixel(x, y, color)
	}
}

// Paints the horizontal segment in the row Y between the columns X1 and X2
// inclusive in colour C
func (t *Table) PaintHorizontal(x1 int, x2 int, y int, color col) {
	for x := x1; x <= x2; x++ {
		t.PaintPixel(x, y, color)
	}
}

// Draws the filled rectangle in colour C. (X1, Y1) is the upper left corner,
// (X2, Y2) is the lower right corner of the rectangle
func (t *Table) FillRect(x1 int, y1 int, x2 int, y2 int, color col) {
	for i := y1; i <= y2; i++ {
		t.PaintHorizontal(x1, x2, i, color)
	}
}

func iter(t *Table, starter col, color col, x int, y int) {
	// out of bounds, stop recursion
	if x < 0 || x >= t.Width() || y < 0 || y >= t.Height() {
		return
	}

	curr := t.GetPixel(x, y)
	// reach wall, stop recursion
	if curr != starter {
		return
	}

	t.PaintPixel(x, y, color)
	iter(t, starter, color, x+1, y) // right
	iter(t, starter, color, x-1, y) // left
	iter(t, starter, color, x, y-1) // up
	iter(t, starter, color, x, y+1) // down
}

// Fills the region with the colour C. The region R to be filled
// is defined as follows. The pixel (X, Y ) belongs to this region.
// The other pixel belongs to the region R if and only if it has
// the same colour as pixel (X, Y ) and a common side with any
// pixel which belongs to this region.
func (t *Table) FillRegion(x int, y int, color col) {
	starter := t.GetPixel(x, y)

	iter(t, starter, color, x, y)
}

// Returns human readable string representation of the table
func (t *Table) ToString() string {
	var buffer bytes.Buffer

	for _, r := range t.rows {
		for _, c := range r {
			buffer.WriteString(strconv.Itoa(int(c)))
			buffer.WriteString(", ")
		}

		buffer.WriteString("\n")
	}

	return buffer.String()
}

// Creates a new table M × N. All the pixels are colored in white (O)
func CreateTable(m int, n int) *Table {
	rows := make([]row, n)

	for i := 0; i < n; i++ {
		rows[i] = make([]col, m)
	}

	return &Table{rows}
}

// Creates a table from an slice of slices
func FromSlice(s [][]int) *Table {
	rows := make([]row, len(s))

	for i := 0; i < len(s); i++ {
		rows[i] = make([]col, len(s[i]))
		for j := 0; j < len(s[i]); j++ {
			rows[i][j] = col(s[i][j])
		}
	}

	return &Table{rows}
}
