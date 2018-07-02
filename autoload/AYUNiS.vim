scriptencoding utf-8

if exists('g:loaded_AYUNiS')
    finish
endif
let g:loaded_AYUNiS = 1

let s:save_cpo = &cpo
set cpo&vim

function! ayunis#get_now_playing() abort
   return system("spotify_util/now_playing.applescript") 
endfunction

function! ayunis#reload_statusline() abort
    let i = 0
    for i < 100
        redrawstatus
        sleep 100m
        let i += 1
    endfor
endfunction

let &cpo = s:save_cpo
unlet s:save_cpo
