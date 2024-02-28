package model

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `tiny_url` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `long_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `short_url` char(7) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT NULL,
  `expired_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `short_rl` (`short_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

JSON Sample
-------------------------------------
{    "id": 1,    "long_url": "DAwToHnljlBkphWqEsjvMWySa",    "short_url": "VLuOntrrIREgIZMSqmLsUAPVg",    "created_at": "2166-02-08T00:37:28.667322062+08:00",    "expired_at": "2036-03-13T10:21:53.575105042+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// TinyURL struct is a row record of the tiny_url table in the dev database
type TinyURL struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] long_url                                       text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	LongURL string `gorm:"column:long_url;type:text;size:65535;" json:"long_url"`
	//[ 2] short_url                                      char(7)              null: false  primary: false  isArray: false  auto: false  col: char            len: 7       default: []
	ShortURL string `gorm:"column:short_url;type:char;size:7;" json:"short_url"`
	//[ 3] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [NULL]
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 4] expired_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [NULL]
	ExpiredAt time.Time `gorm:"column:expired_at;type:datetime;" json:"expired_at"`
}

var tiny_urlTableInfo = &TableInfo{
	Name: "tiny_url",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "long_url",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "LongURL",
			GoFieldType:        "string",
			JSONFieldName:      "long_url",
			ProtobufFieldName:  "long_url",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "short_url",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(7)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       7,
			GoFieldName:        "ShortURL",
			GoFieldType:        "string",
			JSONFieldName:      "short_url",
			ProtobufFieldName:  "short_url",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "expired_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ExpiredAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "expired_at",
			ProtobufFieldName:  "expired_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TinyURL) TableName() string {
	return "tiny_url"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TinyURL) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TinyURL) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TinyURL) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TinyURL) TableInfo() *TableInfo {
	return tiny_urlTableInfo
}
