# Manglet
Tool for quick mangle sensitive data in textish files

What it does is replace all letters by `x`, all digits by `0` and all symbols except `-,` by `-`. Makes this
```text
UserID,Date,SecretID
1032,2020-02-03,http:fyr.nrk?sv=7_2
6074,2020-08-12,cbfb7c1e-67ba-4430-bc0d-fd672517101d
```

looks like that 

```text
UserID,Date,SecretID
0000,0000-00-00,xxxx:xxx.xxx?xx=0_0
0000,0000-00-00,xxxx0x0x-00xx-0000-xx0x-xx000000000x
```

There are no other mangle scenarios yet.

You also can skip certain rows by specifying their numbers in `-skip` flag.


## Usage

```go
manglet -i "input.txt" -o "output.txt" -skip "1,3,59"
```


## Options

```text
-i string
   Input file path
-m string
   Mangle groups (default 'ld'):
       l = letters
       d = digits
       s = symbols
-o string
   Output file path. If empty, then input file path with suffix '_manglet' will be used instead
-v bool
   Verbose output (default false)
-skip string
   Rows to skip, separated by commas
```


