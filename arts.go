/*
in order to keep my precious main.go tidy and simple,
this file will contain all of the distro arts. :)))
*/

package main

import (
	"gotcha/color"
	"slices"
)

type Art struct {
	Name string
	Art  []string
}

var Arts []Art = []Art{
	{
		Name: "nixos",
		Art: []string{
			color.Colorize("     ≠≠   ÷÷ ÷÷     ", color.Blue),
			color.Colorize("      ≠≠   ÷÷÷      ", color.Blue),
			color.Colorize("  ≠≠≠≠≠≠≠≠≠≠÷÷  ≠≠  ", color.Blue),
			color.Colorize("     ÷÷      ÷÷≠≠   ", color.Blue),
			color.Colorize("÷÷÷÷÷÷        ≠≠≠≠≠≠", color.Blue),
			color.Colorize("   ÷÷≠≠      ≠≠     ", color.Blue),
			color.Colorize("  ÷÷  ≠≠÷÷÷÷÷÷÷÷÷÷  ", color.Blue),
			color.Colorize("      ≠≠≠   ÷÷      ", color.Blue),
			color.Colorize("     ≠≠ ≠≠   ÷÷      ", color.Blue),
		},
	},
	{
		Name: "debian",
		Art: []string{
			color.Colorize("       ≈≈≈≈≈≈       ", color.Red),
			color.Colorize("    ≈≈≈≈     ≠≈≈≈   ", color.Red),
			color.Colorize("  ≈≈≈          ≈≈≈  ", color.Red),
			color.Colorize("  ≈≈     ≈      ≈≈  ", color.Red),
			color.Colorize("  ≈     ≈       ≈≈  ", color.Red),
			color.Colorize("  ≈     ≈    ≈  ≈   ", color.Red),
			color.Colorize("  ≈     ≈≈    ≈     ", color.Red),
			color.Colorize("   ≈≈               ", color.Red),
			color.Colorize("    ≈               ", color.Red),
			color.Colorize("     ≈              ", color.Red),
			color.Colorize("       ≈≈           ", color.Red),
		},
	},
	{
		Name: "arch",
		Art: []string{
			color.Colorize("         ÷          ", color.Blue),
			color.Colorize("        ÷÷÷         ", color.Blue),
			color.Colorize("       ÷÷÷÷÷        ", color.Blue),
			color.Colorize("      ÷÷÷=÷÷÷       ", color.Blue),
			color.Colorize("     ÷÷÷÷÷÷÷÷÷      ", color.Blue),
			color.Colorize("    ÷÷÷÷÷ ÷÷=÷÷     ", color.Blue),
			color.Colorize("   ÷÷=÷÷   ÷÷÷÷÷    ", color.Blue),
			color.Colorize("  ÷÷÷÷÷     ÷÷=÷÷   ", color.Blue),
			color.Colorize(" ÷÷=           ÷÷÷  ", color.Blue),
			color.Colorize("=                 = ", color.Blue),
		},
	},
	{
		Name: "gentoo",
		Art: []string{
			color.Colorize("      ==≠==         ", color.White),
			color.Colorize("   =≈≈≈≈≈≠≠≠≠=      ", color.White),
			color.Colorize("  ≠≈≈≈≈≈≈≠≠"+color.Colorize("0", color.BrightMagenta)+color.Colorize("0", color.BrightMagenta)+color.Colorize("0", color.BrightMagenta)+"≠≠=   ", color.White),
			color.Colorize(" ×≠≈≈≈≈≈=××"+color.Colorize("0", color.BrightMagenta)+color.Colorize("0", color.BrightMagenta)+color.Colorize("0", color.BrightMagenta)+color.Colorize("0", color.BrightMagenta)+"≠≠=  ", color.White),
			color.Colorize("  ×÷÷≠≈≈≈≈≈≠≠≠≠≠≠≠≠ ", color.White),
			color.Colorize("     ÷≈≈≈≈≈≠≠≠≠≠≠≈× ", color.White),
			color.Colorize("   ÷≈≈≈≈≈≈≈≠≠≠≠≈÷÷  ", color.White),
			color.Colorize("  ≠≈≈≈≈≈≈≈≠≠≠≈=×    ", color.White),
			color.Colorize(" =≠≠≠≠≠≠≠≠≈=÷×      ", color.White),
			color.Colorize(" ××≠≠≠≠÷÷÷÷         ", color.White),
			color.Colorize("   ÷÷÷×             ", color.White),
		},
	},
	{
		Name: "windows", // note: i was FORCED by a friend to make windows support,
		// my unique condition was that he had to make windows support.
		Art: []string{
			color.Colorize(" ÷÷÷×××× ====÷÷÷÷÷÷÷", color.BrightBlue),
			color.Colorize("÷÷×××××× ===÷÷=÷÷÷÷÷", color.BrightBlue),
			color.Colorize("÷××××××× =÷÷÷÷÷÷÷÷÷÷", color.BrightBlue),
			color.Colorize("×××××××× ÷÷÷÷÷÷÷÷÷÷÷", color.BrightBlue),
			color.Colorize("                    ", color.BrightBlue),
			color.Colorize("××××××-- ÷÷÷÷÷÷÷÷÷××", color.BrightBlue),
			color.Colorize("××××---- ÷÷÷÷÷÷÷÷×××", color.BrightBlue),
			color.Colorize("××------ ÷÷÷÷÷÷×××××", color.BrightBlue),
			color.Colorize(" ------- ÷÷÷÷×××××××", color.BrightBlue),
		},
	},
}

// quick function to find art by name lol
func FindArt(name string) *Art {
	art := slices.IndexFunc(Arts, func(e Art) bool {
		return e.Name == name
	})

	if art == -1 {
		return nil
	}

	return &Arts[art]
}
