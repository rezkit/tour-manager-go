package tourmanager

// EntityType identifies a kind of polymorphic entity. It's used as a path
// parameter on sub-resource endpoints such as categories, fields and
// content, which can attach to more than one kind of parent resource.
//
// Named EntityType rather than Type because "type" is a Go keyword.
type EntityType string

const (
	EntityTypeAccommodation  EntityType = "accommodation"
	EntityTypeCake           EntityType = "cake"
	EntityTypeCategory       EntityType = "category"
	EntityTypeDeparture      EntityType = "departure"
	EntityTypeElement        EntityType = "element"
	EntityTypeElementOption  EntityType = "element_option"
	EntityTypeExtra          EntityType = "extra"
	EntityTypeHoliday        EntityType = "holiday"
	EntityTypeHolidayVersion EntityType = "holiday_version"
	EntityTypeImage          EntityType = "image"
	EntityTypeItineraryEntry EntityType = "itinerary_entry"
	EntityTypeLocation       EntityType = "location"
	EntityTypeMap            EntityType = "map"
	EntityTypeRoomType       EntityType = "room_type"
)

// String returns t as a plain string.
func (t EntityType) String() string { return string(t) }

// Valid reports whether t is one of the EntityType values documented by
// the API.
func (t EntityType) Valid() bool {
	switch t {
	case EntityTypeAccommodation, EntityTypeCake, EntityTypeCategory, EntityTypeDeparture,
		EntityTypeElement, EntityTypeElementOption, EntityTypeExtra, EntityTypeHoliday,
		EntityTypeHolidayVersion, EntityTypeImage, EntityTypeItineraryEntry, EntityTypeLocation,
		EntityTypeMap, EntityTypeRoomType:
		return true
	default:
		return false
	}
}
