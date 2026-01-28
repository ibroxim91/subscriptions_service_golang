package models

import "time"

// Subscription represents a subscription object
type Subscription struct {
    ID          uint       `json:"id" example:"1"`
    CreatedAt   time.Time  `json:"created_at" example:"2026-01-28T15:04:05Z"`
    UpdatedAt   time.Time  `json:"updated_at" example:"2026-01-28T15:04:05Z"`
    DeletedAt   *time.Time `json:"deleted_at,omitempty"`
    ServiceName string     `json:"service_name" example:"Netflix"`
    Price       int        `json:"price" example:"4500"`
    UserID      string     `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
    StartDate   time.Time  `json:"start_date" example:"2026-01-28"`
    EndDate     *time.Time `json:"end_date,omitempty" example:"2026-06-28"`
}
