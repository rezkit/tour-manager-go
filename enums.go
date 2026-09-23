package tourmanager

// SortOrder is a sort direction, used for the "order" query parameter on
// list endpoints.
type SortOrder string

const (
	SortAsc  SortOrder = "asc"
	SortDesc SortOrder = "desc"
)

// Valid reports whether o is a documented SortOrder value.
func (o SortOrder) Valid() bool { return o == SortAsc || o == SortDesc }

// OrderingCommand is a reorder/move instruction, sent as the "ordering"
// field on some update requests.
type OrderingCommand string

const (
	OrderUp    OrderingCommand = "up"
	OrderDown  OrderingCommand = "down"
	OrderFirst OrderingCommand = "first"
	OrderLast  OrderingCommand = "last"
)

// Valid reports whether c is a documented OrderingCommand value.
func (c OrderingCommand) Valid() bool {
	switch c {
	case OrderUp, OrderDown, OrderFirst, OrderLast:
		return true
	default:
		return false
	}
}
