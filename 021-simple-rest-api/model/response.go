package model

// ResponseFormat is the generic format of API Response
type ResponseFormat struct {
	Data   interface{}  `json:"data,omitempty"`
	Meta   MetaResponse `json:"meta,omitempty"`
	Errors []Error      `json:"errors,omitempty"`
}

// MetaResponse is the metadata of API Response
type MetaResponse map[string]interface{}

// Error is the error format of API Response
type Error struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
