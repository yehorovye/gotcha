#!/usr/bin/env bash

echo "building statically..."
make build-static
echo "done! :D"

echo "generating screenshots..."

distros=("nixos" "arch" "debian" "bazzite" "void" "gentoo")

for i in "${!distros[@]}"; do
	freeze --execute "./bin/gotcha --distro=${distros[i]}" -o assets/${distros[i]}.png
done

echo "generated and saved screenshots! see you next time, cunt. x)"

