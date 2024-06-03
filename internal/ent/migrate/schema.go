// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StadiaColumns holds the columns for the "stadia" table.
	StadiaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "team", Type: field.TypeString},
		{Name: "fdcouk", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "stadium", Type: field.TypeString, Unique: true},
		{Name: "capacity", Type: field.TypeInt},
		{Name: "latitude", Type: field.TypeFloat64},
		{Name: "longitude", Type: field.TypeFloat64},
		{Name: "country", Type: field.TypeString},
	}
	// StadiaTable holds the schema information for the "stadia" table.
	StadiaTable = &schema.Table{
		Name:       "stadia",
		Columns:    StadiaColumns,
		PrimaryKey: []*schema.Column{StadiaColumns[0]},
	}
	// WeathersColumns holds the columns for the "weathers" table.
	WeathersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "temperature", Type: field.TypeFloat64},
		{Name: "description", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString, Default: "01d"},
		{Name: "stadium_weather", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// WeathersTable holds the schema information for the "weathers" table.
	WeathersTable = &schema.Table{
		Name:       "weathers",
		Columns:    WeathersColumns,
		PrimaryKey: []*schema.Column{WeathersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "weathers_stadia_weather",
				Columns:    []*schema.Column{WeathersColumns[6]},
				RefColumns: []*schema.Column{StadiaColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StadiaTable,
		WeathersTable,
	}
)

func init() {
	WeathersTable.ForeignKeys[0].RefTable = StadiaTable
}
