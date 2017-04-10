package minimax

import (
	"dots/board"
)

type AlphaBeta struct {
	heuristic Heuristic
}

func (alphabeta *AlphaBeta) Search(board board.Board, depth_left uint, heuristic Heuristic, alpha int) (heur int) {
	alphabeta.heuristic = heuristic
	heur = alphabeta.doAlphaBeta(board, depth_left, alpha, Max_heuristic)
	return
}

func (alphabeta *AlphaBeta) polish(heur, alpha, beta int) (outheur int) {
	if heur < alpha {
		outheur = alpha
	} else if heur > beta {
		outheur = beta
	} else {
		outheur = heur
	}
	return
}

func (alphabeta *AlphaBeta) doAlphaBeta(board board.Board, depth_left uint, alpha, beta int) (heur int) {

	if depth_left == 0 {
		heur = alphabeta.polish(alphabeta.heuristic(board), alpha, beta)
		return
	}

	if moves := board.Moves(); moves != 0 {
		heur = alpha
		for child := range board.GenChildren() {
			child_heur := -alphabeta.doAlphaBeta(child, depth_left-1, -beta, -heur)
			if child_heur > heur {
				heur = child_heur
			}
			if heur >= beta {
				heur = beta
				break
			}
		}
		return
	}

	board.SwitchTurn()
	if moves := board.Moves(); moves != 0 {
		heur = -alphabeta.doAlphaBeta(board, depth_left, -beta, -alpha)
		return
	}

	heur = alphabeta.polish(Exact_score_factor*board.ExactScore(), alpha, beta)
	return
}

func (alphabeta *AlphaBeta) ExactSearch(board board.Board, alpha int) (heur int) {
	heur = -alphabeta.doAlphaBetaExact(board, alpha, Max_exact_heuristic)
	return
}

func (alphabeta *AlphaBeta) doAlphaBetaExact(board board.Board, alpha, beta int) (heur int) {

	if moves := board.Moves(); moves != 0 {
		heur = alpha
		for child := range board.GenChildren() {
			child_heur := -alphabeta.doAlphaBetaExact(child, -beta, -heur)
			if child_heur > heur {
				heur = child_heur
			}
			if heur >= beta {
				heur = beta
				break
			}
		}
		return
	}

	board.SwitchTurn()
	if moves := board.Moves(); moves != 0 {
		heur = -alphabeta.doAlphaBetaExact(board, -beta, -alpha)
		return
	}

	heur = alphabeta.polish(board.ExactScore(), alpha, beta)
	return
}
