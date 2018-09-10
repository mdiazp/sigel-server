package models

type ReservationServices interface {
	CreateReservation(r Reservation) (Reservation, error)
	GetReservationById(id int) (Reservation, error)
	UpdateReservation(r Reservation) (Reservation, error)
	DeleteReservation(id int) error
	GetReservationQuerySeter() ReservationQuerySeter
}
