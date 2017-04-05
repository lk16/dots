package minimax

const (
	Max_exact_heuristic = 64
	Min_exact_heuristic = -Max_exact_heuristic
	Exact_score_factor  = 10000
	Max_heuristic       = Exact_score_factor * Max_exact_heuristic
	Min_heuristic       = Exact_score_factor * Min_exact_heuristic
)
