package players

import (
	"dots/bitset"
	"dots/board"
)

func SquaredHeuristic(board board.Board) (heur int) {
	corner_mask := bitset.Bitset(0)
	corner_mask.SetBit(0).SetBit(7).SetBit(56).SetBit(63)

	me_corners := (corner_mask & board.Me()).Count()
	opp_corners := (corner_mask & board.Opp()).Count()
	corner_diff := me_corners - opp_corners

	me_moves := board.Moves().Count()
	board.SwitchTurn()
	opp_moves := board.Moves().Count()
	move_diff := me_moves - opp_moves

	heur = int((3 * corner_diff) + move_diff)
	return
}
