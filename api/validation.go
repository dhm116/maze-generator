package api

type validationError struct {
	ActualTag string `json:"tag"`
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	Param     string `json:"param"`
}
