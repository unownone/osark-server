package models

import "time"

// Intent is the intent of the event
type Intent string

const (
	IntentInit Intent = "init"

	// App events
	IntentAppOpen      Intent = "app_open"
	IntentAppFocus     Intent = "app_focus"
	IntentAppBlur      Intent = "app_blur"
	IntentAppSwitch    Intent = "app_switch"
	IntentAppTerminate Intent = "app_terminate"
	IntentAppLaunch    Intent = "app_launch"
	IntentAppClose     Intent = "app_close"

	// Process events
	IntentRunningProcesses Intent = "running_processes"
)

// LogEvent is the event that is logged to the server
type LogEvent struct {
	Intent     Intent         `json:"intent"`                // Intent is the intent of the event
	AppInfo    []*AppInfo     `json:"app_info,omitempty"`    // AppInfo is the information about an app
	Error      string         `json:"error,omitempty"`       // Error is the error message
	SystemInfo *SystemInfo    `json:"system_info,omitempty"` // SystemInfo is the information about the system
	CreatedAt  time.Time      `json:"created_at"`            // CreatedAt is the time the event was created
	Processes  []*ProcessInfo `json:"processes,omitempty"`   // Processes is the information about the processes
}

