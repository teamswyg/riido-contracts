(fsm-pattern-gen
  (sum-type
    :package fsmmeta
    :type FSMPattern
    :code-type FSMPatternCode
    :string-type FSMPatternString
    :const-prefix Pattern
    :output fsmmeta/fsm_pattern_gen.go
    (variant PatternFlat "flat")
    (variant PatternEventDriven "event-driven")
    (variant PatternStateDriven "state-driven")
    (variant PatternExplicitBoundary "explicit-boundary")
    (variant PatternMultiTerminal "multi-terminal")
    (variant PatternMultiTargetEvent "multi-target-event")
    (variant PatternSameStateAllowed "same-state-allowed")
    (variant PatternHierarchical "hierarchical")
    (variant PatternParallel "parallel")
    (variant PatternHistory "history")
    (variant PatternGuardedTransition "guarded-transition")
    (variant PatternEntryAction "entry-action")
    (variant PatternExitAction "exit-action")
    (variant PatternInternalTransition "internal-transition")
    (variant PatternEventlessTransition "eventless-transition"))

  (conformance-profile
    :name RiidoPublicFlatFSM
    :allowed-patterns
      (PatternFlat
       PatternEventDriven
       PatternStateDriven
       PatternExplicitBoundary
       PatternMultiTerminal
       PatternMultiTargetEvent
       PatternSameStateAllowed)
    :rejected-patterns
      (PatternHierarchical
       PatternParallel
       PatternHistory
       PatternGuardedTransition
       PatternEntryAction
       PatternExitAction
       PatternInternalTransition
       PatternEventlessTransition)))
