(transitions
  :package assignment
  :name AssignmentTransitionCode
  :from-enum assignment.AssignmentState
  :to-enum assignment.AssignmentState
  :all AssignmentTransitionCodes
  :validate CanTransitionCode
  :allow-same true
  :fsm-name AssignmentFSM
  :fsm-type-union AssignmentPollingFSM
  :pattern-source fsmgen/patterns.lisp
  :conformance-profile RiidoPublicFlatFSM
  :patterns (PatternFlat PatternStateDriven PatternExplicitBoundary PatternMultiTerminal PatternSameStateAllowed)
  :start-points (AssignmentQueued)
  :end-points (AssignmentCancelled AssignmentCompleted AssignmentFailed)
  :readme-section assignment
  (transition AssignmentQueued AssignmentLeased)
  (transition AssignmentQueued AssignmentCancelling)
  (transition AssignmentQueued AssignmentCancelled)
  (transition AssignmentQueued AssignmentFailed)
  (transition AssignmentLeased AssignmentReady)
  (transition AssignmentLeased AssignmentRunning)
  (transition AssignmentLeased AssignmentCancelling)
  (transition AssignmentLeased AssignmentCancelled)
  (transition AssignmentLeased AssignmentFailed)
  (transition AssignmentReady AssignmentRunning)
  (transition AssignmentReady AssignmentCancelling)
  (transition AssignmentReady AssignmentCancelled)
  (transition AssignmentReady AssignmentFailed)
  (transition AssignmentRunning AssignmentCompleted)
  (transition AssignmentRunning AssignmentFailed)
  (transition AssignmentRunning AssignmentCancelling)
  (transition AssignmentRunning AssignmentCancelled)
  (transition AssignmentCancelling AssignmentCancelled)
  (transition AssignmentCancelling AssignmentFailed))
