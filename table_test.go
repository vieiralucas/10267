package table

import "testing"

func TestCreateTableMxM(t *testing.T) {
	table := CreateTable(2, 2)

	if height := table.Height(); height != 2 {
		t.Errorf("Expect table height to be 2, got: %v", height)
	}

	if height := table.Width(); height != 2 {
		t.Errorf("Expect table width to be 2, got: %v", height)
	}

}

func TestCreateTableMxN(t *testing.T) {
	table := CreateTable(2, 3)

	if height := table.Height(); height != 3 {
		t.Errorf("Expect table height to be 3, got: %v", height)
	}

	if width := table.Width(); width != 2 {
		t.Errorf("Expect table width to be 2, got: %v", width)
	}
}

func TestCreateTableInitialize0(t *testing.T) {
	table := CreateTable(3, 4)

	for r := 0; r < len(table.rows); r++ {
		for c := 0; c < len(table.rows[r]); c++ {
			if value := table.rows[r][c]; value != 0 {
				t.Errorf("Expect pos %v, %v to be 0, got %v", r, c, value)
			}
		}
	}
}

func TestClear(t *testing.T) {
	table := CreateTable(3, 4)
	table.rows[1][2] = 2

	table.Clear()

	for r := 0; r < len(table.rows); r++ {
		for c := 0; c < len(table.rows[r]); c++ {
			if value := table.rows[r][c]; value != 0 {
				t.Errorf("Expect pos %v, %v to be 0, got %v", r, c, value)
			}
		}
	}
}

func TestGetPixel(t *testing.T) {
	table := CreateTable(3, 4)
	table.rows[2][1] = 2

	if pixel := table.GetPixel(1, 2); pixel != 2 {
		t.Errorf("Expect pos 1, 2 to be 2, got %v", pixel)
	}
}

func TestPaintPixel(t *testing.T) {
	table := CreateTable(3, 4)
	table.PaintPixel(1, 2, 3)

	if color := table.GetPixel(1, 2); color != 3 {
		t.Errorf("Expect pos 1, 2 to be 3, got %v", color)
	}
}

func TestPaintVertical(t *testing.T) {
	table := CreateTable(3, 4)
	table.PaintVertical(2, 1, 2, 4)

	if p := table.GetPixel(2, 1); p != 4 {
		t.Errorf("Expect pos 2, 1 to be 4, got %v", p)
	}

	if p := table.GetPixel(2, 2); p != 4 {
		t.Errorf("Expect pos 2, 2 to be 4, got %v", p)
	}
}

func TestPaintHorizontal(t *testing.T) {
	table := CreateTable(4, 3)
	table.PaintHorizontal(1, 3, 2, 4)

	if p := table.GetPixel(1, 2); p != 4 {
		t.Errorf("Expect pos 1, 2 to be 4, got %v", p)
	}

	if p := table.GetPixel(2, 2); p != 4 {
		t.Errorf("Expect pos 2, 2 to be 4, got %v", p)
	}

	if p := table.GetPixel(3, 2); p != 4 {
		t.Errorf("Expect pos 3, 2 to be 4, got %v", p)
	}
}

func TestFillRect(t *testing.T) {
	table := CreateTable(5, 5)
	table.FillRect(1, 1, 3, 3, 5)

	var toTest = []struct {
		actual   col
		expected col
	}{
		{table.GetPixel(1, 1), 5},
		{table.GetPixel(1, 2), 5},
		{table.GetPixel(1, 3), 5},
		{table.GetPixel(2, 1), 5},
		{table.GetPixel(2, 2), 5},
		{table.GetPixel(2, 3), 5},
		{table.GetPixel(3, 1), 5},
		{table.GetPixel(3, 2), 5},
		{table.GetPixel(3, 3), 5},
	}

	for _, tt := range toTest {
		if tt.actual != tt.expected {
			t.Errorf("Expect %v, got %v", tt.expected, tt.actual)
		}
	}
}

func TestFillRegion(t *testing.T) {
	table := FromSlice([][]int{
		{0, 1, 0, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 0, 0, 1},
		{0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0},
	})

	table.FillRegion(1, 1, 2)

	expected := FromSlice([][]int{
		{0, 1, 0, 0, 0},
		{1, 2, 1, 1, 0},
		{1, 2, 2, 1, 0},
		{0, 1, 2, 2, 1},
		{0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}).ToString()
	actual := table.ToString()

	if expected != actual {
		t.Errorf("Expected\n%v\ngot\n%v", expected, actual)
	}
}

func TestFromSlice(t *testing.T) {
	table := FromSlice([][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	})

	if width := table.Width(); width != 5 {
		t.Errorf("Expect table width to be 5, got: %v", width)
	}

	if height := table.Height(); height != 5 {
		t.Errorf("Expect table height to be 5, got: %v", height)
	}

	if pixel := table.GetPixel(1, 1); pixel != 1 {
		t.Errorf("Expect pos 1, 1 to be 1, got %v", pixel)
	}
}

func TestToString(t *testing.T) {
	table := FromSlice([][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	})
	expected := "0, 0, 0, 0, 0, \n0, 1, 0, 0, 0, \n0, 0, 0, 0, 0, \n0, 0, 0, 0, 0, \n0, 0, 0, 0, 0, \n"
	actual := table.ToString()

	if expected != actual {
		t.Errorf("Expected\n%v\ngot\n%v", expected, actual)
	}
}
