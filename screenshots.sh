#!/usr/bin/env bash

echo "cleaning assets dir..."
rm -rf assets/*
echo "cleaned that shit."
echo "generating screenshots..."

distros=("nixos" "arch" "debian" "bazzite" "void" "gentoo")

for i in "${!distros[@]}"; do
	freeze --execute "go run . --distro=${distros[i]}" -o assets/${distros[i]}.png --theme github-dark
done

echo "generated and saved screenshots! see you next time, cunt. x)"

