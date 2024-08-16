package enum

type EOrderState string

const (
	Created   EOrderState = "created"
	Confirmed EOrderState = "confirmed"
	Delivered EOrderState = "delivered"
	Cancelled EOrderState = "canceled"
)
