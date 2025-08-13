package main

import "fmt"

func FormatDuration(secs int) string {
	if secs < 0 {
		return "unknown duration"
	}

	h := secs / 3600
	m := (secs % 3600) / 60
	s := secs % 60

	return fmt.Sprintf("%dh %dm %ds", h, m, s)
}

func HumanBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}

	div, exp := unit, 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	pre := "KMGTPE"[exp]

	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), pre)
}
