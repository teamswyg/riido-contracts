package main

type manifestLoopReport struct {
	Complete       int
	Missing        int
	MissingGroups  []manifestGroupCount
	MissingSamples []manifestGroupSample
}
