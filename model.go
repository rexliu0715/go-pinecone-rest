package pinecone

import "github.com/go-resty/resty/v2"

type Client struct {
	Index       string
	Environment string
	APIKey      string
	Debug       bool
	Resty       *resty.Client
}

type DescribeIndexStatsRequest struct {
	Filter *map[string]interface{} `json:"filter,omitempty"`
}

type QueryRequest struct {
	Namespace       *string                 `json:"namespace,omitempty"`
	Filter          *map[string]interface{} `json:"filter,omitempty"`
	TopK            int64                   `json:"topK,omitempty"`
	IncludeValues   *bool                   `json:"includeValues,omitempty"`
	IncludeMetadata *bool                   `json:"includeMetadata,omitempty"`
	Vector          *[]float32              `json:"vector,omitempty"`
	SparseVector    *[]SparseVector         `json:"sparseVector,omitempty"`
	ID              *string                 `json:"id,omitempty"`
}

type UpsertRequest struct {
	Namespace *string  `json:"namespace,omitempty"`
	Vectors   []Vector `json:"vectors,omitempty"`
}

type UpdateRequest struct {
	ID           string                  `json:"id,omitempty"`
	Score        *float32                `json:"score,omitempty"`
	Values       *[]float32              `json:"values,omitempty"`
	SparseValues *[]SparseVector         `json:"sparseValues,omitempty"`
	SetMetadata  *map[string]interface{} `json:"setMetadata,omitempty"`
	Namespace    *string                 `json:"namespace,omitempty"`
}

type DeleteRequest struct {
	IDs       []string                `json:"ids,omitempty"`
	Namespace *string                 `json:"namespace,omitempty"`
	DeleteAll *bool                   `json:"deleteAll,omitempty"`
	Filter    *map[string]interface{} `json:"filter,omitempty"`
}

type FetchRequest struct {
	IDs       []string `json:"ids,omitempty"`
	Namespace *string  `json:"namespace,omitempty"`
}

type DescribeIndexStatsResponse struct {
	Namespace        *map[string]string `json:"namespace,omitempty"`
	Dimension        *int64             `json:"dimension,omitempty"`
	IndexFullness    *float32           `json:"indexFullness,omitempty"`
	TotalVectorCount *int64             `json:"totalVectorCount,omitempty"`
}

type QueryResponse struct {
	Matches   []Vector `json:"matches,omitempty"`
	Namespace *string  `json:"namespace,omitempty"`
}

type DeleteResponse struct {
	Result map[string]interface{}
}

type FetchResponse struct {
	Vectors   map[string]Vector `json:"vectors,omitempty"`
	Namespace *string           `json:"namespace,omitempty"`
}

type UpdateResponse struct {
	Result *map[string]interface{}
}

type UpsertResponse struct {
	UpsertedCount int64 `json:"upsertedCount,omitempty"`
}

type SparseVector struct {
	Indices []int64   `json:"indices,omitempty"`
	Values  []float32 `json:"values,omitempty"`
}

type Vector struct {
	ID           string                  `json:"id,omitempty"`
	Score        *float32                `json:"score,omitempty"`
	Values       *[]float32              `json:"values,omitempty"`
	SparseValues *[]SparseVector         `json:"sparseValues,omitempty"`
	Metadata     *map[string]interface{} `json:"metadata,omitempty"`
}

type ErrorResponse struct {
	Code    int32    `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Details []Detail `json:"details,omitempty"`
}

type Detail struct {
	TypeUrl string `json:"typeUrl,omitempty"`
	Value   string `json:"value,omitempty"`
}
