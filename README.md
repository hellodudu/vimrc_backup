# Vimrc Backup

make some modify from [amix's Ultimate Vimrc(Awesome mode)](https://github.com/amix/vimrc)

## Changed List

* **Add Plugins**
    * [auto-pairs](https://github.com/jiangmiao/auto-pairs) :
    ``
    Insert or delete brackets, parens, quotes in pair.
    ``
    * [neocomplete](https://github.com/Shougo/neocomplete) :
    ``
    Code auto complete
    ``
    * [tagbar](https://github.com/majutsushi/tagbar) :
    ``
    A class outline viewer for Vim
    ``
    * [vim-easymotion](https://github.com/easymotion/vim-easymotion) :
    ``
    Move cursor quickly
    ``

* **Plugins Modifies**
    * change NERDTree pos 
        >let g:NERDTreeWinPos = "left"

    * always show NERDTree bookmarks 
        >let NERDTreeShowBookmarks=1

    * change vim-multiple-cursors hotkeys
        >let g:multi_cursor_start_key='\<C-n\>'

    * add tagbar toggle hotkey
        >nnoremap \<silent\> \<leader\>tt :TagbarToggle\<cr\>

    * add vim-go hotkey
        >map \<leader\>gr :GoRun\<cr\>

        >map \<leader\>gd :GoDoc\<cr\>

* **Vimrc Modifies**
    * add set number
        >set number

    * change colorscheme
        >colorscheme desert
