package metadatakeys

type Key string

const (
	WorkspaceID Key = "workspace_id"
	Workspace   Key = "workspace"
	TaskID      Key = "task_id"
	RunID       Key = "run_id"

	AgentName     Key = "agent_name"
	AgentIdentity Key = "agent_identity"
	Workflow      Key = "workflow"

	WorkdirRoot         Key = "workdir_root"
	Workdir             Key = "workdir"
	OutputDir           Key = "output_dir"
	LogsDir             Key = "logs_dir"
	ArtifactsDir        Key = "artifacts_dir"
	NativeConfigDir     Key = "native_config_dir"
	NativeConfigHome    Key = "native_config_home"
	IRDir               Key = "ir_dir"
	NativeConfigVersion Key = "native_config_version"

	RequiredSurfaces         Key = "required_surfaces"
	AllowExperimentalRuntime Key = "allow_experimental_runtime"

	RuntimeLeaseID               Key = "runtime_lease_id"
	RuntimeFencingToken          Key = "runtime_fencing_token"
	RuntimeCapabilityFingerprint Key = "runtime_capability_fingerprint"

	ProgressMessageCode      Key = "riido_progress_message_code"
	ProgressMessageKey       Key = "riido_progress_message_key"
	ProgressMessageArgPrefix Key = "riido_progress_message_arg."
	ThreadProgressSeq        Key = "thread_progress_seq"

	HTTPRequestMethod      Key = "http.request.method"
	HTTPRoute              Key = "http.route"
	HTTPResponseStatusCode Key = "http.response.status_code"
	HTTPStatusCode         Key = "http.status_code"

	AWSService   Key = "aws.service"
	AWSOperation Key = "aws.operation"
	AWSRegion    Key = "aws.region"

	RiidoTraceSurface         Key = "riido.trace.surface"
	RiidoStoreOperation       Key = "riido.store.operation"
	RiidoTaskContextOperation Key = "riido.task_context.operation"
	RiidoPollAction           Key = "riido.poll.action"
)

func (key Key) String() string {
	return string(key)
}

func All() []Key {
	return []Key{
		WorkspaceID, Workspace, TaskID, RunID,
		AgentName, AgentIdentity, Workflow,
		WorkdirRoot, Workdir, OutputDir, LogsDir, ArtifactsDir,
		NativeConfigDir, NativeConfigHome, IRDir, NativeConfigVersion,
		RequiredSurfaces, AllowExperimentalRuntime,
		RuntimeLeaseID, RuntimeFencingToken, RuntimeCapabilityFingerprint,
		ProgressMessageCode, ProgressMessageKey, ProgressMessageArgPrefix, ThreadProgressSeq,
		HTTPRequestMethod, HTTPRoute, HTTPResponseStatusCode, HTTPStatusCode,
		AWSService, AWSOperation, AWSRegion,
		RiidoTraceSurface, RiidoStoreOperation, RiidoTaskContextOperation, RiidoPollAction,
	}
}
