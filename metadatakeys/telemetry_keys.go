package metadatakeys

const (
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
