package models

import (
	"time"

	"gorm.io/gorm"
)

// AppInfo is the information about an app
type AppInfo struct {
	ID             *string         `gorm:"primaryKey" json:"id,omitempty"` // ID of the app
	CreatedAt      *time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt      *time.Time      `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt      *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Name           string          `json:"name,omitempty" gorm:"name"`                         // Name of the app
	BundleName     string          `json:"bundle_name,omitempty" gorm:"bundle_name"`           // Bundle name of the app
	BundleID       string          `json:"bundle_id,omitempty" gorm:"bundle_id"`               // Bundle ID of the app
	BundleVersion  string          `json:"bundle_version,omitempty" gorm:"bundle_version"`     // Bundle version of the app
	Path           string          `json:"path,omitempty" gorm:"path"`                         // Path of the app
	LastOpenedTime time.Time       `json:"last_opened_time,omitempty" gorm:"last_opened_time"` // Last opened time of the app
	SysInfoID      *string         `json:"sys_info_id,omitempty" gorm:"sys_info_id"`
	// Using embedded assignment for SysInfo rather than a foreign key reference
	SysInfo *SystemInfo `json:"sys_info,omitempty" gorm:"foreignKey:SysInfoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// ProcessInfo is the information about a process
type ProcessInfo struct {
	ID            string          `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt     time.Time       `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt     time.Time       `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt     *gorm.DeletedAt `gorm:"index,omitempty" json:"deleted_at,omitempty"`
	PID           string          `json:"pid,omitempty" gorm:"pid"`
	Name          string          `json:"name,omitempty" gorm:"name"`
	BundleID      string          `json:"bundle_id,omitempty" gorm:"bundle_id"`
	BundleVersion string          `json:"bundle_version,omitempty" gorm:"bundle_version"`
	Path          string          `json:"path,omitempty" gorm:"path"`
	AppInfoID     string          `json:"app_info_id,omitempty" gorm:"column:app_info_id"`
	SysInfoID     string          `json:"sys_info_id,omitempty" gorm:"column:sys_info_id"`
	SysInfo       *SystemInfo     `json:"sys_info,omitempty" gorm:"foreignKey:SysInfoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
