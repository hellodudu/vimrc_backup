#!/bin/bash

echo "start getting vimrc from github..."
git clone --depth=1 git://github.com/hellodudu/vimrc_backup.git ~/.vim_runtime

echo "start using vimrc config..."
cd ~/.vim_runtime/
sh ./install_awesome_vimrc.sh

echo "start getting dependence tools..."

if ! [ -x "$(command -v cmake)" ]; then
  echo "installing cmake..."
  brew install cmake
fi

if ! [ -x "$(command -v ctags)" ]; then
  echo "installing ctags..."
  brew install ctags
fi

if ! [ -x "$(command -v the_silver_searcher)" ]; then
  echo "installing the_silver_searcher..."
  brew install the_silver_searcher
fi

