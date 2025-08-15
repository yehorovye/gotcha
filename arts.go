/*
in order to keep my precious main.go tidy and simple,
this file will contain all of the distro arts. :)))
*/

package main

import (
	"gotcha/color"
	"regexp"
	"strings"
)

type Art struct {
	Name   string
	Art    []string
	Accent string
}

var Arts []Art = []Art{
	{
		Name:   "nixos",
		Accent: color.Blue,
		Art: []string{
			color.Colorize("          ++      -----     --          ", color.Blue),
			color.Colorize("         ++++      -----   ----         ", color.Blue),
			color.Colorize("          ++++      ----- ----          ", color.Blue),
			color.Colorize("           +++++     --------           ", color.Blue),
			color.Colorize("      ++++++++++++++++------            ", color.Blue),
			color.Colorize("     ++++++++++++++++++-----     ++     ", color.Blue),
			color.Colorize("           -----        -----   ++++    ", color.Blue),
			color.Colorize("          -----          ----- ++++     ", color.Blue),
			color.Colorize("         -----            --- ++++      ", color.Blue),
			color.Colorize("-------------              -++++++++++++", color.Blue),
			color.Colorize("------------+              +++++++++++++", color.Blue),
			color.Colorize("      ---- +++            +++++         ", color.Blue),
			color.Colorize("     ---- +++++          +++++          ", color.Blue),
			color.Colorize("    ----   +++++        +++++           ", color.Blue),
			color.Colorize("     --     +++++------------------     ", color.Blue),
			color.Colorize("            ++++++----------------      ", color.Blue),
			color.Colorize("           ++++++++      ----           ", color.Blue),
			color.Colorize("          ++++ +++++      ----          ", color.Blue),
			color.Colorize("         ++++   +++++      ----         ", color.Blue),
			color.Colorize("          ++     +++++      --          ", color.Blue),
		},
	},
	{
		Name:   "debian",
		Accent: color.Red,
		Art: []string{
			color.Colorize("                -- .                    ", color.Red),
			color.Colorize("            ----+++------------         ", color.Red),
			color.Colorize("         ---+#+---------.----++---      ", color.Red),
			color.Colorize("       --++----              --+#+--    ", color.Red),
			color.Colorize("     --++--                    --+#+--  ", color.Red),
			color.Colorize("    -++--                        -+#+-. ", color.Red),
			color.Colorize("   ----                           -++ . ", color.Red),
			color.Colorize("  -+-               ......         -+-  ", color.Red),
			color.Colorize(" -++-             .                -#-  ", color.Red),
			color.Colorize(" -+-            ..                 .#-  ", color.Red),
			color.Colorize(" --             -                  -#-  ", color.Red),
			color.Colorize(" --            --                  -+-  ", color.Red),
			color.Colorize(" -.            --                  --   ", color.Red),
			color.Colorize(".--             --                --    ", color.Red),
			color.Colorize(" --              ..             .-.     ", color.Red),
			color.Colorize(" -+-               --        ....       ", color.Red),
			color.Colorize(" -+-               .-........           ", color.Red),
			color.Colorize("  ---.                                  ", color.Red),
			color.Colorize("   .+-                                  ", color.Red),
			color.Colorize("    ---                                 ", color.Red),
			color.Colorize("     ---                                ", color.Red),
			color.Colorize("      ---                               ", color.Red),
			color.Colorize("        .-.                             ", color.Red),
			color.Colorize("          ...                           ", color.Red),
			color.Colorize("             ...                        ", color.Red),
			color.Colorize("                ...                     ", color.Red),
		},
	},
	{
		Name:   "bazzite",
		Accent: color.Magenta,
		Art: []string{
			color.Colorize("      ######++++##########              ", color.Magenta),
			color.Colorize("    ########----##############          ", color.Magenta),
			color.Colorize("    ########-..-################        ", color.Magenta),
			color.Colorize("   #########----#################       ", color.Magenta),
			color.Colorize("   #########----##############+#+#+     ", color.Magenta),
			color.Colorize("   ++--..----------..--++++++++#+++     ", color.Magenta),
			color.Colorize("   +++----------------++++++-++++#++    ", color.Magenta),
			color.Colorize("   #########----########+#++++-+++++    ", color.Magenta),
			color.Colorize("   #########-..-######+#+#+++-+-+++++   ", color.Magenta),
			color.Colorize("   #########----####+##+#+#+-+-++++++   ", color.Magenta),
			color.Colorize("    ########++++#####+#+++++-+-+++++    ", color.Magenta),
			color.Colorize("    ########++++#+#+#+++++-+-+++++++    ", color.Magenta),
			color.Colorize("     #######+++++#+#+++++-+--++++++     ", color.Magenta),
			color.Colorize("     ########++++++-+++--+-++++++++     ", color.Magenta),
			color.Colorize("       ####+#++-+-++--+-+++++++++       ", color.Magenta),
			color.Colorize("        ##+##+#++++-++++++++++++        ", color.Magenta),
			color.Colorize("          #+++++++++++++++++++          ", color.Magenta),
			color.Colorize("              ++++++++++++              ", color.Magenta),
		},
	},
	{
		Name:   "arch",
		Accent: color.Blue,
		Art: []string{
			color.Colorize("                   ++                   ", color.Blue),
			color.Colorize("                  ++++                  ", color.Blue),
			color.Colorize("                 ++++++                 ", color.Blue),
			color.Colorize("                ++++++++                ", color.Blue),
			color.Colorize("                +++++++++               ", color.Blue),
			color.Colorize("              +  +++++++++              ", color.Blue),
			color.Colorize("             +++++  +++++++             ", color.Blue),
			color.Colorize("            ++++++++++++++++            ", color.Blue),
			color.Colorize("           ++++++++++++++++++           ", color.Blue),
			color.Colorize("          ++++++++++++++++++++          ", color.Blue),
			color.Colorize("         ++++++++++++++++++++++         ", color.Blue),
			color.Colorize("        ++++++++++    ++++++++++        ", color.Blue),
			color.Colorize("       +++++++++        +++++++++       ", color.Blue),
			color.Colorize("      ++++++++++        ++++++++++      ", color.Blue),
			color.Colorize("     +++++++++++         +++++++ ++     ", color.Blue),
			color.Colorize("    ++++++++++++         +++++++++      ", color.Blue),
			color.Colorize("   +++++++++++            +++++++++++   ", color.Blue),
			color.Colorize("  +++++++                      +++++++  ", color.Blue),
			color.Colorize(" ++++                              ++++ ", color.Blue),
			color.Colorize("++                                    ++", color.Blue),
		},
	},
	{
		Name:   "gentoo",
		Accent: color.Magenta,
		Art: []string{
			color.Colorize("           --.   ...---                 ", color.Magenta),
			color.Colorize("        -. .      ......--              ", color.Magenta),
			color.Colorize("      -.          ........---           ", color.Magenta),
			color.Colorize("    --    .  .  . .........----         ", color.Magenta),
			color.Colorize("   --          .. ..---....------       ", color.Magenta),
			color.Colorize("   -.          -.-----++...------.-     ", color.Magenta),
			color.Colorize("   ---          ---  -++...-------..    ", color.Magenta),
			color.Colorize("   -----.         .........--------     ", color.Magenta),
			color.Colorize("     ----....     ........---------  -  ", color.Magenta),
			color.Colorize("        ----..   .........--------. --  ", color.Magenta),
			color.Colorize("          .    ...........------.  ---  ", color.Magenta),
			color.Colorize("        ..     ..........-----.  .--    ", color.Magenta),
			color.Colorize("      .       ..........----.  .---     ", color.Magenta),
			color.Colorize("    -       ...........---   ----       ", color.Magenta),
			color.Colorize("   .     .............-.  .++++         ", color.Magenta),
			color.Colorize("  - .................   -+++            ", color.Magenta),
			color.Colorize("  -...............   +++++              ", color.Magenta),
			color.Colorize("  +-..........   -+++++                 ", color.Magenta),
			color.Colorize("   ++.      -+++++++                    ", color.Magenta),
			color.Colorize("     +++++++++++                        ", color.Magenta),
		},
	},
	{
		Name:   "manjaro",
		Accent: color.Green,
		Art: []string{
			color.Colorize("--------------------------   -----------", color.Green),
			color.Colorize("--------------------------   -----------", color.Green),
			color.Colorize("--------------------------   -----------", color.Green),
			color.Colorize("--------------------------   -----------", color.Green),
			color.Colorize("--------------------------   -----------", color.Green),
			color.Colorize("--------------------------   -----------", color.Green),
			color.Colorize("-----------                  -----------", color.Green),
			color.Colorize("-----------                  -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
			color.Colorize("-----------   ------------   -----------", color.Green),
		},
	},
	{
		Name:   "void",
		Accent: color.Green,
		Art: []string{
			color.Colorize("                   ------------                   ", color.Green),
			color.Colorize("               --------------------               ", color.Green),
			color.Colorize("              ------------------------            ", color.Green),
			color.Colorize("                --------  --------------          ", color.Green),
			color.Colorize("                                ----------        ", color.Green),
			color.Colorize("       +++                         --------       ", color.Green),
			color.Colorize("      ++++++                        --------      ", color.Green),
			color.Colorize("     ++++++++           --           --------     ", color.Green),
			color.Colorize("     +++++++        ----------        -------     ", color.Green),
			color.Colorize(" ######++++   ## ############--##### #############", color.White),
			color.Colorize("  #####++++##  #####----#####-##### #####---+#####", color.White),
			color.Colorize("   #######+   #####----#####-##### #####----##### ", color.White),
			color.Colorize("    ###++++    ###########--##### #############   ", color.White),
			color.Colorize("     +++++++        ----------        -------     ", color.Green),
			color.Colorize("      +++++++           --           -------      ", color.Green),
			color.Colorize("      ++++++++                        ------      ", color.Green),
			color.Colorize("       +++++++++                        ---       ", color.Green),
			color.Colorize("        ++++++++++              +                 ", color.Green),
			color.Colorize("          ++++++++++++++++++++++++                ", color.Green),
			color.Colorize("            ++++++++++++++++++++++++              ", color.Green),
			color.Colorize("               ++++++++++++++++++++               ", color.Green),
			color.Colorize("                   ++++++++++++                   ", color.Green),
		},
	},
	{
		Name:   "void-textless",
		Accent: color.Green,
		Art: []string{
			color.Colorize("              ------------              ", color.Green),
			color.Colorize("          --------------------          ", color.Green),
			color.Colorize("         ------------------------       ", color.Green),
			color.Colorize("           ------------------------     ", color.Green),
			color.Colorize("            --            ----------    ", color.Green),
			color.Colorize("   +++                       ---------  ", color.Green),
			color.Colorize("  ++++++                       -------  ", color.Green),
			color.Colorize(" +++++++           --           ------- ", color.Green),
			color.Colorize("+++++++        ----------        -------", color.Green),
			color.Colorize("+++++++       ------------       -------", color.Green),
			color.Colorize("+++++++       ------------       -------", color.Green),
			color.Colorize("+++++++       ------------       -------", color.Green),
			color.Colorize("+++++++       ------------       -------", color.Green),
			color.Colorize("+++++++        ----------        -------", color.Green),
			color.Colorize(" +++++++           --           ------- ", color.Green),
			color.Colorize("  +++++++                       ------  ", color.Green),
			color.Colorize("  +++++++++                       ----  ", color.Green),
			color.Colorize("    ++++++++++            ++            ", color.Green),
			color.Colorize("     ++++++++++++++++++++++++           ", color.Green),
			color.Colorize("       ++++++++++++++++++++++++         ", color.Green),
			color.Colorize("         +++++++++++++++++++++          ", color.Green),
			color.Colorize("              ++++++++++++              ", color.Green),
		},
	},
	{
		Name:   "artix",
		Accent: color.Cyan,
		Art: []string{
			color.Colorize("                   --                   ", color.Cyan),
			color.Colorize("                   --                   ", color.Cyan),
			color.Colorize("                  -.--                  ", color.Cyan),
			color.Colorize("                 -+.---                 ", color.Cyan),
			color.Colorize("                --#.----                ", color.Cyan),
			color.Colorize("               --+ .-----               ", color.Cyan),
			color.Colorize("              ----  .-----              ", color.Cyan),
			color.Colorize("             ---+###  .----             ", color.Cyan),
			color.Colorize("               ----+## ..---            ", color.Cyan),
			color.Colorize("                   ----# .---           ", color.Cyan),
			color.Colorize("          ---         ---#..--          ", color.Cyan),
			color.Colorize("         ----.---        --+.--         ", color.Cyan),
			color.Colorize("         ------..-----      ---         ", color.Cyan),
			color.Colorize("        ---------. ##-----              ", color.Cyan),
			color.Colorize("       -----------.   ---------         ", color.Cyan),
			color.Colorize("      ----------.   +##+------          ", color.Cyan),
			color.Colorize("     --------..  ###-----       ---     ", color.Cyan),
			color.Colorize("    ------... #+-----        ---.---    ", color.Cyan),
			color.Colorize("   -----...#----          -----  .---   ", color.Cyan),
			color.Colorize("  ----..----              ------#  .--  ", color.Cyan),
			color.Colorize(" ---.--                         ----..- ", color.Cyan),
			color.Colorize("---                                  ---", color.Cyan),
		},
	},
	{
		Name:   "freebsd",
		Accent: color.BrightRed,
		Art: []string{
			color.Colorize("++++++                             +++++", color.BrightRed),
			color.Colorize("++++++++++  +++++++++++++++++  +++++++++", color.BrightRed),
			color.Colorize(" +++++++ +++++++++++++++++  +++++++++++ ", color.BrightRed),
			color.Colorize(" +++++ +++++++++++++++++++ ++++++++++++ ", color.BrightRed),
			color.Colorize("  ++ +++++++++++++++++++++ +++++++++++  ", color.BrightRed),
			color.Colorize("    +++++++++++++++++++++++ +++++++++   ", color.BrightRed),
			color.Colorize("   ++++++++++++++++++++++++++ ++++++++  ", color.BrightRed),
			color.Colorize("  ++++++++++++++++++++++++++++++    +++ ", color.BrightRed),
			color.Colorize("  +++++++++++++++++++++++++++++++++++++ ", color.BrightRed),
			color.Colorize("  ++++++++++++++++++++++++++++++++++++++", color.BrightRed),
			color.Colorize(" +++++++++++++++++++++++++++++++++++++++", color.BrightRed),
			color.Colorize(" +++++++++++++++++++++++++++++++++++++++", color.BrightRed),
			color.Colorize("  +++++++++++++++++++++++++++++++++++++ ", color.BrightRed),
			color.Colorize("  +++++++++++++++++++++++++++++++++++++ ", color.BrightRed),
			color.Colorize("   +++++++++++++++++++++++++++++++++++  ", color.BrightRed),
			color.Colorize("    ++++++++++++++++++++++++++++++++++  ", color.BrightRed),
			color.Colorize("     +++++++++++++++++++++++++++++++    ", color.BrightRed),
			color.Colorize("      +++++++++++++++++++++++++++++     ", color.BrightRed),
			color.Colorize("        +++++++++++++++++++++++++       ", color.BrightRed),
			color.Colorize("           +++++++++++++++++++          ", color.BrightRed),
			color.Colorize("               +++++++++++              ", color.BrightRed),
		},
	},
	{
		Name:   "opensuse",
		Accent: color.Green,
		Art: []string{
			color.Colorize("     +++++++++++++++++++++              ", color.Green),
			color.Colorize("     +++             ++++++++           ", color.Green),
			color.Colorize("     +++                  +++++         ", color.Green),
			color.Colorize("     +++           ++       ++++        ", color.Green),
			color.Colorize("     +++       ++++++++++     +++       ", color.Green),
			color.Colorize("     +++     ++++      ++++    +++      ", color.Green),
			color.Colorize("     +++    +++     +++  +++   +++      ", color.Green),
			color.Colorize("     +++    +++     +++  +++    ++      ", color.Green),
			color.Colorize("     +++    +++          +++    +++     ", color.Green),
			color.Colorize("     +++    +++          +++    +++     ", color.Green),
			color.Colorize("     +++     ++++      ++++     +++     ", color.Green),
			color.Colorize("     +++       ++++++++++       +++     ", color.Green),
			color.Colorize("      +++ +++     ++++          +++     ", color.Green),
			color.Colorize("      +++  ++++                 +++     ", color.Green),
			color.Colorize("       +++   ++++++++++++++++++++++     ", color.Green),
			color.Colorize("        ++++    +++++++++++++++++++     ", color.Green),
			color.Colorize("         +++++++                        ", color.Green),
			color.Colorize("            +++++++++++++++++++         ", color.Green),
		},
	},
	{
		Name:   "none",
		Accent: color.Blue,
		Art:    []string{},
	},
}

/*
 * advanced function to find a distro art
 * based on the user input.
 *
 * priority: target distro (fake), actual distro.
 */
func FindArt(targetName, distro string) *Art {
	for i := range Arts {
		if Arts[i].Name == targetName {
			return &Arts[i]
		}
	}

	targetLower := strings.ToLower(distro)

	for i := range Arts {
		nameLower := strings.ToLower(Arts[i].Name)
		if strings.Contains(targetLower, nameLower) {
			return &Arts[i]
		}
	}

	pattern := `(?i)` + regexp.QuoteMeta(distro)
	re := regexp.MustCompile(pattern)

	for i := range Arts {
		if re.MatchString(Arts[i].Name) {
			return &Arts[i]
		}
	}

	return nil
}
