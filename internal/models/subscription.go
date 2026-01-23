package models

import (
    "time"
    "gorm.io/gorm"
)

type Subscription struct {
    gorm.Model
    ServiceName string    `gorm:"type:varchar(255);not null" json:"service_name"`
    Price       int       `gorm:"not null;check:price >= 0" json:"price"`
    UserID      string    `gorm:"type:uuid;not null" json:"user_id"`
    StartDate   time.Time `gorm:"not null" json:"start_date"`
    EndDate     *time.Time `json:"end_date,omitempty"`
}
