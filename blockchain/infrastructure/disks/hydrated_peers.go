package disks

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type hydratedPeers struct {
	ID           string          `json:"id" hydro:"0"`
	SyncInterval int64           `json:"sync_interval" hydro:"1"`
	List         []*hydratedPeer `json:"list" hydro:"2"`
	LastSyncTime string          `json:"last_sync_time" hydro:"3"`
}

func peersOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "ID" {
		if id, ok := ins.(*uuid.UUID); ok {
			return id.String(), nil
		}
	}

	if fieldName == "SyncInterval" {
		if syncInterval, ok := ins.(time.Duration); ok {
			return syncInterval.Nanoseconds(), nil
		}
	}

	if fieldName == "LastSyncTime" {
		if lastSyncTime, ok := ins.(*time.Time); ok {
			return lastSyncTime.Format(timeLayout), nil
		}
	}

	return nil, nil
}

func peersOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "ID" {
		id, err := uuid.FromString(ins.(string))
		if err != nil {
			return nil, err
		}

		return &id, nil
	}

	if fieldName == "SyncInterval" {
		return time.Duration(ins.(int64)), nil
	}

	if fieldName == "LastSyncTime" {
		if str, ok := ins.(string); ok {
			if str == "" {
				return nil, nil
			}

			lastSyncTime, err := time.Parse(timeLayout, str)
			if err != nil {
				return nil, err
			}

			return &lastSyncTime, nil
		}
	}

	return nil, nil
}
