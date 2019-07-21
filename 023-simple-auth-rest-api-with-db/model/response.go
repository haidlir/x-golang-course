package model

import (
	"encoding/json"
)

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

// NewResponseFormat returns new response format
func NewResponseFormat() *ResponseFormat {
	newResp := &ResponseFormat{}
	newResp.Meta = MetaResponse{}
	newResp.Errors = []Error{}
	return newResp
}

// SetData sets response data
func (resp *ResponseFormat) SetData(data interface{}) {
	resp.Data = data
}

// AddMeta add response meta
func (resp *ResponseFormat) AddMeta(key string, val interface{}) {
	resp.Meta[key] = val
}

// AddError add response meta
func (resp *ResponseFormat) AddError(title, detail string) {
	newError := Error{
		Title:  title,
		Detail: detail,
	}
	resp.Errors = append(resp.Errors, newError)
}

// EncodeToJSON marshals response to json
func (resp *ResponseFormat) EncodeToJSON() ([]byte, error) {
	encodedResp, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return encodedResp, nil
}
