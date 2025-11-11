package board

import (
	"boardgames/rows"
	"fmt"
	"strings"
)

type Board []rows.Row

// New erwartet Höhe, Breite und ein Zeichen.
// Liefert ein neues `Board` zurück, das mit dem Zeichen gefüllt ist.
func New(height, width int, fill string) Board {
	board := make(Board, height)
	for i := 0; i < height; i++ {
		board[i] = rows.New(width, fill)
	}
	return board
}

// String gibt das `Board` als String zurück.
// Die Zeilen sind durch Trenner der Form `+---+---+---+` getrennt.
func (b Board) String() string {
	rowStrings := make([]string, len(b))
	divider := fmt.Sprintf("\n%s+\n", strings.Repeat("+---", len(b)))

	for i, row := range b {
		rowStrings[i] = row.String()
	}

	return fmt.Sprintf("%s%s%s", divider, strings.Join(rowStrings, divider), divider)
}

// Row erwartet eine Zeilennummer und liefert diese Zeile zurück.
func (b Board) Row(row int) rows.Row {
	// TODO
	return b[row]
}

// Set erwartet eine Spaltennummer und liefert diese Spalte zurück.
// Der Rückgabetype ist Row, da Zeilen und Spalten gleich sind.
func (b Board) Col(col int) rows.Row {
	c := make(rows.Row, len(b))
	for i := range b {
		c[i] = b[i][col]
	}

	return c
}

// Diag erwartet eine Diagonalennummer und liefert diese Diagonale zurück.
// Der Rückgabetype ist Row, da Diagonalen und Zeilen gleich sind.
// Die Diagonalennummer ist 0 für die Hauptdiagonale und 1 für die Nebendiagonale.
func (b Board) Diag(diag int) rows.Row {
	d := make(rows.Row, len(b))
	if diag == 0 {
		for i := range b {
			d[i] = b[i][i]
		}
	}
	if diag == 1 {
		for l := range b {
			d[l] = b[l][len(b)-1-l]
		}
	}
	return d
}

// Set erwartet eine Zeilen- und eine Spaltennummer und ein Zeichen.
// Setzt das Zeichen an die entsprechende Stelle.
func (b Board) Set(row, col int, fill string) {
	b[row][col] = fill
}

// Full gibt `true` zurück, wenn das Board voll ist.
func (b Board) Full() bool {
	for i := range b {
		for j := range b[i] {
			if b[i][j] != "X" && b[i][j] != "O" {
				return false
			}

		}
	}
	return true
}

// RowContainsOnly erwartet eine Zeilennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Zeile nur aus dem Zeichen besteht.
func (b Board) RowContainsOnly(row int, s string) bool {
	for i := range b[row] {
		if b[row][i] != s {
			return false
		}
	}
	return true
}

// ColContainsOnly erwartet eine Spaltennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Spalte nur aus dem Zeichen besteht.
func (b Board) ColContainsOnly(col int, s string) bool {
	for i := range b {
		if b[i][col] != s {
			return false
		}
	}
	return true
}

// DiagContainsOnly erwartet eine Diagonalennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Diagonale nur aus dem Zeichen besteht.
// Die Diagonalennummer ist 0 für die Hauptdiagonale und 1 für die Nebendiagonale.
func (b Board) DiagContainsOnly(diag int, s string) bool {
	if diag == 0 {
		for i := range b {
			if b[i][i] != s {
				return false
			}
		}
	}
	if diag == 1 {
		for l := range b {
			if b[l][len(b)-1-l] != s {
				return false
			}
		}
	}

	return true
}
