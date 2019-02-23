#!/bin/bash

echo "start getting vimrc from github..."
git clone git://github.com/hellodudu/vimrc_backup.git ~/.vim_runtime

echo "start using vimrc config..."
cd ~/.vim_runtime/
sh ./install_awesome_vimrc.sh

echo "start getting submodule plugins..."
git submodule update --init --recursive

echo "start getting dependence tools..."
./sources_non_forked/YouCompleteMe/install.py --go-completer
