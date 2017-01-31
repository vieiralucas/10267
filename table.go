package table

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

// Creates a new table M × N. All the pixels are colored in white (O)
func CreateTable(m int, n int) *Table {
	rows := make([]row, n)

	for i := 0; i < n; i++ {
		rows[i] = make([]col, m)
	}

	return &Table{rows}
}
