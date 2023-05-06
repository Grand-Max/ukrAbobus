package models

import "time"

type User struct {
	UserID    uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsAdmin   bool
	Document  *Document `json:"-" gorm:"foreignKey:DocumentID"`
}

type Document struct {
	DocumentID  uint `gorm:"primaryKey"`
	Name        string
	Description string
	Discount    float64
}

type Ticket struct {
	TicketsID uint   `gorm:"primaryKey"`
	User      []User `json:"-" gorm:"foreignKey:UserID"`
	Trip      []Trip `json:"-" gorm:"foreignKey:TripsID"`
	Price     float64
}

type Trip struct {
	TripsID        uint `gorm:"primaryKey"`
	CityFrom       string
	CityTo         string
	ArrivalTime    time.Time
	DepartureTime  time.Time
	NumberOfPlaces uint
	TransportType  string
}
