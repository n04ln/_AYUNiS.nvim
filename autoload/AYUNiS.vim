scriptencoding utf-8

if !exists('g:loaded_AYUNiS')
    finish
endif
let g:loaded_AYUNiS = 1

let s:save_cpo = &cpo
set cpo&vim

" 

let &cpo = s:save_cpo
unlet s:save_cpo
