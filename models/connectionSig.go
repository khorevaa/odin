package models

import uuid "github.com/satori/go.uuid"

type ConnectionSig struct {
	ClusterID  uuid.UUID `json:"cluster_id" example:"6d6958e1-a96c-4999-a995-698a0298161e"`
	InfobaseID uuid.UUID `json:"infobase_id" example:"6d6958e1-a96c-4999-a995-698a0298161e"`
	Process    uuid.UUID `json:"process" example:"6d6958e1-a96c-4999-a995-698a0298161e"`
	UUID       uuid.UUID `json:"uuid" example:"6d6958e1-a96c-4999-a995-698a0298161e"`
}

type TerminateConnectionsRequest struct {
	InfobaseID  string          `json:"infobase_id" example:"6d6958e1-a96c-4999-a995-698a0298161e or testib2"`
	Connections []ConnectionSig `json:"connections"`
}

type TerminateConnectionSig struct {
	ConnectionSig
	Terminated bool   `json:"terminated" example:"false"`
	Err        string `json:"err,omitempty" example:"error terminate connection"`
}

type TerminateConnectionsResponse struct {
	Count       int                      `json:"count" example:"0"`
	Connections []TerminateConnectionSig `json:"connections,omitempty"`
}

func (r *TerminateConnectionsResponse) AddResult(sig ConnectionSig, err error) {

	msg := ""
	terminated := true

	if err != nil {
		terminated = false
		msg = err.Error()
	}

	r.Connections = append(r.Connections, TerminateConnectionSig{
		ConnectionSig: sig,
		Terminated:    terminated,
		Err:           msg,
	})

	r.Count++
}
