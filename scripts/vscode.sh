# Clone microsoft vscode from aur
git clone https://aur.archlinux.org/visual-studio-code-bin.git /tmp/visual-studio-code-bin

# Install vscode
cd /tmp/visual-studio-code-bin
yes | makepkg -si
