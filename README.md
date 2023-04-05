# stool

有时用到一些工具，在网上找可能还没有自己撸的快，而且定制化还需要了解它的结构。不如大致学习下原理，简单撸一个。

simple tools write in go


# 1. a tcp echo tools 


```
➜  stool git:(main) ✗ ./stool tec -h
2023/04/05 21:55:33 start running stool
tcp echo client

Usage:
   tec [flags]

Flags:
  -h, --help        help for tec
  -i, --ip string   server addr
  -p, --port int    server port (default 8181)
```


./stool tec -i 127.0.0.1 -p 8182 -s afafasfsfsfsfasfsdf

