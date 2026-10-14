(transitions
  :package task
  :name TaskTransitionCode
  :from-enum task.TaskState
  :to-enum task.TaskState
  :event-enum ir.EventType
  :all LegalTransitionCodes
  :validate ValidateTransitionCode
  :fsm-name TaskFSM
  :fsm-type-union TaskLifecycleFSM
  :pattern-source fsmgen/patterns.lisp
  :conformance-profile RiidoPublicFlatFSM
  :patterns (PatternFlat PatternEventDriven PatternExplicitBoundary PatternMultiTerminal PatternMultiTargetEvent)
  :start-points (StateCreated)
  :end-points (StateCompleted StateFailed StateCancelled StateTimedOut)
  :readme-section task
  (transition StateCreated StateQueued EventTaskQueued)
  (transition StateQueued StateClaimed EventTaskClaimed)
  (transition StateClaimed StatePreparing EventWorkdirPreparing)
  (transition StatePreparing StateRunning EventRunStarted)
  (transition StateRunning StateValidating EventRunReportedDone)
  (transition StateValidating StatePatchReady EventValidationPassed)
  (transition StatePatchReady StateHumanReview EventReviewRequested)
  (transition StateHumanReview StateCompleted EventHumanApproved)
  (transition StatePatchReady StateCompleted EventAutoApproved)
  (transition StateRunning StateNeedsInput EventInputRequested)
  (transition StateNeedsInput StateRunning EventInputProvided)
  (transition StatePreparing StateBlocked EventBlockerRaised)
  (transition StateRunning StateBlocked EventBlockerRaised)
  (transition StateBlocked StateRunning EventBlockerResolved)
  (transition StateBlocked StateQueued EventBlockerResolvedRequeue)
  (transition StateValidating StateFailed EventValidationFailed)
  (transition StateHumanReview StateReworkQueued EventHumanRejected)
  (transition StateHumanReview StateCancelled EventHumanRejected)
  (transition StateReworkQueued StateQueued EventReworkAccepted)
  (transition StateRunning StateFailed EventRuntimePinViolated)
  (transition StateValidating StateFailed EventRuntimePinViolated)
  (transition StateClaimed StateFailed EventTaskFailed)
  (transition StatePreparing StateFailed EventTaskFailed)
  (transition StateRunning StateFailed EventTaskFailed)
  (transition StateNeedsInput StateFailed EventTaskFailed)
  (transition StateBlocked StateFailed EventTaskFailed)
  (transition StateValidating StateFailed EventTaskFailed)
  (transition StateCreated StateCancelled EventTaskCancelled)
  (transition StateQueued StateCancelled EventTaskCancelled)
  (transition StateClaimed StateCancelled EventTaskCancelled)
  (transition StatePreparing StateCancelled EventTaskCancelled)
  (transition StateRunning StateCancelled EventTaskCancelled)
  (transition StateNeedsInput StateCancelled EventTaskCancelled)
  (transition StateBlocked StateCancelled EventTaskCancelled)
  (transition StateValidating StateCancelled EventTaskCancelled)
  (transition StatePatchReady StateCancelled EventTaskCancelled)
  (transition StateHumanReview StateCancelled EventTaskCancelled)
  (transition StateReworkQueued StateCancelled EventTaskCancelled)
  (transition StateRunning StateTimedOut EventTaskTimedOut)
  (transition StateNeedsInput StateTimedOut EventTaskTimedOut)
  (transition StateBlocked StateTimedOut EventTaskTimedOut)
  (transition StateValidating StateTimedOut EventTaskTimedOut)
  (transition StateHumanReview StateTimedOut EventTaskTimedOut))
