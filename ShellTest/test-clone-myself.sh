#! /bin/bash
# add following line into .zshrc or .bashrc
alias gcms='cloneMyself() { git clone git@github.com:yongliang/$1.git ${2:-$1} ; }; cloneMyself'