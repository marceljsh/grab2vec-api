package models

import (
	"fmt"
	"time"
)

type Trajectory struct {
	TrjID       int       `json:"trj_id" gorm:"primaryKey"`
	DrivingMode string    `json:"driving_mode" gorm:"size:16;not null"`
	OsName      string    `json:"os_name" gorm:"size:8;not null"`
	StartedAt   time.Time `json:"started_at" gorm:"not null"`
	StartHour   int       `json:"start_hour"`
	EndedAt     time.Time `json:"ended_at" gorm:"not null"`
	EndHour     int       `json:"end_hour"`
	DayName     string    `json:"day_name" gorm:"size:10"`
	DayOfWeek   string    `json:"day_of_week" gorm:"size:8"`
	IsHoliday   bool      `json:"is_holiday" gorm:"type:boolean"`
	StartLat    float64   `json:"start_lat" gorm:"type:decimal(18,8);not null"`
	EndLat      float64   `json:"end_lat" gorm:"type:decimal(18,8);not null"`
	StartLon    float64   `json:"start_lon" gorm:"type:decimal(18,8);not null"`
	EndLon      float64   `json:"end_lon" gorm:"type:decimal(18,8);not null"`
	Distance    float64   `json:"distance" gorm:"type:decimal(10,4)"`
	AvgSpeed    float64   `json:"avg_speed" gorm:"type:decimal(10,4)"`
	Duration    int       `json:"duration"`
	TotalPing   int       `json:"total_ping"`
	Points      []Point   `json:"points" gorm:"foreignKey:TrjID;references:TrjID"`
}

func (t *Trajectory) String() string {
	return fmt.Sprintf(
		"Trajectory<%d|%s|%s|%v|%v>",
		t.TrjID,
		t.DrivingMode,
		t.OsName,
		t.StartedAt,
		t.EndedAt,
	)
}
