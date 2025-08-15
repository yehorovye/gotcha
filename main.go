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

var unknown string = color.Colorize("unknown", color.Red)

func getDistro() string {
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
	meminfo, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return color.Colorize("couldn't read meminfo", color.BrightRed)
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
		return color.Colorize("meminfo missing MemTotal", color.BrightRed)
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
		color.Colorize(HumanBytes(availableBytes), color.Green),
		color.Colorize(HumanBytes(totalBytes), color.Green),
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

	return FormatDuration(int(secs))
}

func getShell() string {
	sh := os.Getenv("SHELL")
	if sh == "" {
		sh = "unknown"
	}

	shell := strings.Split(sh, "/")

	return shell[len(shell)-1]
}

func getDiskUsage() string {
	out, err := exec.Command("df", "-B1", "/").Output()
	if err != nil {
		return color.Colorize("couldn't get disk usage", color.BrightRed)
	}

	lines := bytes.Split(out, []byte("\n"))
	if len(lines) < 2 {
		return color.Colorize("unexpected df output", color.BrightRed)
	}

	cols := strings.Fields(string(lines[1]))
	if len(cols) < 3 {
		return color.Colorize("unexpected df columns", color.BrightRed)
	}

	total, err1 := strconv.ParseUint(cols[1], 10, 64)
	used, err2 := strconv.ParseUint(cols[2], 10, 64)
	if err1 != nil || err2 != nil {
		return color.Colorize("failed parsing df output", color.BrightRed)
	}

	available := total - used
	usedPct := float64(used) / float64(total) * 100

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
		color.Colorize(HumanBytes(available), color.Green),
		color.Colorize(HumanBytes(total), color.Green),
		color.Colorize(fmt.Sprintf("%.1f%%", usedPct), usageColor),
	)
}

func getKernel() string {
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

func getDE() string {
	de := os.Getenv("XDG_CURRENT_DESKTOP")
	if de == "" {
		de = os.Getenv("DESKTOP_SESSION")
		if de == "" {
			de = unknown
		}
	}

	return de
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
	pkgCounts := getPkgCounts()
	diskUsage := getDiskUsage()
	distroName := getDistro()
	kernel := getKernel()
	de := getDE()

	term, colorterm := os.Getenv("TERM"), os.Getenv("COLORTERM")

	art := FindArt(distro, strings.ToLower(distroName))
	if art == nil {
		art = &Arts[len(Arts)-1]
	}

	user := os.Getenv("USER")
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	divider := os.Getenv("DIVIDER")
	if divider == "" {
		divider = "-"
	}

	info := []string{
		fmt.Sprintf("%s@%s", color.Colorize(user, color.Green), color.Colorize(hostname, color.Yellow)),
		fmt.Sprintf("%s", strings.Repeat(color.Colorize(divider, color.Yellow), len(fmt.Sprintf("%s@%s", user, hostname)))),
		fmt.Sprintf("%s %s", color.Colorize("os:", color.BrightCyan), distroName),
		fmt.Sprintf("%s %s", color.Colorize("kernel:", color.BrightCyan), kernel),
		fmt.Sprintf("%s %s", color.Colorize("memory:", color.BrightCyan), memoryUsage),
		fmt.Sprintf("%s %s", color.Colorize("disk usage:", color.BrightCyan), diskUsage),
		fmt.Sprintf("%s %s", color.Colorize("uptime:", color.BrightCyan), uptime),
		fmt.Sprintf("%s %s", color.Colorize("shell:", color.BrightCyan), shell),
		fmt.Sprintf("%s %s", color.Colorize("desktop:", color.BrightCyan), de),
		fmt.Sprintf("%s %s (%s)", color.Colorize("term:", color.BrightCyan), term, color.Colorize(colorterm, color.Green)),
		fmt.Sprintf("%s %s", color.Colorize("packages:", color.BrightCyan), pkgCounts),
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

		fmt.Printf("%s %s\n", left, right)
	}
}
