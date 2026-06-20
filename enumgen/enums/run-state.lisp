(enum
  :package runstate
  :type RunState
  :code-type RunStateCode
  :string-type RunStateString
  :const-prefix State
  :all AllStates
  :code-all AllRunStateCodes
  (value StatePending "pending")
  (value StatePreparing "preparing")
  (value StateStartingProvider "starting_provider")
  (value StateHandshaking "handshaking")
  (value StateRunning "running")
  (value StateWaitingToolApproval "waiting_tool_approval")
  (value StateToolRunning "tool_running")
  (value StateWaitingProvider "waiting_provider")
  (value StateCompleting "completing")
  (value StateCompleted "completed" :terminal true)
  (value StateFailed "failed" :terminal true)
  (value StateCancelled "cancelled" :terminal true)
  (value StateTimedOut "timed_out" :terminal true)
  (value StateIdleStopped "idle_stopped" :terminal true))
