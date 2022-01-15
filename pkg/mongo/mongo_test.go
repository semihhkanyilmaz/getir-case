package mongo

import (
	"context"
	"testing"
)

func Test_NewMongoClient(t *testing.T) {
	client := MewMongoClient("mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/?retryWrites=true")
	defer client.Disconnect(context.Background())

	if client == nil {
		t.Error("Mongo connection failed")
	}

}
