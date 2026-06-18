package metadatakeys

func All() []Key {
	return []Key{
		WorkspaceID, Workspace, TaskID, RunID,
		AgentName, AgentIdentity, Workflow,
		WorkdirRoot, Workdir, OutputDir, LogsDir, ArtifactsDir,
		NativeConfigDir, NativeConfigHome, IRDir, NativeConfigVersion,
		RequiredSurfaces, AllowExperimentalRuntime,
		RuntimeLeaseID, RuntimeFencingToken, RuntimeCapabilityFingerprint,
		ProgressMessageCode, ProgressMessageKey, ProgressMessageArgPrefix, ThreadProgressSeq,
		AssignmentRecovery, AssignmentResultStatus, AssignmentFailureCategory, AssignmentEventKey,
		HTTPRequestMethod, HTTPRoute, HTTPResponseStatusCode, HTTPStatusCode,
		AWSService, AWSOperation, AWSRegion,
		RiidoTraceSurface, RiidoStoreOperation, RiidoTaskContextOperation, RiidoPollAction,
	}
}
