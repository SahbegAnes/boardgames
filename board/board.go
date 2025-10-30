package board

import (
	"boardgames/rows"
	"strings"
)

type Board []rows.Row

// New erwartet Höhe, Breite und ein Zeichen.
// Liefert ein neues `Board` zurück, das mit dem Zeichen gefüllt ist.
func New(height, width int, fill string) Board {
	board := make(Board, height)
	// TODO
	return board
}

// String gibt das `Board` als String zurück.
// Die Zeilen sind durch Trenner der Form `+---+---+---+` getrennt.
func (b Board) String() string {
	rowStrings := make([]string, len(b))
	divider := "\n"
	// TODO
	return strings.Join(rowStrings, divider)
}

// Row erwartet eine Zeilennummer und liefert diese Zeile zurück.
func (b Board) Row(row int) rows.Row {
	// TODO
	return rows.Row{}
}

// Set erwartet eine Spaltennummer und liefert diese Spalte zurück.
// Der Rückgabetype ist Row, da Zeilen und Spalten gleich sind.
func (b Board) Col(col int) rows.Row {
	c := make(rows.Row, len(b))
	// TODO
	return c
}

// Diag erwartet eine Diagonalennummer und liefert diese Diagonale zurück.
// Der Rückgabetype ist Row, da Diagonalen und Zeilen gleich sind.
// Die Diagonalennummer ist 0 für die Hauptdiagonale und 1 für die Nebendiagonale.
func (b Board) Diag(diag int) rows.Row {
	d := make(rows.Row, len(b))
	// TODO
	return d
}

// Set erwartet eine Zeilen- und eine Spaltennummer und ein Zeichen.
// Setzt das Zeichen an die entsprechende Stelle.
func (b Board) Set(row, col int, fill string) {
	// TODO
}

// Full gibt `true` zurück, wenn das Board voll ist.
func (b Board) Full() bool {
	// TODO
	return true
}

// RowContainsOnly erwartet eine Zeilennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Zeile nur aus dem Zeichen besteht.
func (b Board) RowContainsOnly(row int, s string) bool {
	// TODO
	return false
}

// ColContainsOnly erwartet eine Spaltennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Spalte nur aus dem Zeichen besteht.
func (b Board) ColContainsOnly(col int, s string) bool {
	// TODO
	return false
}

// DiagContainsOnly erwartet eine Diagonalennummer und ein Zeichen.
// Gibt `true` zurück, wenn die Diagonale nur aus dem Zeichen besteht.
// Die Diagonalennummer ist 0 für die Hauptdiagonale und 1 für die Nebendiagonale.
func (b Board) DiagContainsOnly(diag int, s string) bool {
	// TODO
	return false
}
