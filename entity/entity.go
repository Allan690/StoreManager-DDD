package entity

import (
	"fmt"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)


//ID entity ID
type ID struct {
	uuid.UUID
}


//NewID create a new entity ID
func NewID() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}

func (mu ID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bsontype.Binary, bsoncore.AppendBinary(nil, 4, mu.UUID[:]), nil
}

func (mu *ID) UnmarshalBSONValue(t bsontype.Type, raw []byte) error {
	if t != bsontype.Binary {
		return fmt.Errorf("invalid format on unmarshal bson value")
	}

	_, data, _, ok := bsoncore.ReadBinary(raw)
	if !ok {
		return fmt.Errorf("not enough bytes to unmarshal bson value")
	}

	copy(mu.UUID[:], data)

	return nil
}
