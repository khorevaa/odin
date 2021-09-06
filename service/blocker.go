package service

import (
	errors2 "github.com/khorevaa/odin/errors"
	"github.com/khorevaa/odin/models"
	"github.com/khorevaa/ras-client/serialize"
	uuid "github.com/satori/go.uuid"
	"time"
)

func (s *service) Block(client ClientContext, blocker *models.InfobaseBlocker) (*models.InfobaseUnblocker, error) {

	infobase := blocker.Infobase

	if len(infobase) == 0 {
		infobaseName, ok := client.GetInfobaseID()
		if !ok {
			return nil, errors2.BadRequest.New("infobase id or name must be set")
		}
		infobase = infobaseName.String()
	}

	infobaseSummaryInfo, err := s.findInfobase(client, infobase)
	if err != nil {
		return nil, err
	}

	infobaseInfo, err := client.GetInfobaseInfo(infobaseSummaryInfo.ClusterID, infobaseSummaryInfo.UUID)

	if err != nil {
		return nil, err
	}

	unblock := &models.InfobaseUnblocker{
		ClusterID:         infobaseInfo.ClusterID,
		InfobaseID:        infobaseInfo.UUID,
		PermissionCode:    infobaseInfo.PermissionCode,
		SessionsDeny:      infobaseInfo.SessionsDeny,
		ScheduledJobsDeny: infobaseInfo.ScheduledJobsDeny,
	}

	block := infobaseInfo.Blocker(blocker.Reload)
	block.Msg(blocker.Message).
		To(blocker.DeniedTo).
		From(blocker.DeniedFrom).
		ScheduledJobs(blocker.ScheduledJobsDeny).
		Code(blocker.PermissionCode)

	client.AddAuth(infobaseInfo.ClusterID, infobaseInfo.UUID)
	err = block.Block(client.Context(), client.GetApiClient())
	return unblock, err
}

func (s *service) Unblock(client ClientContext, unblocker *models.InfobaseUnblocker) (*serialize.InfobaseSummaryInfo, error) {

	infobase := unblocker.Infobase

	if len(infobase) == 0 && unblocker.InfobaseID == uuid.Nil {
		infobaseName, ok := client.GetInfobaseID()
		if !ok {
			return nil, errors2.BadRequest.New("infobase id or name must be set")
		}
		infobase = infobaseName.String()
	} else {
		infobase = unblocker.InfobaseID.String()
	}

	infobaseSummaryInfo, err := s.findInfobase(client, infobase)
	if err != nil {
		return nil, err
	}

	infobaseInfo, err := client.GetInfobaseInfo(infobaseSummaryInfo.ClusterID, infobaseSummaryInfo.UUID)

	if err != nil {
		return nil, err
	}

	infobaseInfo.ScheduledJobsDeny = unblocker.ScheduledJobsDeny
	infobaseInfo.PermissionCode = unblocker.PermissionCode
	infobaseInfo.SessionsDeny = unblocker.SessionsDeny
	infobaseInfo.DeniedFrom = time.Time{}
	infobaseInfo.DeniedTo = time.Time{}

	err = client.UpdateInfobase(infobaseInfo.ClusterID, infobaseInfo)
	if err != nil {
		return nil, err
	}
	summary := infobaseInfo.Summary()
	return &summary, nil

}
