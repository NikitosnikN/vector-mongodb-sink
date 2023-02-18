package app

import (
	"context"
	"log"
	"strings"
	"time"
	"vector-mongodb-sink/adapters"
	"vector-mongodb-sink/internal/common"
)

func StorePlainText(b []byte, collection string) {
	now := time.Now()
	var dataToInsert []interface{}

	for _, m := range strings.Split(string(b), "\n") {
		dataToInsert = append(dataToInsert, common.LogRecord{Timestamp: now, Message: m})
	}

	coll := adapters.GetCollection(adapters.DB, collection)

	if len(dataToInsert) > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		_, err := coll.InsertMany(ctx, dataToInsert)

		if err != nil {
			log.Fatal(err)
		}
	}
}
