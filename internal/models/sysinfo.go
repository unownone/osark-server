package models

import (
	"time"

	"gorm.io/gorm"
)

// SystemInfo is the information about the system
type SystemInfo struct {
	ID             string         `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt      *time.Time     `gorm:"index, autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt      *time.Time     `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UptimeSeconds  *int64         `json:"uptime_seconds" gorm:"uptime_ns"`        // Uptime seconds of the system
	OSQueryVersion string         `json:"osquery_version" gorm:"osquery_version"` // Version of osquery
	OSName         string         `json:"os_name" gorm:"os_name"`                 // Name of the operating system
	OSVersion      string         `json:"os_version" gorm:"os_version"`           // Version of the operating system
	OSArch         string         `json:"os_arch" gorm:"os_arch"`                 // Architecture of the operating system
	MacAddress     string         `json:"mac_address" gorm:"mac_address"`         // Mac address of the system
}
