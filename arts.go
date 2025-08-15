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
	Name string
	Art  []string
}

var Arts []Art = []Art{
	{
		Name: "nixos",
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
		Name: "debian",
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
		Name: "bazzite",
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
		Name: "arch",
		Art: []string{
			color.Colorize("                   ++                   ", color.Blue),
			color.Colorize("                  ++++                  ", color.Blue),
			color.Colorize("                 ++++++                 ", color.Blue),
			color.Colorize("                 +++++++                ", color.Blue),
			color.Colorize("                ++++++++                ", color.Blue),
			color.Colorize("               ++++++++++               ", color.Blue),
			color.Colorize("              ++ +++++++++              ", color.Blue),
			color.Colorize("             ++++++ +++++++             ", color.Blue),
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
		Name: "gentoo",
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
		Name: "manjaro",
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
		Name: "void",
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
		Name: "void-textless",
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
		Name: "artix",
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
		Name: "freebsd",
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
		Name: "opensuse",
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
		Name: "windows", // note: i was FORCED by a friend to make windows support,
		// my unique condition was that he had to make windows support.
		Art: []string{
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("                                        ", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
			color.Colorize("+++++++++++++++++++  +++++++++++++++++++", color.Cyan),
		},
	},
	{
		Name: "none",
		Art:  []string{},
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
