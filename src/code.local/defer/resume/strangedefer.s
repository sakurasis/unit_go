go : # command-line-arguments
所在位置 行:1 字符: 1
+ go build -gcflags="-l -S" strangedefer.go > strangedefer.s 2>&1
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : NotSpecified: (# command-line-arguments:String) [], RemoteException
    + FullyQualifiedErrorId : NativeCommandError
 
main.(*temp).Add STEXT size=116 args=0x10 locals=0x30 funcid=0x0 align=0x0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	TEXT	main.(*temp).Add(SB), ABIInternal, $48-16
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	CMPQ	SP, 16(R14)
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	PCDATA	$0, $-2
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	JLS	89
	0x0006 00006 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	PCDATA	$0, $-1
	0x0006 00006 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	SUBQ	$48, SP
	0x000a 00010 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	MOVQ	BP, 40(SP)
	0x000f 00015 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	LEAQ	40(SP), BP
	0x0014 00020 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	FUNCDATA	$0, gclocals路wBS4fiKwwXBG0Q3AcyXF/Q==(SB)
	0x0014 00020 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	FUNCDATA	$1, gclocals路5aa34RaZcmo0NkRpBHp2fg==(SB)
	0x0014 00020 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	FUNCDATA	$2, main.(*temp).Add.stkobj(SB)
	0x0014 00020 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	FUNCDATA	$5, main.(*temp).Add.arginfo1(SB)
	0x0014 00020 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	FUNCDATA	$6, main.(*temp).Add.argliveinfo(SB)
	0x0014 00020 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	PCDATA	$3, $1
	0x0014 00020 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	MOVUPS	X15, main..autotmp_3+24(SP)
	0x001a 00026 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	MOVQ	BX, AX
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	PCDATA	$1, $1
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	NOP
	0x0020 00032 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	CALL	runtime.convT64(SB)
	0x0025 00037 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	LEAQ	type.int(SB), CX
	0x002c 00044 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	MOVQ	CX, main..autotmp_3+24(SP)
	0x0031 00049 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	MOVQ	AX, main..autotmp_3+32(SP)
	0x0036 00054 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	LEAQ	main..autotmp_3+24(SP), AX
	0x003b 00059 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	MOVL	$1, BX
	0x0040 00064 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	MOVQ	BX, CX
	0x0043 00067 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	PCDATA	$1, $0
	0x0043 00067 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:8)	CALL	fmt.Println(SB)
	0x0048 00072 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:9)	LEAQ	runtime.zerobase(SB), AX
	0x004f 00079 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:9)	MOVQ	40(SP), BP
	0x0054 00084 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:9)	ADDQ	$48, SP
	0x0058 00088 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:9)	RET
	0x0059 00089 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:9)	NOP
	0x0059 00089 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	PCDATA	$1, $-1
	0x0059 00089 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	PCDATA	$0, $-2
	0x0059 00089 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	MOVQ	AX, 8(SP)
	0x005e 00094 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	MOVQ	BX, 16(SP)
	0x0063 00099 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0068 00104 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	MOVQ	8(SP), AX
	0x006d 00109 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	MOVQ	16(SP), BX
	0x0072 00114 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	PCDATA	$0, $-1
	0x0072 00114 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:7)	JMP	0
	0x0000 49 3b 66 10 76 53 48 83 ec 30 48 89 6c 24 28 48  I;f.vSH..0H.l$(H
	0x0010 8d 6c 24 28 44 0f 11 7c 24 18 48 89 d8 0f 1f 00  .l$(D..|$.H.....
	0x0020 e8 00 00 00 00 48 8d 0d 00 00 00 00 48 89 4c 24  .....H......H.L$
	0x0030 18 48 89 44 24 20 48 8d 44 24 18 bb 01 00 00 00  .H.D$ H.D$......
	0x0040 48 89 d9 e8 00 00 00 00 48 8d 05 00 00 00 00 48  H.......H......H
	0x0050 8b 6c 24 28 48 83 c4 30 c3 48 89 44 24 08 48 89  .l$(H..0.H.D$.H.
	0x0060 5c 24 10 e8 00 00 00 00 48 8b 44 24 08 48 8b 5c  \$......H.D$.H.\
	0x0070 24 10 eb 8c                                      $...
	rel 2+0 t=23 type.int+0
	rel 33+4 t=7 runtime.convT64+0
	rel 40+4 t=14 type.int+0
	rel 68+4 t=7 fmt.Println+0
	rel 75+4 t=14 runtime.zerobase+0
	rel 100+4 t=7 runtime.morestack_noctxt+0
main.main STEXT size=159 args=0x0 locals=0x38 funcid=0x0 align=0x0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	TEXT	main.main(SB), ABIInternal, $56-0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	CMPQ	SP, 16(R14)
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	PCDATA	$0, $-2
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	JLS	149
	0x000a 00010 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	PCDATA	$0, $-1
	0x000a 00010 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	SUBQ	$56, SP
	0x000e 00014 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	MOVQ	BP, 48(SP)
	0x0013 00019 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	LEAQ	48(SP), BP
	0x0018 00024 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	MOVQ	$0, R13
	0x001f 00031 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	MOVQ	R13, 40(SP)
	0x0024 00036 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	FUNCDATA	$0, gclocals路J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x0024 00036 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	FUNCDATA	$1, gclocals路wdmTuppZUxZYftR7OCq88Q==(SB)
	0x0024 00036 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	FUNCDATA	$2, main.main.stkobj(SB)
	0x0024 00036 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	FUNCDATA	$4, main.main.opendefer(SB)
	0x0024 00036 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	MOVB	$0, main..autotmp_7+23(SP)
	0x0029 00041 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	LEAQ	main..autotmp_4+23(SP), AX
	0x002e 00046 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVL	$1, BX
	0x0033 00051 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	PCDATA	$1, $1
	0x0033 00051 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	CALL	main.(*temp).Add(SB)
	0x0038 00056 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVUPS	X15, main..autotmp_6+24(SP)
	0x003e 00062 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	LEAQ	main.main.func1(SB), CX
	0x0045 00069 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVQ	CX, main..autotmp_6+24(SP)
	0x004a 00074 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVQ	AX, main..autotmp_6+32(SP)
	0x004f 00079 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	LEAQ	main..autotmp_6+24(SP), CX
	0x0054 00084 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVQ	CX, main..autotmp_8+40(SP)
	0x0059 00089 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVB	$1, main..autotmp_7+23(SP)
	0x005e 00094 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:15)	LEAQ	main..autotmp_4+23(SP), AX
	0x0063 00099 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:15)	MOVL	$3, BX
	0x0068 00104 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:15)	CALL	main.(*temp).Add(SB)
	0x006d 00109 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	MOVB	$0, main..autotmp_7+23(SP)
	0x0072 00114 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	MOVQ	main..autotmp_8+40(SP), DX
	0x0077 00119 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	MOVQ	(DX), CX
	0x007a 00122 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	CALL	CX
	0x007c 00124 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	MOVQ	48(SP), BP
	0x0081 00129 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	ADDQ	$56, SP
	0x0085 00133 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	RET
	0x0086 00134 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	CALL	runtime.deferreturn(SB)
	0x008b 00139 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	MOVQ	48(SP), BP
	0x0090 00144 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	ADDQ	$56, SP
	0x0094 00148 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	RET
	0x0095 00149 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:16)	NOP
	0x0095 00149 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	PCDATA	$1, $-1
	0x0095 00149 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	PCDATA	$0, $-2
	0x0095 00149 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	CALL	runtime.morestack_noctxt(SB)
	0x009a 00154 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	PCDATA	$0, $-1
	0x009a 00154 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:12)	JMP	0
	0x0000 49 3b 66 10 0f 86 8b 00 00 00 48 83 ec 38 48 89  I;f.......H..8H.
	0x0010 6c 24 30 48 8d 6c 24 30 49 c7 c5 00 00 00 00 4c  l$0H.l$0I......L
	0x0020 89 6c 24 28 c6 44 24 17 00 48 8d 44 24 17 bb 01  .l$(.D$..H.D$...
	0x0030 00 00 00 e8 00 00 00 00 44 0f 11 7c 24 18 48 8d  ........D..|$.H.
	0x0040 0d 00 00 00 00 48 89 4c 24 18 48 89 44 24 20 48  .....H.L$.H.D$ H
	0x0050 8d 4c 24 18 48 89 4c 24 28 c6 44 24 17 01 48 8d  .L$.H.L$(.D$..H.
	0x0060 44 24 17 bb 03 00 00 00 e8 00 00 00 00 c6 44 24  D$............D$
	0x0070 17 00 48 8b 54 24 28 48 8b 0a ff d1 48 8b 6c 24  ..H.T$(H....H.l$
	0x0080 30 48 83 c4 38 c3 e8 00 00 00 00 48 8b 6c 24 30  0H..8......H.l$0
	0x0090 48 83 c4 38 c3 e8 00 00 00 00 e9 61 ff ff ff     H..8.......a...
	rel 52+4 t=7 main.(*temp).Add+0
	rel 65+4 t=14 main.main.func1+0
	rel 105+4 t=7 main.(*temp).Add+0
	rel 122+0 t=10 +0
	rel 135+4 t=7 runtime.deferreturn+0
	rel 150+4 t=7 runtime.morestack_noctxt+0
main.main.func1 STEXT size=77 args=0x0 locals=0x18 funcid=0x15 align=0x0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	TEXT	main.main.func1(SB), WRAPPER|NEEDCTXT|ABIInternal, $24-0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	CMPQ	SP, 16(R14)
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	PCDATA	$0, $-2
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	JLS	53
	0x0006 00006 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	PCDATA	$0, $-1
	0x0006 00006 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	SUBQ	$24, SP
	0x000a 00010 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVQ	BP, 16(SP)
	0x000f 00015 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	LEAQ	16(SP), BP
	0x0014 00020 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVQ	32(R14), R12
	0x0018 00024 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	TESTQ	R12, R12
	0x001b 00027 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	JNE	60
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	NOP
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	FUNCDATA	$0, gclocals路g2BeySu+wFnoycgXfElmcg==(SB)
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	FUNCDATA	$1, gclocals路g2BeySu+wFnoycgXfElmcg==(SB)
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	FUNCDATA	$7, main.(*temp).Add.wrapinfo(SB)
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVQ	8(DX), AX
	0x0021 00033 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVL	$2, BX
	0x0026 00038 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	PCDATA	$1, $0
	0x0026 00038 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	CALL	main.(*temp).Add(SB)
	0x002b 00043 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVQ	16(SP), BP
	0x0030 00048 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	ADDQ	$24, SP
	0x0034 00052 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	RET
	0x0035 00053 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	NOP
	0x0035 00053 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	PCDATA	$1, $-1
	0x0035 00053 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	PCDATA	$0, $-2
	0x0035 00053 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	CALL	runtime.morestack(SB)
	0x003a 00058 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	PCDATA	$0, $-1
	0x003a 00058 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	JMP	0
	0x003c 00060 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	LEAQ	32(SP), R13
	0x0041 00065 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	CMPQ	(R12), R13
	0x0045 00069 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	JNE	29
	0x0047 00071 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	MOVQ	SP, (R12)
	0x004b 00075 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\resume\strangedefer.go:14)	JMP	29
	0x0000 49 3b 66 10 76 2f 48 83 ec 18 48 89 6c 24 10 48  I;f.v/H...H.l$.H
	0x0010 8d 6c 24 10 4d 8b 66 20 4d 85 e4 75 1f 48 8b 42  .l$.M.f M..u.H.B
	0x0020 08 bb 02 00 00 00 e8 00 00 00 00 48 8b 6c 24 10  ...........H.l$.
	0x0030 48 83 c4 18 c3 e8 00 00 00 00 eb c4 4c 8d 6c 24  H...........L.l$
	0x0040 20 4d 39 2c 24 75 d6 49 89 24 24 eb d0            M9,$u.I.$$..
	rel 39+4 t=7 main.(*temp).Add+0
	rel 54+4 t=7 runtime.morestack+0
go.cuinfo.producer.main SDWARFCUINFO dupok size=0
	0x0000 2d 6c 20 72 65 67 61 62 69                       -l regabi
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
main..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
runtime.memequal0路f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal0+0
runtime.memequal64路f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*main.temp- SRODATA dupok size=12
	0x0000 00 0a 2a 6d 61 69 6e 2e 74 65 6d 70              ..*main.temp
type..namedata.*func(*main.temp, int) *main.temp- SRODATA dupok size=35
	0x0000 00 21 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 74 65  .!*func(*main.te
	0x0010 6d 70 2c 20 69 6e 74 29 20 2a 6d 61 69 6e 2e 74  mp, int) *main.t
	0x0020 65 6d 70                                         emp
type.*func(*main.temp, int) *main.temp SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 b2 d4 91 33 08 08 08 36 00 00 00 00 00 00 00 00  ...3...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.temp, int) *main.temp-+0
	rel 48+8 t=1 type.func(*main.temp, int) *main.temp+0
type.func(*main.temp, int) *main.temp SRODATA dupok size=80
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 2f d7 b4 b8 02 08 08 33 00 00 00 00 00 00 00 00  /......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 02 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.temp, int) *main.temp-+0
	rel 44+4 t=-32763 type.*func(*main.temp, int) *main.temp+0
	rel 56+8 t=1 type.*main.temp+0
	rel 64+8 t=1 type.int+0
	rel 72+8 t=1 type.*main.temp+0
type..importpath.main. SRODATA dupok size=6
	0x0000 00 04 6d 61 69 6e                                ..main
type..namedata.Add. SRODATA dupok size=5
	0x0000 01 03 41 64 64                                   ..Add
type..namedata.*func(int) *main.temp- SRODATA dupok size=23
	0x0000 00 15 2a 66 75 6e 63 28 69 6e 74 29 20 2a 6d 61  ..*func(int) *ma
	0x0010 69 6e 2e 74 65 6d 70                             in.temp
type.*func(int) *main.temp SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9b 39 a9 23 08 08 08 36 00 00 00 00 00 00 00 00  .9.#...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int) *main.temp-+0
	rel 48+8 t=1 type.func(int) *main.temp+0
type.func(int) *main.temp SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 e3 2c ed 02 08 08 33 00 00 00 00 00 00 00 00  ..,....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int) *main.temp-+0
	rel 44+4 t=-32763 type.*func(int) *main.temp+0
	rel 56+8 t=1 type.int+0
	rel 64+8 t=1 type.*main.temp+0
type.*main.temp SRODATA size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 c3 2c 67 97 09 08 08 36 00 00 00 00 00 00 00 00  .,g....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 01 00 01 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.temp-+0
	rel 48+8 t=1 type.main.temp+0
	rel 56+4 t=5 type..importpath.main.+0
	rel 72+4 t=5 type..namedata.Add.+0
	rel 76+4 t=26 type.func(int) *main.temp+0
	rel 80+4 t=26 main.(*temp).Add+0
	rel 84+4 t=26 main.(*temp).Add+0
runtime.gcbits. SRODATA dupok size=0
type.main.temp SRODATA size=96
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 c5 ef dc 1e 0f 01 01 19 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.memequal0路f+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*main.temp-+0
	rel 44+4 t=5 type.*main.temp+0
	rel 56+8 t=1 type.main.temp+96
	rel 80+4 t=5 type..importpath.main.+0
runtime.nilinterequal路f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 3b fc f8 8f 08 08 08 36 00 00 00 00 00 00 00 00  ;......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 39 7a 09 0f 02 08 08 14 00 00 00 00 00 00 00 00  9z..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal路f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9d 9c 0e 59 08 08 08 36 00 00 00 00 00 00 00 00  ...Y...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 76 de 99 0d 02 08 08 17 00 00 00 00 00 00 00 00  v...............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=18
	0x0000 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65 20  ..*[1]interface 
	0x0010 7b 7d                                            {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a8 0e 57 36 08 08 08 36 00 00 00 00 00 00 00 00  ..W6...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 6e 20 6a 3d 02 08 08 11 00 00 00 00 00 00 00 00  n j=............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.nilinterequal路f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=-32763 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..namedata.*func()- SRODATA dupok size=9
	0x0000 00 07 2a 66 75 6e 63 28 29                       ..*func()
type.*func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 75 ac 29 27 08 08 08 36 00 00 00 00 00 00 00 00  u.)'...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 48+8 t=1 type.func()+0
type.func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 fe fa b9 80 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00                                      ....
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 44+4 t=-32763 type.*func()+0
type..namedata.*struct { F uintptr; .autotmp_2 *main.temp }- SRODATA dupok size=46
	0x0000 00 2c 2a 73 74 72 75 63 74 20 7b 20 46 20 75 69  .,*struct { F ui
	0x0010 6e 74 70 74 72 3b 20 2e 61 75 74 6f 74 6d 70 5f  ntptr; .autotmp_
	0x0020 32 20 2a 6d 61 69 6e 2e 74 65 6d 70 20 7d        2 *main.temp }
type.*struct { F uintptr; main..autotmp_2 *main.temp } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 12 fe cf 8f 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; .autotmp_2 *main.temp }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; main..autotmp_2 *main.temp }+0
type..namedata..F- SRODATA dupok size=4
	0x0000 00 02 2e 46                                      ...F
type..namedata..autotmp_2- SRODATA dupok size=12
	0x0000 00 0a 2e 61 75 74 6f 74 6d 70 5f 32              ...autotmp_2
type.noalg.struct { F uintptr; main..autotmp_2 *main.temp } SRODATA dupok size=128
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 4e d5 59 12 02 08 08 19 00 00 00 00 00 00 00 00  N.Y.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; .autotmp_2 *main.temp }-+0
	rel 44+4 t=-32763 type.*struct { F uintptr; main..autotmp_2 *main.temp }+0
	rel 48+8 t=1 type..importpath.main.+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; main..autotmp_2 *main.temp }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata..autotmp_2-+0
	rel 112+8 t=1 type.*main.temp+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals路wBS4fiKwwXBG0Q3AcyXF/Q== SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 00                    ..........
gclocals路5aa34RaZcmo0NkRpBHp2fg== SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 02                    ..........
main.(*temp).Add.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
main.(*temp).Add.arginfo1 SRODATA static dupok size=5
	0x0000 00 08 08 08 ff                                   .....
main.(*temp).Add.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals路J5F+7Qw7O7ve2QcWC7DpeQ== SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals路wdmTuppZUxZYftR7OCq88Q== SRODATA dupok size=10
	0x0000 02 00 00 00 03 00 00 00 00 04                    ..........
main.main.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 e8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
main.main.opendefer SRODATA dupok size=3
	0x0000 19 01 08                                         ...
gclocals路g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
main.(*temp).Add.wrapinfo SRODATA static dupok size=4
	0x0000 00 00 00 00                                      ....
	rel 0+4 t=5 main.(*temp).Add+0
