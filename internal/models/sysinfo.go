package models

import (
	"time"

	"gorm.io/gorm"
)

// DBSystemInfo is the information about the system
// Database model for SystemInfo
type DBSystemInfo struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"index, autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	*SystemInfo
}

// SystemInfo is the information about the system
type SystemInfo struct {
	UptimeSeconds  time.Duration `json:"uptime_seconds" gorm:"uptime_seconds"`   // Uptime seconds of the system
	OSQueryVersion string        `json:"osquery_version" gorm:"osquery_version"` // Version of osquery
	OSName         string        `json:"os_name" gorm:"os_name"`                 // Name of the operating system
	OSVersion      string        `json:"os_version" gorm:"os_version"`           // Version of the operating system
	OSArch         string        `json:"os_arch" gorm:"os_arch"`                 // Architecture of the operating system
	MacAddress     string        `json:"mac_address" gorm:"mac_address"`         // Mac address of the system
}
