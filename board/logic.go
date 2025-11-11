package board

// PlayerWins erwartet ein Zeichen.
// Gibt `true` zurück, wenn es eine Zeile, Spalte oder Diagonale gibt, die nur aus dem Zeichen besteht.
func (b Board) PlayerWins(s string) bool {
	for i := range b {
		if b.RowContainsOnly(i, s) || b.ColContainsOnly(i, s) {
			return true
		}
	}
	if b.DiagContainsOnly(0, s) || b.DiagContainsOnly(1, s) {
		return true
	}
	return false
}

// GameOver gibt `true` zurück, wenn das Spiel vorbei ist.
// Das Spiel ist vorbei, wenn ein Spieler gewonnen hat oder das Board voll ist.
func (b Board) GameOver() bool {
	if b.PlayerWins("X") || b.PlayerWins("O") || b.Full() {
		return true
	}
	return false
}
