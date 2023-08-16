package main

type DescribeIndexStatsRequest struct {
	Filter *map[string]interface{} `json:"filter,omitempty"`
}

type QueryRequest struct {
	Namespace *string `json:"namespace,omitempty"`
	Filter *map[string]interface{} `json:"filter,omitempty"`
	TopK int64 `json:"top_k,omitempty"`
	IncludeValues *bool `json:"include_values,omitempty"`
	IncludeMetadata *bool `json:"include_metadata,omitempty"`
	Vector *[]float32 `json:"vector,omitempty"`
	SparseVector *[]SparseVector	`json:"sparse_vector,omitempty"`
	ID *string `json:"id,omitempty"`
}

type UpsertRequest struct {
	Namespace *string `json:"namespace,omitempty"`
	Vectors []Vector `json:"vectors,omitempty"`
}

type UpdateRequest struct {
	ID string `json:"id,omitempty"`
	Score *float32 `json:"score,omitempty"`
	Values *[]float32 `json:"values,omitempty"`
	SparseValues *[]SparseVector `json:"sparse_values,omitempty"`
	SetMetadata *map[string]string `json:"set_metadata,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

type DeleteRequest struct {
	IDs []string `json:"ids,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	DeleteAll *bool `json:"delete_all,omitempty"`
	Filter *map[string]interface{} `json:"filter,omitempty"`
}

type FetchRequest struct {
	IDs []string `json:"ids,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

type DescribeIndexStatsResponse struct {
	Namespace *map[string]string `json:"namespace,omitempty"`
	Dimension *int64 `json:"dimension,omitempty"`
	IndexFullness *float32 `json:"index_fullness,omitempty"`
	TotalVectorCount *int64 `json:"total_vector_count,omitempty"`
}

type QueryResponse struct {
	Matches []Vector `json:"matches,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

type DeleteResponse struct{
	Result map[string]string
} 

type FetchResponse struct {
	Vectors []Vector `json:"vectors,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

type UpdateResponse struct {
	Result *map[string]string
}

type UpsertResponse struct {
	UpsertedCount int64 `json:"upserted_count,omitempty"`
}

type SparseVector struct {
	Indices []int64 `json:"indices,omitempty"`
	Values []float32 `json:"values,omitempty"`
}

type Vector struct {
	ID string `json:"id,omitempty"`
	Score *float32 `json:"score,omitempty"`
	Values *[]float32 `json:"values,omitempty"`
	SparseValues *[]SparseVector `json:"sparse_values,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
}
