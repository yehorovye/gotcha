package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"gotcha/color"
)

type DiskUsage struct {
	MountPoint string
	Used       uint64
	Total      uint64
	UsedPct    float64
}

type Field struct {
	Name string
	Text string
}

func IsDisabled(name string) bool {
	if v, ok := config["DISABLE"]; ok {
		for p := range strings.SplitSeq(v, ",") {
			if strings.EqualFold(strings.TrimSpace(p), name) {
				return true
			}
		}
	}
	return false
}

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

var unknown string = color.Colorize("unknown", color.Red)

func GetDistro() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return unknown
	}

	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		if value, ok := strings.CutPrefix(line, "PRETTY_NAME="); ok {
			return strings.Trim(value, `"'`)
		}
	}

	return unknown
}

func GetPkgCounts() string {
	counts := make(map[string]int)

	pathExists := func(path string) bool {
		_, err := os.Stat(path)
		return err == nil
	}

	// nix
	if pathExists("/run/current-system/sw/bin") {
		if entries, err := os.ReadDir("/run/current-system/sw/bin"); err == nil {
			counts["nix"] = len(entries)
		}
	}

	// dpkg
	if pathExists("/usr/bin/dpkg-query") {
		cmd := exec.Command("/usr/bin/dpkg-query", "-f", ".", "-W")
		if out, err := cmd.Output(); err == nil {
			counts["dpkg"] = len(out)
		}
	}

	// rpm
	if pathExists("/usr/bin/rpm") {
		cmd := exec.Command("/usr/bin/rpm", "-qa")
		if out, err := cmd.Output(); err == nil {
			counts["rpm"] = bytes.Count(out, []byte{'\n'})
		}
	}

	// pacman
	if pathExists("/usr/bin/pacman") {
		cmd := exec.Command("/usr/bin/pacman", "-Q")
		if out, err := cmd.Output(); err == nil {
			counts["pacman"] = bytes.Count(out, []byte{'\n'})
		}
	}

	// flatpak
	if pathExists("/usr/bin/flatpak") {
		cmd := exec.Command("/usr/bin/flatpak", "list")
		if out, err := cmd.Output(); err == nil {
			counts["flatpak"] = bytes.Count(out, []byte{'\n'})
		}
	}

	s := ""
	for k, v := range counts {
		if s != "" {
			s += ", "
		}
		s += fmt.Sprintf("%s - %d", k, v)
	}

	return s
}

func GetMemoryUsage() string {
	meminfo, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return unknown
	}

	var total, available uint64

	lines := strings.SplitSeq(string(meminfo), "\n")
	for line := range lines {
		if val, ok := strings.CutPrefix(line, "MemTotal:"); ok {
			total, _ = strconv.ParseUint(strings.Fields(strings.TrimSpace(val))[0], 10, 64)
		} else if val, ok := strings.CutPrefix(line, "MemAvailable:"); ok {
			available, _ = strconv.ParseUint(strings.Fields(strings.TrimSpace(val))[0], 10, 64)
		}
	}

	if total == 0 {
		return unknown
	}

	totalBytes := total * 1024
	availableBytes := available * 1024
	usedBytes := totalBytes - availableBytes
	usedPct := float64(usedBytes) / float64(totalBytes) * 100

	var usageColor string
	switch {
	case usedPct >= 80:
		usageColor = color.BrightRed
	case usedPct >= 50:
		usageColor = color.BrightYellow
	default:
		usageColor = color.BrightGreen
	}

	return fmt.Sprintf("%s / %s (%s used)",
		HumanBytes(usedBytes),
		HumanBytes(totalBytes),
		color.Colorize(fmt.Sprintf("%.1f%%", usedPct), usageColor),
	)
}

func GetUptime() string {
	up, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return unknown
	}

	uptime := strings.Split(string(up), " ")[0]
	secs, err := strconv.ParseFloat(uptime, 64)
	if err != nil {
		return unknown
	}

	return FormatDuration(int(secs))
}

func GetShell() string {
	sh := os.Getenv("SHELL")
	if sh == "" {
		sh = unknown
	}

	shell := strings.Split(sh, "/")

	return shell[len(shell)-1]
}

func GetDisksUsage() []DiskUsage {
	out, err := exec.Command("df", "-B1").Output()
	if err != nil {
		return nil
	}

	lines := bytes.Split(out, []byte("\n"))
	if len(lines) < 2 {
		return nil
	}

	var results []DiskUsage

	for _, line := range lines[1:] {
		if len(strings.TrimSpace(string(line))) == 0 {
			continue
		}

		cols := strings.Fields(string(line))
		if len(cols) < 6 {
			continue
		}

		total, err1 := strconv.ParseUint(cols[1], 10, 64)
		used, err2 := strconv.ParseUint(cols[2], 10, 64)
		mountPoint := cols[5]

		if err1 != nil || err2 != nil || total == 0 {
			continue
		}

		usedPct := float64(used) / float64(total) * 100

		results = append(results, DiskUsage{
			MountPoint: mountPoint,
			Used:       used,
			Total:      total,
			UsedPct:    usedPct,
		})
	}

	if len(results) == 0 {
		return nil
	}

	return results
}

func GetKernel() string {
	data, err := os.ReadFile("/proc/version")
	if err == nil {
		fields := strings.Fields(string(data))
		if len(fields) >= 3 {
			return fields[2]
		}
	}

	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return unknown
	}
	return strings.TrimSpace(string(out))
}

func GetDE() string {
	de := os.Getenv("XDG_CURRENT_DESKTOP")
	if de == "" {
		de = os.Getenv("DESKTOP_SESSION")
		if de == "" {
			de = unknown
		}
	}

	return de
}
