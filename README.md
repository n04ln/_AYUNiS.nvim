# AYUNiS.nvim
All YoU Need Is Spotify

## Require
- go >= 1.9.1
- [Shougo/dein.vim](https://github.com/Shougo/dein.vim) (Plugin Manager)

## Installation
1. Plz write below code in `$XDG_CONFIG_HOME/nvim/init.vim` if you use dein.vim
``` vim
call dein#add("NoahOrberg/AYUNiS.nvim")
```

2. `:call dein#update()`

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

" use airline
let g:airline_section_x = airline#section#create(['%{AYUNiSGetNowPlaying()}'])
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

