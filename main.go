package main

import (
	"flag"
	"fmt"
	"os"
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

func getMemoryUsage() int {
	return 0
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

var distro string

func init() {
	flag.StringVar(&distro, "distro", "", "distro name, art purposes")
	flag.Parse()
}

func main() {
	memoryUsage := getMemoryUsage()
	uptime := getUptime()
	shell := getShell()

	art := FindArt(distro)
	if art == nil {
		art = &Arts[0] // todo: better fallbacks ig
	}

	info := []string{
		fmt.Sprintf("%s %d (mocked)", color.Colorize("memory:", color.BrightCyan), memoryUsage),
		fmt.Sprintf("%s %s", color.Colorize("uptime:", color.BrightCyan), uptime),
		fmt.Sprintf("%s %s", color.Colorize("shell:", color.BrightCyan), shell),
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
