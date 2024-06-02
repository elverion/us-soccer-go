// Code generated by ent, DO NOT EDIT.

package stadium

import (
	"us-soccer-go-test/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Stadium {
	return predicate.Stadium(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Stadium {
	return predicate.Stadium(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Stadium {
	return predicate.Stadium(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Stadium {
	return predicate.Stadium(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Stadium {
	return predicate.Stadium(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Stadium {
	return predicate.Stadium(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Stadium {
	return predicate.Stadium(sql.FieldLTE(FieldID, id))
}

// Team applies equality check predicate on the "team" field. It's identical to TeamEQ.
func Team(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldTeam, v))
}

// Fdcouk applies equality check predicate on the "fdcouk" field. It's identical to FdcoukEQ.
func Fdcouk(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldFdcouk, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldCity, v))
}

// Stadium applies equality check predicate on the "stadium" field. It's identical to StadiumEQ.
func Stadium(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldStadium, v))
}

// Capacity applies equality check predicate on the "capacity" field. It's identical to CapacityEQ.
func Capacity(v int) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldCapacity, v))
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldLatitude, v))
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldLongitude, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldCountry, v))
}

// TeamEQ applies the EQ predicate on the "team" field.
func TeamEQ(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldTeam, v))
}

// TeamNEQ applies the NEQ predicate on the "team" field.
func TeamNEQ(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldNEQ(FieldTeam, v))
}

// TeamIn applies the In predicate on the "team" field.
func TeamIn(vs ...string) predicate.Stadium {
	return predicate.Stadium(sql.FieldIn(FieldTeam, vs...))
}

// TeamNotIn applies the NotIn predicate on the "team" field.
func TeamNotIn(vs ...string) predicate.Stadium {
	return predicate.Stadium(sql.FieldNotIn(FieldTeam, vs...))
}

// TeamGT applies the GT predicate on the "team" field.
func TeamGT(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldGT(FieldTeam, v))
}

// TeamGTE applies the GTE predicate on the "team" field.
func TeamGTE(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldGTE(FieldTeam, v))
}

// TeamLT applies the LT predicate on the "team" field.
func TeamLT(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldLT(FieldTeam, v))
}

// TeamLTE applies the LTE predicate on the "team" field.
func TeamLTE(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldLTE(FieldTeam, v))
}

// TeamContains applies the Contains predicate on the "team" field.
func TeamContains(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldContains(FieldTeam, v))
}

// TeamHasPrefix applies the HasPrefix predicate on the "team" field.
func TeamHasPrefix(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldHasPrefix(FieldTeam, v))
}

// TeamHasSuffix applies the HasSuffix predicate on the "team" field.
func TeamHasSuffix(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldHasSuffix(FieldTeam, v))
}

// TeamEqualFold applies the EqualFold predicate on the "team" field.
func TeamEqualFold(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEqualFold(FieldTeam, v))
}

// TeamContainsFold applies the ContainsFold predicate on the "team" field.
func TeamContainsFold(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldContainsFold(FieldTeam, v))
}

// FdcoukEQ applies the EQ predicate on the "fdcouk" field.
func FdcoukEQ(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldFdcouk, v))
}

// FdcoukNEQ applies the NEQ predicate on the "fdcouk" field.
func FdcoukNEQ(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldNEQ(FieldFdcouk, v))
}

// FdcoukIn applies the In predicate on the "fdcouk" field.
func FdcoukIn(vs ...string) predicate.Stadium {
	return predicate.Stadium(sql.FieldIn(FieldFdcouk, vs...))
}

// FdcoukNotIn applies the NotIn predicate on the "fdcouk" field.
func FdcoukNotIn(vs ...string) predicate.Stadium {
	return predicate.Stadium(sql.FieldNotIn(FieldFdcouk, vs...))
}

// FdcoukGT applies the GT predicate on the "fdcouk" field.
func FdcoukGT(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldGT(FieldFdcouk, v))
}

// FdcoukGTE applies the GTE predicate on the "fdcouk" field.
func FdcoukGTE(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldGTE(FieldFdcouk, v))
}

// FdcoukLT applies the LT predicate on the "fdcouk" field.
func FdcoukLT(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldLT(FieldFdcouk, v))
}

// FdcoukLTE applies the LTE predicate on the "fdcouk" field.
func FdcoukLTE(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldLTE(FieldFdcouk, v))
}

// FdcoukContains applies the Contains predicate on the "fdcouk" field.
func FdcoukContains(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldContains(FieldFdcouk, v))
}

// FdcoukHasPrefix applies the HasPrefix predicate on the "fdcouk" field.
func FdcoukHasPrefix(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldHasPrefix(FieldFdcouk, v))
}

// FdcoukHasSuffix applies the HasSuffix predicate on the "fdcouk" field.
func FdcoukHasSuffix(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldHasSuffix(FieldFdcouk, v))
}

// FdcoukEqualFold applies the EqualFold predicate on the "fdcouk" field.
func FdcoukEqualFold(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEqualFold(FieldFdcouk, v))
}

// FdcoukContainsFold applies the ContainsFold predicate on the "fdcouk" field.
func FdcoukContainsFold(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldContainsFold(FieldFdcouk, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Stadium {
	return predicate.Stadium(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Stadium {
	return predicate.Stadium(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldContainsFold(FieldCity, v))
}

// StadiumEQ applies the EQ predicate on the "stadium" field.
func StadiumEQ(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldStadium, v))
}

// StadiumNEQ applies the NEQ predicate on the "stadium" field.
func StadiumNEQ(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldNEQ(FieldStadium, v))
}

// StadiumIn applies the In predicate on the "stadium" field.
func StadiumIn(vs ...string) predicate.Stadium {
	return predicate.Stadium(sql.FieldIn(FieldStadium, vs...))
}

// StadiumNotIn applies the NotIn predicate on the "stadium" field.
func StadiumNotIn(vs ...string) predicate.Stadium {
	return predicate.Stadium(sql.FieldNotIn(FieldStadium, vs...))
}

// StadiumGT applies the GT predicate on the "stadium" field.
func StadiumGT(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldGT(FieldStadium, v))
}

// StadiumGTE applies the GTE predicate on the "stadium" field.
func StadiumGTE(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldGTE(FieldStadium, v))
}

// StadiumLT applies the LT predicate on the "stadium" field.
func StadiumLT(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldLT(FieldStadium, v))
}

// StadiumLTE applies the LTE predicate on the "stadium" field.
func StadiumLTE(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldLTE(FieldStadium, v))
}

// StadiumContains applies the Contains predicate on the "stadium" field.
func StadiumContains(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldContains(FieldStadium, v))
}

// StadiumHasPrefix applies the HasPrefix predicate on the "stadium" field.
func StadiumHasPrefix(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldHasPrefix(FieldStadium, v))
}

// StadiumHasSuffix applies the HasSuffix predicate on the "stadium" field.
func StadiumHasSuffix(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldHasSuffix(FieldStadium, v))
}

// StadiumEqualFold applies the EqualFold predicate on the "stadium" field.
func StadiumEqualFold(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEqualFold(FieldStadium, v))
}

// StadiumContainsFold applies the ContainsFold predicate on the "stadium" field.
func StadiumContainsFold(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldContainsFold(FieldStadium, v))
}

// CapacityEQ applies the EQ predicate on the "capacity" field.
func CapacityEQ(v int) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldCapacity, v))
}

// CapacityNEQ applies the NEQ predicate on the "capacity" field.
func CapacityNEQ(v int) predicate.Stadium {
	return predicate.Stadium(sql.FieldNEQ(FieldCapacity, v))
}

// CapacityIn applies the In predicate on the "capacity" field.
func CapacityIn(vs ...int) predicate.Stadium {
	return predicate.Stadium(sql.FieldIn(FieldCapacity, vs...))
}

// CapacityNotIn applies the NotIn predicate on the "capacity" field.
func CapacityNotIn(vs ...int) predicate.Stadium {
	return predicate.Stadium(sql.FieldNotIn(FieldCapacity, vs...))
}

// CapacityGT applies the GT predicate on the "capacity" field.
func CapacityGT(v int) predicate.Stadium {
	return predicate.Stadium(sql.FieldGT(FieldCapacity, v))
}

// CapacityGTE applies the GTE predicate on the "capacity" field.
func CapacityGTE(v int) predicate.Stadium {
	return predicate.Stadium(sql.FieldGTE(FieldCapacity, v))
}

// CapacityLT applies the LT predicate on the "capacity" field.
func CapacityLT(v int) predicate.Stadium {
	return predicate.Stadium(sql.FieldLT(FieldCapacity, v))
}

// CapacityLTE applies the LTE predicate on the "capacity" field.
func CapacityLTE(v int) predicate.Stadium {
	return predicate.Stadium(sql.FieldLTE(FieldCapacity, v))
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldLatitude, v))
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldNEQ(FieldLatitude, v))
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldIn(FieldLatitude, vs...))
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldNotIn(FieldLatitude, vs...))
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldGT(FieldLatitude, v))
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldGTE(FieldLatitude, v))
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldLT(FieldLatitude, v))
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldLTE(FieldLatitude, v))
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldLongitude, v))
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldNEQ(FieldLongitude, v))
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldIn(FieldLongitude, vs...))
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldNotIn(FieldLongitude, vs...))
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldGT(FieldLongitude, v))
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldGTE(FieldLongitude, v))
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldLT(FieldLongitude, v))
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Stadium {
	return predicate.Stadium(sql.FieldLTE(FieldLongitude, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Stadium {
	return predicate.Stadium(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Stadium {
	return predicate.Stadium(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Stadium {
	return predicate.Stadium(sql.FieldContainsFold(FieldCountry, v))
}

// HasWeather applies the HasEdge predicate on the "weather" edge.
func HasWeather() predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, WeatherTable, WeatherColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWeatherWith applies the HasEdge predicate on the "weather" edge with a given conditions (other predicates).
func HasWeatherWith(preds ...predicate.Weather) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		step := newWeatherStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Stadium) predicate.Stadium {
	return predicate.Stadium(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Stadium) predicate.Stadium {
	return predicate.Stadium(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Stadium) predicate.Stadium {
	return predicate.Stadium(sql.NotPredicates(p))
}
