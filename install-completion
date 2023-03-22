#!/bin/sh
shell="$1"

if [ "$shell" = "" ]; then
  echo "usage: install-completion <shell>"
  exit
fi

completion=$(apisense completion "$shell")

# check completion shell
if ! command -v "$shell" >/dev/null; then
  echo "$shell could not be found"
  exit
fi

# install completion
case "$shell" in
bash)
  echo "$completion" | sudo tee /etc/bash_completion.d/apisense
  ;;
zsh)
  echo "$completion" | sudo tee /usr/share/zsh/site-functions/_apisense
  ;;
fish)
  echo "$completion" | sudo tee /usr/local/share/fish/vendor_completions.d/apisense.fish
  ;;
powershell)
  echo "$completion" | sudo tee /usr/local/share/powershell/Modules/apisense/apisense.psm1
  ;;
*)
  echo "unsupported shell"
  exit
  ;;
esac
