(enum
  :package task
  :type TaskState
  :code-type TaskStateCode
  :string-type TaskStateString
  :const-prefix State
  :all AllStates
  :code-all AllTaskStateCodes
  (value StateCreated "Created")
  (value StateQueued "Queued")
  (value StateClaimed "Claimed")
  (value StatePreparing "Preparing")
  (value StateRunning "Running" :active true)
  (value StateNeedsInput "NeedsInput")
  (value StateBlocked "Blocked")
  (value StateValidating "Validating" :active true)
  (value StatePatchReady "PatchReady")
  (value StateHumanReview "HumanReview")
  (value StateReworkQueued "ReworkQueued")
  (value StateCompleted "Completed" :terminal true)
  (value StateFailed "Failed" :terminal true)
  (value StateCancelled "Cancelled" :terminal true)
  (value StateTimedOut "TimedOut" :terminal true))
