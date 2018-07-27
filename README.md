# AYUNiS.nvim
All YoU Need Is Spotify (only macOS)

## Require
- go >= 1.9.1
- Spotify.app

## Installation
1. Plz write below code in `$XDG_CONFIG_HOME/nvim/init.vim` if you use dein.vim
``` vim
" e.g.
call dein#add("NoahOrberg/AYUNiS.nvim")
" or
Plug 'NoahOrberg/AYUNiS.nvim'
```
and set runtimepath
``` vim
" e.g.
let g:ayunis_rtp = $HOME . '/.vim/plugged/AYUNiS.nvim'
```

2. please execute install command

3. execute `$ go get -u github.com/NoahOrberg/AYUNiS.nvim` because install binary

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
or `<SPACE>sl` in normal mode

### Prev Track
``` vim
:call AYUNiSPrev()
```
or `<SPACE>sh` in normal mode

### Toggle Play & Pause
``` vim
:call AYUNiSToggle()
```
or `<SPACE>st` in normal mode

### Volume Control
``` vim
:call AYUNiSVolumeUp()
:call AYUNiSVolumeDown()
```
or `<SPACE>s+` and `<SPACE>s-` in normal mode

### Toggle REPEAT
``` vim
:call AYUNiSToggleRepeat()
```
or `<SPACE>sr` in normal mode

### Toggle SHUFFLE
``` vim
:call AYUNiSToggleShuffle()
```
or `<SPACE>sf` in normal mode

