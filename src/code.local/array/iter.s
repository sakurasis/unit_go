go : # command-line-arguments
所在位置 行:1 字符: 1
+ go build -gcflags="-l -S" iter.go > iter.s 2>&1
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : NotSpecified: (# command-line-arguments:String) [], RemoteException
    + FullyQualifiedErrorId : NativeCommandError
 
.\iter.go:18:18: invalid argument: index 5 out of bounds [0:5]
