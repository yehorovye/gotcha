package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"gotcha/color"
)

func formatDuration(secs int) string {
	if secs < 0 {
		return "unknown duration"
	}

	h := secs / 3600
	m := (secs % 3600) / 60
	s := secs % 60

	return fmt.Sprintf("%dh %dm %ds", h, m, s)
}

func getPkgCounts() string {
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

func getMemoryUsage() string {
	var total uint64
	var avail uint64

	mem, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return color.Colorize("couldn't read meminfo", color.BrightRed)
	}

	lines := strings.SplitSeq(string(mem), "\n")
	for line := range lines {
		if valueStr, ok := strings.CutPrefix(line, "MemTotal:"); ok {
			total, _ = strconv.ParseUint(strings.Fields(strings.TrimSpace(valueStr))[0], 10, 64)
		} else if valueStr, ok := strings.CutPrefix(line, "MemAvailable:"); ok {
			avail, _ = strconv.ParseUint(strings.Fields(strings.TrimSpace(valueStr))[0], 10, 64)
		}
	}

	if total == 0 {
		return color.Colorize("meminfo missing MemTotal", color.BrightRed)
	}

	totalF := float64(total) * 1024
	availF := float64(avail) * 1024

	usedPct := (totalF - availF) / totalF * 100

	var unit string
	var div float64
	if totalF >= 1024*1024*1024 {
		unit = "GB"
		div = 1024 * 1024 * 1024
	} else {
		unit = "MB"
		div = 1024 * 1024
	}

	var usageColor string
	switch {
	case usedPct >= 80:
		usageColor = color.BrightRed
	case usedPct >= 50:
		usageColor = color.BrightYellow
	default:
		usageColor = color.BrightGreen
	}

	return fmt.Sprintf("%s available (%s total) (%s used)",
		color.Colorize(fmt.Sprintf("%.2f%s", availF/div, unit), color.BrightGreen),
		color.Colorize(fmt.Sprintf("%.2f%s", totalF/div, unit), color.BrightGreen),
		color.Colorize(fmt.Sprintf("%.1f%%", usedPct), usageColor),
	)
}

func getUptime() string {
	up, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "couldn't read uptime"
	}

	uptime := strings.Split(string(up), " ")[0]
	secs, err := strconv.ParseFloat(uptime, 64)
	if err != nil {
		return "invalid duration"
	}

	return formatDuration(int(secs))
}

func getShell() string {
	sh := os.Getenv("SHELL")
	if sh == "" {
		sh = "unknown"
	}

	shell := strings.Split(sh, "/")

	return shell[len(shell)-1]
}

func getGPUNames() []string {
	// this function is slow asfuck lol
	if os.Getenv("SKIP_GPU_DATA") == "1" {
		return []string{}
	}

	out, err := exec.Command("lspci").Output()
	if err != nil {
		return []string{}
	}

	lines := strings.Split(string(out), "\n")
	var gpus []string

	for _, line := range lines {
		if strings.Contains(line, "VGA") || strings.Contains(line, "3D controller") {
			descStart := strings.Index(line, ": ")
			if descStart == -1 {
				continue
			}
			desc := strings.TrimSpace(line[descStart+2:])

			if start := strings.Index(desc, "["); start != -1 {
				if end := strings.Index(desc, "]"); end != -1 && end > start {
					gpus = append(gpus, strings.TrimSpace(desc[start+1:end]))
					continue
				}
			}

			parts := strings.Fields(desc)
			if len(parts) >= 3 {
				short := strings.Join(parts[len(parts)-3:], " ")
				gpus = append(gpus, short)
			} else {
				gpus = append(gpus, desc)
			}
		}
	}

	return gpus
}

var distro string

func init() {
	flag.StringVar(&distro, "distro", "", "distro name, art purposes")
	flag.Parse()
}

func main() {
	memoryUsage := getMemoryUsage()
	uptime := getUptime()
	shell := getShell()
	gpuNames := getGPUNames()
	pkgCounts := getPkgCounts()

	art := FindArt(distro)
	if art == nil {
		art = &Arts[0] // todo: better fallbacks ig
	}

	user := os.Getenv("USER")
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	info := []string{
		fmt.Sprintf("%s@%s", color.Colorize(user, color.Green), color.Colorize(hostname, color.Yellow)),
		fmt.Sprintf("%s %s", color.Colorize("memory:", color.BrightCyan), memoryUsage),
		fmt.Sprintf("%s %s", color.Colorize("uptime:", color.BrightCyan), uptime),
		fmt.Sprintf("%s %s", color.Colorize("shell:", color.BrightCyan), shell),
		fmt.Sprintf("%s %s", color.Colorize("packages:", color.BrightCyan), pkgCounts),
		fmt.Sprintf("%s %s", color.Colorize("disk usage:", color.BrightCyan), color.Colorize("not yet :P", color.BgRed)),
	}

	for gpu := range gpuNames {
		info = append(info, fmt.Sprintf("%s %s", color.Colorize(fmt.Sprintf("GPU #%d:", gpu), color.BrightCyan), color.Colorize(gpuNames[gpu], color.Green)))
	}

	m := max(len(info), len(art.Art))

	for i := range m {
		left := ""
		right := ""

		if i < len(art.Art) {
			left = art.Art[i]
		}
		if i < len(info) {
			right = info[i]
		}

		fmt.Printf("%-10s %s\n", left, right)
	}
}
