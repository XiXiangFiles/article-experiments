package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Collection interface {
	InitCollection(ctx context.Context) error
	AddSingleData(ctx context.Context, raw interface{}) error
	Query(ctx context.Context, filter interface{}, dest interface{}) error
}

type sortType int

var (
	Asc  sortType = 0
	Desc sortType = 1
)

func (s sortType) ToString() string {
	if s == Desc {
		return "DESC"
	}
	return "ASC"
}

type GUIDID struct {
	Id string `gorm:"column:id;primaryKey;type:char(36);" json:"id"`
}

type Created struct {
	CreatedAt int64 `gorm:"column:created_at;not null;index;type:BIGINT UNSIGNED;autoCreateTime;<-:create" bson:"created_at" json:"created_at"`
}

type Updated struct {
	UpdatedAt int64 `gorm:"column:updated_at;not null;index;type:BIGINT UNSIGNED;autoUpdateTime" bson:"updated_at" json:"updated_at"`
}

type Deleted struct {
	DeletedAt int64 `gorm:"column:deleted_at;not null;index;type:BIGINT UNSIGNED;default:0" bson:"deleted_at" json:"deleted_at"`
}

func (t *Deleted) SetDeletedTime() {
	t.DeletedAt = time.Now().Unix()
}

func (t *Deleted) ClearDeletedTime() {
	t.DeletedAt = 0
}

func (t *Updated) SetUpdatedTime() {
	t.UpdatedAt = time.Now().Unix()
}

func (t *Created) SetCreatedTime() {
	t.CreatedAt = time.Now().Unix()
}

func (u *GUIDID) BeforeCreate(tx *gorm.DB) (err error) {
	if len(u.Id) == 0 {
		u.Id = uuid.NewString()
	}

	return nil
}
