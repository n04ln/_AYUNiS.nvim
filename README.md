# AYUNiS.nvim
All YoU Need Is Spotify (only macOS)

## Require
- go >= 1.11.0
- Spotify.app

## Installation
1. execute `$ go get -u github.com/n04ln/AYUNiS.nvim` because install a binary

2. Plz write below code in `$XDG_CONFIG_HOME/nvim/init.vim`
``` vim
" if you use dein.vim
call dein#add("n04ln/AYUNiS.nvim")
" if you use vim-plug
Plug 'n04ln/AYUNiS.nvim'
```
and set runtimepath
``` vim
" e.g. if you use vim-plug, perhaps runtime path is this.
let g:ayunis_rtp = $HOME . '/.vim/plugged/AYUNiS.nvim'
```

3. please execute some command to install plugin (e.g. `:PlugInstall`)

4. Let's Enjoy `AYUNiS.nvim` Life!

## Usage
### Get the name of the song that is playing now
``` vim
:echo AYUNiSGetNowPlaying()
```
if you want to add to `statusline`, plz set as below
``` vim
" default
set statusline+=%!AYUNiSGetNowPlaying()

" use airline (show in the statusline
let g:airline_section_x = airline#section#create(['%{AYUNiSGetNowPlaying()}'])

" use lightline (show in the tabline
let g:lightline = {
      \ 'component_function': {
      \   'ayunis': 'AYUNiSGetNowPlaying'
      \ },
      \ }
let g:lightline.tabline          = {
      \ 'left': [['ayunis', 'buffers']],
      \ 'right': [['close']]
      \ }
```

### Next Track
``` vim
:call AYUNiSNext()
```

### Prev Track
``` vim
:call AYUNiSPrev()
```

### Toggle Play & Pause
``` vim
:call AYUNiSToggle()
```

### Volume Control
``` vim
:call AYUNiSVolumeUp()
:call AYUNiSVolumeDown()
```

### Toggle REPEAT
``` vim
:call AYUNiSToggleRepeat()
```

### Toggle SHUFFLE
``` vim
:call AYUNiSToggleShuffle()
```

## Keybinds
###  Example
``` vim
" Next
nnoremap <silent><SPACE>sl :call AYUNiSNext()<CR>
" Prev
nnoremap <silent><SPACE>sh :call AYUNiSPrev()<CR>
" Toggle(playpause)
nnoremap <silent><SPACE>st :call AYUNiSToggle()<CR>
" Volume up
nnoremap <silent><SPACE>s+ :call AYUNiSVolumeUp()<CR>
" Volume down
nnoremap <silent><SPACE>s- :call AYUNiSVolumeDown()<CR>
" Toggle Repeat
nnoremap <silent><SPACE>sr :call AYUNiSToggleRepeat()<CR>
" Toggle Shuffle
nnoremap <silent><SPACE>sf :call AYUNiSToggleShuffle()<CR>
```
