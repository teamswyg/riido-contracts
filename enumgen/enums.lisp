(enum-gen
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

  (enum
    :package ir
    :type EventType
    :code-type EventTypeCode
    :string-type EventTypeString
    :const-prefix Event
    :all AllEventTypes
    :code-all AllEventTypeCodes
    (value EventTaskCreated "TaskCreated" :transition true :native-config forbidden)
    (value EventTaskQueued "TaskQueued" :transition true :native-config forbidden)
    (value EventTaskClaimed "TaskClaimed" :transition true :native-config pre-execute)
    (value EventWorkdirPreparing "WorkdirPreparing" :transition true :native-config pre-execute)
    (value EventRuntimePinned "RuntimePinned" :transition true :native-config pre-execute)
    (value EventRunStarted "RunStarted" :transition true)
    (value EventInputRequested "InputRequested" :transition true)
    (value EventInputProvided "InputProvided" :transition true)
    (value EventBlockerRaised "BlockerRaised" :transition true :native-config phase-dependent)
    (value EventBlockerResolved "BlockerResolved" :transition true :native-config phase-dependent)
    (value EventBlockerResolvedRequeue "BlockerResolvedRequeue" :transition true :native-config phase-dependent)
    (value EventRunReportedDone "RunReportedDone" :transition true)
    (value EventValidationPassed "ValidationPassed" :transition true)
    (value EventValidationFailed "ValidationFailed" :transition true)
    (value EventReviewRequested "ReviewRequested" :transition true)
    (value EventAutoApproved "AutoApproved" :transition true)
    (value EventHumanApproved "HumanApproved" :transition true)
    (value EventHumanRejected "HumanRejected" :transition true)
    (value EventReworkAccepted "ReworkAccepted" :transition true :native-config phase-dependent)
    (value EventTaskCancelled "TaskCancelled" :transition true :native-config phase-dependent)
    (value EventTaskTimedOut "TaskTimedOut" :transition true :native-config phase-dependent)
    (value EventRuntimePinViolated "RuntimePinViolated" :transition true :native-config phase-dependent)
    (value EventTaskFailed "TaskFailed" :transition true :native-config phase-dependent)
    (value EventRuntimeRegistered "RuntimeRegistered" :native-config forbidden)
    (value EventRuntimeRejected "RuntimeRejected" :native-config forbidden)
    (value EventRuntimeFingerprintChanged "RuntimeFingerprintChanged" :native-config forbidden)
    (value EventCapabilityReevaluated "CapabilityReevaluated" :native-config forbidden)
    (value EventLeaseInvalidated "LeaseInvalidated" :native-config forbidden)
    (value EventRuntimeHandshakeOK "RuntimeHandshakeOK" :native-config pre-execute)
    (value EventTextDelta "TextDelta")
    (value EventReasoningDelta "ReasoningDelta")
    (value EventToolCallStarted "ToolCallStarted")
    (value EventToolCallFinished "ToolCallFinished")
    (value EventFileChanged "FileChanged")
    (value EventCommandStarted "CommandStarted")
    (value EventCommandFinished "CommandFinished")
    (value EventSessionPinned "SessionPinned")
    (value EventApprovalRequested "ApprovalRequested")
    (value EventApprovalResolved "ApprovalResolved")
    (value EventStatusUpdate "StatusUpdate")
    (value EventUsageDelta "UsageDelta")
    (value EventLogLine "LogLine")
    (value EventProviderUnknownEvent "ProviderUnknownEvent")
    (value EventValidationStarted "ValidationStarted")
    (value EventValidationRuleExecuted "ValidationRuleExecuted")
    (value EventWorkdirCreated "WorkdirCreated" :native-config pre-execute)
    (value EventNativeConfigInjected "NativeConfigInjected")
    (value EventWorkdirArchived "WorkdirArchived")
    (value EventConfigTemplateReinjected "ConfigTemplateReinjected")
    (value EventPolicyBundleLoaded "PolicyBundleLoaded" :native-config forbidden)
    (value EventPolicyBundleSwitched "PolicyBundleSwitched" :native-config forbidden)
    (value EventPolicyViolationDetected "PolicyViolationDetected")
    (value EventSecretsScopeIssued "SecretsScopeIssued")
    (value EventSecretsScopeRevoked "SecretsScopeRevoked")
    (value EventUpgradeDetected "UpgradeDetected" :native-config forbidden)
    (value EventUpgradePolicyReevaluated "UpgradePolicyReevaluated" :native-config forbidden)
    (value EventDrainStarted "DrainStarted" :native-config forbidden)
    (value EventDrainTimedOut "DrainTimedOut" :native-config forbidden)
    (value EventTaskHandedOff "TaskHandedOff")
    (value EventCorrection "Correction")
    (value EventOperatorNote "OperatorNote")
    (value EventSnapshot "Snapshot"))

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

  (enum
    :package assignment
    :type PollAction
    :code-type PollActionCode
    :string-type PollActionString
    :const-prefix Poll
    :all AllPollActions
    :code-all AllPollActionCodes
    (value PollNone "none")
    (value PollStart "start")
    (value PollCancel "cancel")
    (value PollActive "active")))
