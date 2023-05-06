package models

import "time"

type User struct {
	UserID     uint      `gorm:"primaryKey"`
	FirstName  string    `gorm:"not null"`
	LastName   string    `gorm:"not null"`
	Email      string    `gorm:"not null"`
	Password   string    `gorm:"not null"`
	IsAdmin    bool      `gorm:"not null"`
	DocumentID int       `gorm:"index"`
	Document   *Document `gorm:"foreignKey:DocumentID"`
}

type Document struct {
	DocumentID  uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Discount    float64 `gorm:"not null"`
}

type Ticket struct {
	TicketsID uint    `gorm:"primaryKey"`
	UserID    int     `gorm:"index; not null;"`
	User      *User   `json:"-" gorm:"foreignKey:UserID"`
	TripID    int     `gorm:"index;not null;"`
	Trip      *Trip   `gorm:"foreignKey:TripsID"`
	Price     float64 `gorm:"not null"`
}

type Trip struct {
	TripsID        uint      `gorm:"primaryKey"`
	CityFrom       string    `gorm:"not null"`
	CityTo         string    `gorm:"not null"`
	ArrivalTime    time.Time `gorm:"not null"`
	DepartureTime  time.Time `gorm:"not null"`
	NumberOfPlaces uint      `gorm:"not null"`
	TransportType  string    `gorm:"not null"`
}
