#!/bin/bash

echo "start getting vimrc from github..."
git clone --depth=1 git://github.com/hellodudu/vimrc_backup.git ~/.vim_runtime

echo "start using vimrc config..."
cd ~/.vim_runtime/
sh ./install_awesome_vimrc.sh

echo "start getting submodule plugins..."
git submodule update --init --recursive

echo "start getting dependence tools..."

if ! [ -x "$(command -v cmake)" ]; then
  echo "installing cmake..."
  brew install cmake
fi

if ! [ -x "$(command -v ctags)" ]; then
  echo "installing ctags..."
  brew install ctags
fi

./sources_non_forked/YouCompleteMe/install.py --go-completer
