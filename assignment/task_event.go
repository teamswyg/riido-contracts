package assignment

import "time"

type TaskEvent struct {
	Seq          int64             `json:"seq"`
	TaskID       string            `json:"task_id"`
	AssignmentID string            `json:"assignment_id"`
	AgentID      string            `json:"agent_id"`
	Type         string            `json:"type"`
	State        AssignmentState   `json:"state,omitempty"`
	Message      string            `json:"message,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	At           time.Time         `json:"at"`
}
