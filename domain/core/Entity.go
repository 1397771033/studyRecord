package core

import (
	"github.com/satori/go.uuid"
)

type Entity struct {
	Id uuid.UUID
}

// GenerateId 生成一个id
func (entity *Entity) GenerateId() {
	entity.Id = uuid.NewV4()
}
