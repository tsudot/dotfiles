Vim�UnDo� 	Ҩy�7������O^��w(+��iK��   N   "nnoremap <C-e> :NERDTreeToggle<CR>   ,                           U-��    _�                             ����                                                                                                                                                                                                                                                                                                                                                             U �    �   I   K          +let g:go_highlight_build_constraints = 1   �   H   J          �let g:go_highlight_operators = 1                                                                                                                                                                            �   G   I          �let g:go_highlight_structs = 1                                                                                                                                                                              �   F   H          �let g:go_highlight_methods = 1                                                                                                                                                                              �   E   G          �let g:go_highlight_functions = 1                                                                                                                                                                            �   D   F          �let g:go_fmt_command = "goimports"                                                                                                                                                                          �   C   E          �" Go specific                                                                                                                                                                                               �   B   D          �                                                                                                                                                                                                            �   A   C          �highlight LineNr ctermfg=cyan                                                                                                                                                                               �   @   B          �                                                                                                                                                                                                            �   ?   A          �                                                                                                                                                                                                            �   >   @          �let g:airline#extensions#tabline#enabled = 1                                                                                                                                                                �   =   ?          �                                                                                                                                                                                                            �   <   >          �set wildignore+=*/tmp/*,*.so,*.swp,*.zip                                                                                                                                                                    �   ;   =          �let g:ctrlp_working_path_mode = 'ra'                                                                                                                                                                        �   :   <          �let g:ctrlp_cmd = 'CtrlP'                                                                                                                                                                                   �   9   ;          �let g:ctrlp_map = '<c-p>'                                                                                                                                                                                   �   8   :          �                                                                                                                                                                                                            �   7   9          �set runtimepath^=~/.vim/bundle/ctrlp.vim                                                                                                                                                                    �   6   8          �                                                                                                                                                                                                            �   5   7          �set backspace=indent,eol,start                                                                                                                                                                              �   4   6          �" Fix for backspace                                                                                                                                                                                         �   3   5          �                                                                                                                                                                                                            �   2   4          �" let g:pymode_lint_ignore = "W0611"                                                                                                                                                                        �   1   3          �" Ignore unused imports pyflakes - pymode                                                                                                                                                                   �   /   1          !map <C-S-Tab> :bprev<CR>         �   .   0          �map <C-Tab> :bnext<CR>                                                                                                                                                                                      �   -   /          �" To switch between buffers                                                                                                                                                                                 �   ,   .          �                                                                                                                                                                                                            �   +   -          �nnoremap <C-e> :NERDTreeToggle<CR>                                                                                                                                                                          �   *   ,          �" NerdTree Toggle                                                                                                                                                                                           �   )   +          �                                                                                                                                                                                                            �   (   *          �autocmd vimenter * if !argc() | NERDTree | endif                                                                                                                                                            �   '   )          �" Open NERDTree on startup if no files are selected                                                                                                                                                         �   &   (          �                                                                                                                                                                                                            �   %   '          �autocmd bufenter * if (winnr("$") == 1 && exists("b:NERDTreeType") && b:NERDTreeType == "primary") | q | endif                                                                                              �   $   &          �" Close window if only nerdtree is open                                                                                                                                                                     �   #   %          �                                                                                                                                                                                                            �   "   $          �map <C-n> :NERDTreeToggle<CR>                                                                                                                                                                               �   !   #          �" NerdTree                                                                                                                                                                                                  �       "          �                                                                                                                                                                                                            �      !          �set undofile                                                                                                                                                                                                �                 �set undodir=$HOME/.vim/undo                                                                                                                                                                                 �                �                                                                                                                                                                                                            �                �filetype indent on                                                                                                                                                                                          �                �" OPTIONAL: This enables automatic indentation as you type.                                                                                                                                                 �                �                                                                                                                                                                                                            �                �filetype plugin on                                                                                                                                                                                          �                �                                                                                                                                                                                                            �                �set cursorline                                                                                                                                                                                              �                �set number                                                                                                                                                                                                  �                �" line numbers                                                                                                                                                                                              �                �                                                                                                                                                                                                            �                �colorscheme molokai                                                                                                                                                                                         �                �syntax on                                                                                                                                                                                                   �                �"enable syntax highlighting                                                                                                                                                                                 �                �set t_Co=256                                                                                                                                                                                                �                �set bg=dark                                                                                                                                                                                                 �                �set guifont=Monaco\ 10                                                                                                                                                                                      �                �" set font to monaco                                                                                                                                                                                        �                �set secure                                                                                                                                                                                                  �                �set linebreak                                                                                                                                                                                               �   
             �set wrap                                                                                                                                                                                                    �   	             �set sm " showmatch                                                                                                                                                                                          �      
          �set sta " smarttab                                                                                                                                                                                          �      	          �set ai " autoindent                                                                                                                                                                                         �                �set et " expandtab                                                                                                                                                                                          �                �set sw=4 " shiftwidth = 4                                                                                                                                                                                   �                �set ts=4 " tabstop = 4                                                                                                                                                                                      �                �                                                                                                                                                                                                            �                �call pathogen#helptags()                                                                                                                                                                                    �                �call pathogen#infect()                                                                                                                                                                                      �                 �filetype off                                                                                                                                                                                                5�_�                   J        ����                                                                                                                                                                                                                                                                                                                                                             U-�     �   J               �   J            5�_�                    M        ����                                                                                                                                                                                                                                                                                                                                                             U-�     �   L   N   N      "5�_�                    M        ����                                                                                                                                                                                                                                                                                                                                                             U-�     �   L   N   N       5�_�                    N        ����                                                                                                                                                                                                                                                                                                                                                             U-�     �   M              "nmap <F8> :TagbarToggle<CR>5�_�      	              N       ����                                                                                                                                                                                                                                                                                                                                                             U-�     �   M              nmap <F8> :TagbarToggle<CR>5�_�      
           	   N   
    ����                                                                                                                                                                                                                                                                                                                                                             U-�     �   M              nmap <C-F8>> :TagbarToggle<CR>5�_�   	              
   N   
    ����                                                                                                                                                                                                                                                                                                                                                             U-��     �   M              nmap <C-F8> :TagbarToggle<CR>5�_�   
                 N   	    ����                                                                                                                                                                                                                                                                                                                                                             U-��     �   M              nmap <C-e>> :TagbarToggle<CR>5�_�                     ,        ����                                                                                                                                                                                                                                                                                                                                                             U-��    �   +   -   N      "nnoremap <C-e> :NERDTreeToggle<CR>5�_�                          ����                                                                                                                                                                                                                                                                                                                                                             U!�    �         J      set guifont=Monaco\ 95�_�                            ����                                                                                                                                                                                                                                                                                                                                                             U!�    �         J      set guifont=Monaco\ 45��