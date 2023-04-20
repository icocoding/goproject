
## 获取当前目录
`os.Getwd()`

## 修改工作目录
`os.Chdir()`

## 读文件
`os.ReadFile()`

## 写文件
`os.WriteFile()`

## 权限
- os.ModePerm : 最高权限 777
- os.FileMode(644).String()
- 只读模式：os.O_RDONLY
- 只写模式：os.O_WRONLY
- 读写模式：os.O_RDWR

更多的操作模式：
- os.O_APPEDN ：写内容是，追加在现有内容的后面；
- os.O_CREATE ：当给定路径上的文件不存在时，创建一个新文件；
- os.O_EXCL：需要和os.O_CREATE一同使用，表示给定的路径上不能有存在的文件。
- os.O_SYNC ：在打开文件之上实施同步I/O。它会保证读写的内容总会与硬盘上的数据保持同步。
- os.O_TRUNC：如果文件已存在，并且是常规的文件，那么就先清空其中已存在的任何内容。

## 创建文件
`os.Create()`

## 创建目录
`os.MkDirAll()`

`os.MkDir()`

## 删除目录
`os.RemoveAll()`

## 删除文件
`os.Remove()`

## 临时目录
`os.TempDir()`

## 重命名
`os.Rename()`