#!/bin/bash
git clone git://github.com/hellodudu/vimrc_backup.git ~/.vim_runtime
cd ~/.vim_runtime/
sh ./install_awesome_vimrc.sh
git submodule update --init --recursive
