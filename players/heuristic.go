package players

import (
	"math/bits"

	"dots/board"
)

func Squared(board board.Board) (heur int) {
	corner_mask := 1<<0 | 1<<7 | uint64(1)<<56 | uint64(1)<<63

	me_corners := bits.OnesCount64(corner_mask & board.Me())
	opp_corners := bits.OnesCount64(corner_mask & board.Opp())
	corner_diff := me_corners - opp_corners

	me_moves := bits.OnesCount64(board.Moves())
	opp_moves := bits.OnesCount64(board.OpponentMoves())
	move_diff := me_moves - opp_moves

	heur = int((3 * corner_diff) + move_diff)
	return
}
