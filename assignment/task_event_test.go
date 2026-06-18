package assignment

func testTaskEvent() TaskEvent {
	return TaskEvent{
		Seq:          1,
		TaskID:       "task-a",
		AssignmentID: "asn-000001",
		AgentID:      "agent-a",
		Type:         EventAssignmentRunning,
		State:        AssignmentRunning,
		Message:      "running",
		Metadata:     map[string]string{"step": "run"},
		At:           testAssignmentTime(),
	}
}
