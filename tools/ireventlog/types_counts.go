package main

type nativeConfigCounts struct {
	Forbidden      int `json:"forbidden"`
	PreExecute     int `json:"pre_execute"`
	Required       int `json:"required"`
	PhaseDependent int `json:"phase_dependent"`
}
