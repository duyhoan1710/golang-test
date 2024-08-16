package enum

type EOrderState int

const (
	Created   EOrderState = 0
	Confirmed EOrderState = 1
	Delivered EOrderState = 2
	Cancelled EOrderState = 3
)
