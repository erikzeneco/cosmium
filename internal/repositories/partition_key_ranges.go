package repositories

import (
	"fmt"

	"github.com/google/uuid"
	repositorymodels "github.com/pikami/cosmium/internal/repository_models"
	"github.com/pikami/cosmium/internal/resourceid"
)

// I have no idea what this is tbh
func GetPartitionKeyRanges(databaseId string, collectionId string) ([]repositorymodels.PartitionKeyRange, repositorymodels.RepositoryStatus) {
	databaseRid := databaseId
	collectionRid := collectionId
	var timestamp int64 = 0

	if database, ok := storeState.Databases[databaseId]; !ok {
		databaseRid = database.ResourceID
	}

	if collection, ok := storeState.Collections[databaseId][collectionId]; !ok {
		collectionRid = collection.ResourceID
		timestamp = collection.TimeStamp
	}

	pkrResourceId := resourceid.NewCombined(databaseRid, collectionRid, resourceid.New())
	pkrSelf := fmt.Sprintf("dbs/%s/colls/%s/pkranges/%s/", databaseRid, collectionRid, pkrResourceId)
	etag := fmt.Sprintf("\"%s\"", uuid.New())

	return []repositorymodels.PartitionKeyRange{
		{
			ResourceID:         pkrResourceId,
			ID:                 "0",
			Etag:               etag,
			MinInclusive:       "",
			MaxExclusive:       "FF",
			RidPrefix:          0,
			Self:               pkrSelf,
			ThroughputFraction: 1,
			Status:             "online",
			Parents:            []interface{}{},
			TimeStamp:          timestamp,
			Lsn:                17,
		},
	}, repositorymodels.StatusOk
}
