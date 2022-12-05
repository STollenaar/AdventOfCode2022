package internal

type Row[T comparable] []T

type Grid[T comparable] struct {
	Rows  []Row[T]
	empty T
}

// Adds object to end of the slice
func (g *Grid[T]) AddUnsafeToColumn(object T, row int) bool {
	g.Rows[row] = append(g.Rows[row], object)
	return true
}

// Adds object to beginning of the slice
func (g *Grid[T]) ShiftUnsafeToColumn(object T, row int) bool {
	g.Rows[row] = append([]T{object}, g.Rows[row]...)
	return true
}

// Adds objects safely to end of the slice
func (g *Grid[T]) AddSafeToColumn(object T, row int) bool {
	g.GetSafeRow(row)
	g.AddUnsafeToColumn(object, row)
	return true
}

// Adds objects safely to beginning of the slice
func (g *Grid[T]) ShiftSafeToColumn(object T, row int) bool {
	g.GetSafeRow(row)
	g.ShiftUnsafeToColumn(object, row)
	return true
}

func (g *Grid[T]) SetUnsafeColumn(object T, x, y int) bool {
	g.Rows[y][x] = object
	return true
}

func (g *Grid[T]) SetSafeColumn(object T, x, y int) bool {
	g.GetSafeColumn(x, y)
	g.GetSafeRow(y)[x] = object
	return true
}

// Adds row to end of the slice
func (g *Grid[T]) AddRow(row Row[T]) bool {
	g.Rows = append(g.Rows, row)
	return true
}

// Adds row to beginning of the slice
func (g *Grid[T]) ShiftRow(row Row[T]) bool {
	g.Rows = append([]Row[T]{row}, g.Rows...)
	return true
}

func (g *Grid[T]) GetUnsafeRow(y int) Row[T] {
	return g.Rows[y]
}

// Adds rows if needed, or shifts once and returns the new row
func (g *Grid[T]) GetSafeRow(y int) Row[T] {
	if y < 0 {
		g.ShiftRow(Row[T]{})
		return g.Rows[0]
	}
	for len(g.Rows) <= y {
		g.AddRow(Row[T]{})
	}
	return g.Rows[y]
}

func (g *Grid[T]) GetUnsafeColumn(x, y int) T {
	return g.GetUnsafeRow(y)[x]
}

// Tries to safely get a column, will apply GetSafeRow function. Throws error if the x value is not valid
func (g *Grid[T]) GetSafeColumn(x, y int) T {
	return g.GetSafeRow(y)[x]
}

// Getting the max Y of a non-empty given column
func (g *Grid[T]) GetHeight(x int) int {
	for y := range g.Rows {
		if g.GetSafeColumn(x, y) != g.empty {
			return y
		}
	}
	return len(g.Rows) - 1
}
