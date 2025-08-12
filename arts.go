/*
in order to keep my precious main.go tidy and simple,
this file will contain all of the distro arts. :)))
*/

package main

import "gotcha/color"

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
}

// quick function to find art by name lol
// btw why go be so complex when trying to filter
// arrays? LOLLLLLLLLEWOJIDGVHJ
func FindArt(name string) *Art {
	for i := range Arts {
		if Arts[i].Name == name {
			return &Arts[i]
		}
	}
	return nil
}
