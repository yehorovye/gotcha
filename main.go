package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gotcha/color"
	cfg "gotcha/config"
)

var (
	distro string
	config map[string]string
)

func init() {
	var cfgPath string

	flag.StringVar(&distro, "distro", "", "distro name, art purposes")
	flag.StringVar(&cfgPath, "config", "", "path to config file")
	flag.Parse()

	if cfgPath == "" {
		homeDir, _ := os.UserHomeDir()
		cfgPath = filepath.Join(homeDir, ".config", "gotcha")
	}

	if _, err := os.Stat(cfgPath); err == nil {
		if m, err := cfg.LoadFile(cfgPath); err == nil {
			config = m
		}
	}

	if config == nil {
		config = make(map[string]string)
	}
}

func main() {
	term, colorterm := os.Getenv("TERM"), os.Getenv("COLORTERM")

	art := FindArt(distro, strings.ToLower(GetDistro()))
	if art == nil {
		art = &Arts[len(Arts)-1]
	}

	user := os.Getenv("USER")
	hostname, err := os.Hostname()
	if err != nil {
		hostname = unknown
	}

	divider := config["DIVIDER"]
	if divider == "" {
		divider = "-"
	}

	info := []string{}

	if !IsDisabled("userhost") {
		info = append(info,
			fmt.Sprintf("%s@%s",
				color.Colorize(user, art.Accent),
				color.Colorize(hostname, color.Red)),
			strings.Repeat(color.Colorize(divider, color.Red), len(fmt.Sprintf("%s@%s", user, hostname))),
		)
	}
	if !IsDisabled("os") {
		distroName := GetDistro()
		info = append(info, fmt.Sprintf("%s %s", color.Colorize("os:", art.Accent), distroName))
	}
	if !IsDisabled("kernel") {
		kernel := GetKernel()
		info = append(info, fmt.Sprintf("%s %s", color.Colorize("kernel:", art.Accent), kernel))
	}
	if !IsDisabled("memory") {
		memoryUsage := GetMemoryUsage()
		info = append(info, fmt.Sprintf("%s %s", color.Colorize("memory:", art.Accent), memoryUsage))
	}
	if !IsDisabled("uptime") {
		uptime := GetUptime()
		info = append(info, fmt.Sprintf("%s %s", color.Colorize("uptime:", art.Accent), uptime))
	}
	if !IsDisabled("shell") {
		shell := GetShell()
		info = append(info, fmt.Sprintf("%s %s", color.Colorize("shell:", art.Accent), shell))
	}
	if !IsDisabled("desktop") {
		de := GetDE()
		info = append(info, fmt.Sprintf("%s %s", color.Colorize("desktop:", art.Accent), de))
	}
	if !IsDisabled("term") {
		info = append(info, fmt.Sprintf("%s %s (%s)",
			color.Colorize("term:", art.Accent),
			term, color.Colorize(colorterm, color.Green)))
	}
	if !IsDisabled("packages") {
		pkgCounts := GetPkgCounts()
		info = append(info, fmt.Sprintf("%s %s", color.Colorize("packages:", art.Accent), pkgCounts))
	}
	if !IsDisabled("disks") {
		for _, d := range GetDisksUsage() {
			info = append(info,
				fmt.Sprintf("%s %s / %s (%.1f%% used)",
					color.Colorize(fmt.Sprintf("disk (%s):", d.MountPoint), art.Accent),
					HumanBytes(d.Used),
					HumanBytes(d.Total),
					d.UsedPct,
				),
			)
		}
	}

	m := max(len(info), len(art.Art))
	for i := range m {
		left, right := "", ""
		if i < len(art.Art) {
			left = art.Art[i]
		}
		if i < len(info) {
			right = info[i]
		}
		fmt.Printf("%s %s\n", left, right)
	}
}
