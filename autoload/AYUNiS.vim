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

let &cpo = s:save_cpo
unlet s:save_cpo
