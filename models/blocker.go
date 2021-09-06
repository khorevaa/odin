package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type InfobaseBlocker struct {
	Infobase  string `query:"infobase" json:"infobase" example:"testdb2"`
	ClusterID string `query:"cluster-id" json:"cluster_id" example:"80f7f2f6-2feb-46bf-92f4-19294a2f5dc7"`

	Message           string    `query:"message" json:"message" example:"Обновление информационной базы"`
	PermissionCode    string    `query:"permission-code" json:"permission_code" example:"123"`
	DeniedParameter   string    `query:"denied-parameter" json:"denied_parameter" example:"code"`
	ScheduledJobsDeny bool      `query:"scheduled-jobs-deny" json:"scheduled_jobs_deny"`
	SessionsDeny      bool      `query:"sessions-deny,required" json:"sessions_deny"`
	Reload            bool      `query:"reload" json:"reload"`
	DeniedFrom        time.Time `query:"denied-from" json:"denied_from" example:"2020-10-01T08:30:00Z"`
	DeniedTo          time.Time `query:"denied-to" json:"denied_to" example:"2020-10-01T08:30:00Z"`
}

func (b *InfobaseBlocker) Empty() bool {
	return !b.SessionsDeny && b.DeniedFrom.IsZero() && b.DeniedTo.IsZero() &&
		len(b.Message) == 0 && len(b.PermissionCode) == 0
}

type InfobaseUnblocker struct {
	Infobase string `query:"infobase-name" json:"infobase" example:"testdb2"`

	InfobaseID uuid.UUID `query:"infobase-id" json:"infobase_id" example:"80f7f2f6-2feb-46bf-92f4-19294a2f5dc7"`
	ClusterID  uuid.UUID `query:"cluster-id" json:"cluster_id" example:"80f7f2f6-2feb-46bf-92f4-19294a2f5dc7"`

	PermissionCode    string `query:"permission-code" json:"permission_code" example:""`
	DeniedParameter   string `query:"denied-parameter" json:"denied_parameter" example:""`
	ScheduledJobsDeny bool   `query:"scheduled-jobs-deny" json:"scheduled_jobs_deny"`
	SessionsDeny      bool   `query:"sessions-deny,required" json:"sessions_deny"`
}
