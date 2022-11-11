package entities

import (
	"context"
	"reflect"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

type ProductRaw struct {
	Id          int    `gorm:"column:p_id;primaryKey;index;not null;" json:"id"`
	Index       int    `gorm:"column:s_index;index;not null;" json:"index"`
	Name        string `gorm:"column:name;index;not null;" json:"name"`
	Grade       string `gorm:"column:grade;index;not null;" json:"grade"`
	ArticleLink string `gorm:"column:article_link;index;not null;" json:"article_link"`
	MarketDate  string `gorm:"column:market_date;index;not null;" json:"market_date"`
	Price       string `gorm:"column:price;index;not null;" json:"price"`
	IsLimit     int    `gorm:"column:is_limit;index;not null;" json:"is_limit"`
	IsDiscount  int    `gorm:"column:is_discount;index;not null;" json:"is_discount"`
	IsWithdraw  int    `gorm:"column:is_withdraw;index;not null;" json:"is_withdraw"`
	Info        string `gorm:"column:info;index;not null;" json:"info"`
}

type ProductFilter struct {
	Id *int `condition:"p_id = ?" json:"id"`
}

type ProductCollection struct {
	db *gorm.DB
}

func (raw ProductRaw) TableName() string {
	return "products"
}

func NewProductCollection(db *gorm.DB) Collection {
	var c Collection = &ProductCollection{db: db}
	return c
}

func (p *ProductCollection) InitCollection(ctx context.Context) error {
	return p.db.AutoMigrate(&ProductRaw{})
}

func (g *ProductCollection) SetQueryParams(sql *gorm.DB, filter *ProductFilter) {
	filterStruct := structs.New(filter)
	for _, fieldName := range filterStruct.Fields() {
		val := filterStruct.Field(fieldName.Name()).Value()
		type_ := reflect.TypeOf(val)
		condition := filterStruct.Field(fieldName.Name()).Tag("condition")
		if type_ == reflect.TypeOf(filter.Id) {
			if valString := reflect.ValueOf(val).Interface().(*int); valString != nil {
				sql.Where(condition, *valString)
			}
		}
	}
}

func (p *ProductCollection) AddSingleData(ctx context.Context, raw interface{}) error {
	raw_ := reflect.ValueOf(raw).Interface().(*ProductRaw)
	return p.db.Model(&ProductRaw{}).Create(raw_).Error
}

func (p *ProductCollection) Query(ctx context.Context, filter interface{}, dest interface{}) error {
	filter_ := reflect.ValueOf(filter).Interface().(*ProductFilter)
	sql := p.db.Model(&ProductRaw{})
	p.SetQueryParams(sql, filter_)

	err := sql.Find(dest).Error
	return err
}
