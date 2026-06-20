package main

type fixtureSummary struct {
	ContractID         string `json:"contract_id"`
	Context            string `json:"context"`
	ServiceSchema      string `json:"service_schema"`
	OperationCount     int    `json:"operation_count"`
	ComponentCount     int    `json:"component_count"`
	EnumCount          int    `json:"enum_count"`
	SumTypeCount       int    `json:"sum_type_count"`
	GeneratedPathCount int    `json:"generated_path_count"`
	V2OperationCount   int    `json:"v2_operation_count"`
}
