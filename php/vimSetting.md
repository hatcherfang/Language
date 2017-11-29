一般的检查方式是回到命令行运行php -l，或者升级一下直接在Vim里运行:!php -l %，或者再将此绑定为快捷键，  
如:map <F5> :!php -l %<CR>，每当按F5便可立即执行语法检查。  
```
" 基于缩进或语法进行代码折叠
set foldmethod=indent
"set foldmethod=syntax
" 设置状态栏主题风格
let g:Powerline_colorscheme='solarized256'

"自动对齐       
set ai
"依据上面的对起格式
set si
set smarttab
set wrap
set lbr
set tw=0
" 从第二层开始可视化显示缩进
let g:indent_guides_start_level=2
" 色块宽度
let g:indent_guides_guide_size=1

" vim 之自动缩进(smartindent) tab 空格数设置为4
set smartindent  
set tabstop=4  
set shiftwidth=4  
set expandtab  
set softtabstop=4  
```
## Reference  
[PHPSyntaxCheck : PHP语法检查](http://www.vim.org/scripts/script.php?script_id=4984)
