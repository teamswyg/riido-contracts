package contextcli

// EnvelopeSchemaVersion is the version of the JSON envelope written to stdout.
const EnvelopeSchemaVersion = "1"

// Envelope is the single JSON object a work-context command writes to stdout
// when --json is set. Slices are never null so agents can iterate safely.
type Envelope struct {
	SchemaVersion string    `json:"schema_version"`
	RequestID     string    `json:"request_id"`
	Data          any       `json:"data"`
	Sources       []Source  `json:"sources"`
	Warnings      []Warning `json:"warnings"`
	Page          *Page     `json:"page"`
	Error         *Error    `json:"error"`
}

// SourceStatus reports whether a source could be read for this request.
type SourceStatus string

const (
	SourceAvailable   SourceStatus = "available"
	SourcePartial     SourceStatus = "partial"
	SourceDenied      SourceStatus = "denied"
	SourceUnavailable SourceStatus = "unavailable"
	SourceUnsupported SourceStatus = "unsupported"
	SourceDeleted     SourceStatus = "deleted"
)

// Source identifies one piece of Riido data that contributed to a result.
type Source struct {
	ID        string       `json:"id"`
	Kind      string       `json:"kind"`
	URL       string       `json:"url,omitempty"`
	Version   string       `json:"version,omitempty"`
	FetchedAt string       `json:"fetched_at,omitempty"`
	Status    SourceStatus `json:"status"`
	Truncated bool         `json:"truncated"`
	Reason    string       `json:"reason,omitempty"`
}

// Warning is a non-fatal note attached to a result.
type Warning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Page describes list truncation and how to fetch the next page.
type Page struct {
	Limit      int    `json:"limit"`
	NextCursor string `json:"next_cursor,omitempty"`
	Truncated  bool   `json:"truncated"`
}

// Error is the structured failure carried in the envelope. It never contains
// tokens or response headers.
type Error struct {
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	Retryable  bool      `json:"retryable"`
	NextAction string    `json:"next_action,omitempty"`
	RequestID  string    `json:"request_id"`
}
