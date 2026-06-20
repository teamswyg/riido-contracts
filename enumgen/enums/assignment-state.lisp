(enum
  :package assignment
  :type AssignmentState
  :code-type AssignmentStateCode
  :string-type AssignmentStateString
  :const-prefix Assignment
  :all AllAssignmentStates
  :code-all AllAssignmentStateCodes
  (value AssignmentQueued "queued")
  (value AssignmentLeased "leased" :agent-active true)
  (value AssignmentReady "ready" :agent-active true)
  (value AssignmentRunning "running" :agent-active true)
  (value AssignmentCancelling "cancelling")
  (value AssignmentCancelled "cancelled" :terminal true)
  (value AssignmentCompleted "completed" :terminal true)
  (value AssignmentFailed "failed" :terminal true))
