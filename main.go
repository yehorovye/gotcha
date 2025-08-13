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

	return fmt.Sprintf("memory: %s available (%s total) (%s used)",
		color.Colorize(HumanBytes(availableBytes), color.BrightGreen),
		color.Colorize(HumanBytes(totalBytes), color.BrightGreen),
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

	return fmt.Sprintf("%s available (%s total) (%s used)",
		color.Colorize(HumanBytes(available), color.BrightGreen),
		color.Colorize(HumanBytes(total), color.BrightGreen),
		color.Colorize(fmt.Sprintf("%.1f%%", usedPct), usageColor),
	)
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
	diskUsage := getDiskUsage()

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
		fmt.Sprintf("%s %s", color.Colorize("disk usage:", color.BrightCyan), diskUsage),
		fmt.Sprintf("%s %s", color.Colorize("uptime:", color.BrightCyan), uptime),
		fmt.Sprintf("%s %s", color.Colorize("shell:", color.BrightCyan), shell),
		fmt.Sprintf("%s %s", color.Colorize("packages:", color.BrightCyan), pkgCounts),
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
