package models

import (
	"fmt"
	"time"
)

type Point struct {
	PointID       int       `json:"point_id" gorm:"primaryKey"`
	TrjID         int       `json:"trj_id" gorm:"foreignKey:trj_id;references:trajectories;index;not null"`
	PingTimestamp time.Time `json:"pingtimestamp" gorm:"not null"`
	Lat           float64   `json:"lat" gorm:"type:decimal(18,8);not null"`
	Lon           float64   `json:"lon" gorm:"type:decimal(18,8);not null"`
	Speed         float64   `json:"speed" gorm:"type:decimal(10,4);not null"`
	Bearing       int       `json:"bearing" gorm:"not null"`
	Accuracy      float64   `json:"accuracy" gorm:"type:decimal(10,4);not null"`
	Trajectory    Trajectory `json:"trajectory" gorm:"foreignKey:TrjID;references:TrjID"`
}

func (p *Point) String() string {
	return fmt.Sprintf(
		"Point<%d|%v|%v|%v>",
		p.PointID,
		p.Lat,
		p.Lon,
		p.PingTimestamp,
	)
}
