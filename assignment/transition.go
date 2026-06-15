package assignment

func CanTransition(from, to AssignmentState) bool {
	return CanTransitionCode(from.Code(), to.Code())
}
