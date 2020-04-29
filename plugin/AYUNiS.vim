scriptencoding utf-8

if exists('g:loaded_AYUNiS')
  finish
endif
let g:loaded_AYUNiS = 1

let s:save_cpo = &cpo
set cpo&vim

function! s:RequireAYUNiS(host) abort
  return jobstart(['AYUNiS.nvim'], { 'rpc': v:true })
endfunction

call remote#host#Register('AYUNiS.nvim', '0', function('s:RequireAYUNiS'))
call remote#host#RegisterPlugin('AYUNiS.nvim', '0', [
\ {'type': 'function', 'name': 'AYUNiSGetNowPlaying', 'sync': 1, 'opts': {}},
\ {'type': 'function', 'name': 'AYUNiSNext', 'sync': 1, 'opts': {}},
\ {'type': 'function', 'name': 'AYUNiSPrev', 'sync': 1, 'opts': {}},
\ {'type': 'function', 'name': 'AYUNiSToggle', 'sync': 1, 'opts': {}},
\ {'type': 'function', 'name': 'AYUNiSToggleRepeat', 'sync': 1, 'opts': {}},
\ {'type': 'function', 'name': 'AYUNiSToggleShuffle', 'sync': 1, 'opts': {}},
\ {'type': 'function', 'name': 'AYUNiSVolumeDown', 'sync': 1, 'opts': {}},
\ {'type': 'function', 'name': 'AYUNiSVolumeUp', 'sync': 1, 'opts': {}},
\ {'type': 'function', 'name': 'InitializeAYUNiS', 'sync': 1, 'opts': {}},
\ ])
"
" variable (set init.vim)
if !exists('g:ayunis_rtp')
  let g:ayunis_rtp = ''
endif

" Initialize
call InitializeAYUNiS()

" Refresh tabline at regular interval
func! RefreshTabline(timer) abort
  " refresh tabline
  set tabline+=""
endfunc

let g:ayunis_refresh_tabline_timer = timer_start(
      \ 1000,
      \ 'RefreshTabline',
      \ {'repeat': -1})

let &cpo = s:save_cpo
unlet s:save_cpo
