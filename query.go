package tourmanager

import (
	"net/url"
	"strconv"
)

// setInt sets key to v's decimal string in q, unless v is zero (the
// project-wide "unset" sentinel for page/limit and similar int params —
// see the ListOptions field-type rule).
func setInt(q url.Values, key string, v int) {
	if v != 0 {
		q.Set(key, strconv.Itoa(v))
	}
}

// setString sets key to v in q, unless v is empty (the "unset" sentinel
// for free-text/enum string params, which never have a legitimate empty
// value).
func setString(q url.Values, key, v string) {
	if v != "" {
		q.Set(key, v)
	}
}

// setBool01 sets key to "1" or "0" in q if v is non-nil, matching the
// API's `integer enum [0, 1]` convention for tri-state filters (nil =
// don't filter, false and true are both meaningful values distinct from
// "unset").
func setBool01(q url.Values, key string, v *bool) {
	if v == nil {
		return
	}
	if *v {
		q.Set(key, "1")
	} else {
		q.Set(key, "0")
	}
}

// setBool sets key to "true" in q if v is true. Used for genuine boolean
// query flags (as opposed to the API's tri-state 0/1 filters) where
// omitting the parameter and sending false are equivalent, so no pointer
// is needed.
func setBool(q url.Values, key string, v bool) {
	if v {
		q.Set(key, "true")
	}
}

// setEnum sets key to string(v) in q, unless v is the empty string.
func setEnum[T ~string](q url.Values, key string, v T) {
	if v != "" {
		q.Set(key, string(v))
	}
}
