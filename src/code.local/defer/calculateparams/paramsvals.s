go : # command-line-arguments
所在位置 行:1 字符: 1
+ go build -gcflags="-l -S" paramsvals.go > paramsvals.s 2>&1
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : NotSpecified: (# command-line-arguments:String) [], RemoteException
    + FullyQualifiedErrorId : NativeCommandError
 
main.deferInloop STEXT size=179 args=0x0 locals=0x28 funcid=0x0 align=0x0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	TEXT	main.deferInloop(SB), ABIInternal, $40-0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	CMPQ	SP, 16(R14)
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	PCDATA	$0, $-2
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	JLS	169
	0x000a 00010 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	PCDATA	$0, $-1
	0x000a 00010 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	SUBQ	$40, SP
	0x000e 00014 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	MOVQ	BP, 32(SP)
	0x0013 00019 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	LEAQ	32(SP), BP
	0x0018 00024 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	FUNCDATA	$0, gclocals路J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x0018 00024 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	FUNCDATA	$1, gclocals路CnDyI2HjYXFz19SsOj98tw==(SB)
	0x0018 00024 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	XORL	AX, AX
	0x001a 00026 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:6)	JMP	36
	0x001c 00028 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:6)	MOVQ	main.i+16(SP), AX
	0x0021 00033 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:6)	INCQ	AX
	0x0024 00036 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:6)	CMPQ	AX, $5
	0x0028 00040 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:6)	JGE	154
	0x002a 00042 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:6)	MOVQ	AX, main.i+16(SP)
	0x002f 00047 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$1, $0
	0x002f 00047 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	CALL	runtime.convT64(SB)
	0x0034 00052 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	AX, main..autotmp_4+24(SP)
	0x0039 00057 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	LEAQ	type.noalg.struct { F uintptr; main..autotmp_1 interface {} }(SB),
 AX
	0x0040 00064 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$1, $1
	0x0040 00064 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	CALL	runtime.newobject(SB)
	0x0045 00069 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	LEAQ	main.deferInloop.func1(SB), CX
	0x004c 00076 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	CX, (AX)
	0x004f 00079 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	LEAQ	type.int(SB), DX
	0x0056 00086 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	DX, 8(AX)
	0x005a 00090 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$0, $-2
	0x005a 00090 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	CMPL	runtime.writeBarrier(SB), $0
	0x0061 00097 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	JNE	110
	0x0063 00099 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	main..autotmp_4+24(SP), BX
	0x0068 00104 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	BX, 16(AX)
	0x006c 00108 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	JMP	124
	0x006e 00110 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	LEAQ	16(AX), DI
	0x0072 00114 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	main..autotmp_4+24(SP), BX
	0x0077 00119 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	CALL	runtime.gcWriteBarrierBX(SB)
	0x007c 00124 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$0, $-1
	0x007c 00124 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$1, $0
	0x007c 00124 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	NOP
	0x0080 00128 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	CALL	runtime.deferproc(SB)
	0x0085 00133 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	TESTL	AX, AX
	0x0087 00135 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	JNE	139
	0x0089 00137 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	JMP	28
	0x008b 00139 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	CALL	runtime.deferreturn(SB)
	0x0090 00144 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	32(SP), BP
	0x0095 00149 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	ADDQ	$40, SP
	0x0099 00153 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	RET
	0x009a 00154 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:9)	CALL	runtime.deferreturn(SB)
	0x009f 00159 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:9)	MOVQ	32(SP), BP
	0x00a4 00164 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:9)	ADDQ	$40, SP
	0x00a8 00168 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:9)	RET
	0x00a9 00169 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:9)	NOP
	0x00a9 00169 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	PCDATA	$1, $-1
	0x00a9 00169 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	PCDATA	$0, $-2
	0x00a9 00169 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x00ae 00174 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	PCDATA	$0, $-1
	0x00ae 00174 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:5)	JMP	0
	0x0000 49 3b 66 10 0f 86 9f 00 00 00 48 83 ec 28 48 89  I;f.......H..(H.
	0x0010 6c 24 20 48 8d 6c 24 20 31 c0 eb 08 48 8b 44 24  l$ H.l$ 1...H.D$
	0x0020 10 48 ff c0 48 83 f8 05 7d 70 48 89 44 24 10 e8  .H..H...}pH.D$..
	0x0030 00 00 00 00 48 89 44 24 18 48 8d 05 00 00 00 00  ....H.D$.H......
	0x0040 e8 00 00 00 00 48 8d 0d 00 00 00 00 48 89 08 48  .....H......H..H
	0x0050 8d 15 00 00 00 00 48 89 50 08 83 3d 00 00 00 00  ......H.P..=....
	0x0060 00 75 0b 48 8b 5c 24 18 48 89 58 10 eb 0e 48 8d  .u.H.\$.H.X...H.
	0x0070 78 10 48 8b 5c 24 18 e8 00 00 00 00 0f 1f 40 00  x.H.\$........@.
	0x0080 e8 00 00 00 00 85 c0 75 02 eb 91 e8 00 00 00 00  .......u........
	0x0090 48 8b 6c 24 20 48 83 c4 28 c3 e8 00 00 00 00 48  H.l$ H..(......H
	0x00a0 8b 6c 24 20 48 83 c4 28 c3 e8 00 00 00 00 e9 4d  .l$ H..(.......M
	0x00b0 ff ff ff                                         ...
	rel 3+0 t=23 type.int+0
	rel 48+4 t=7 runtime.convT64+0
	rel 60+4 t=14 type.noalg.struct { F uintptr; main..autotmp_1 interface {} }+0
	rel 65+4 t=7 runtime.newobject+0
	rel 72+4 t=14 main.deferInloop.func1+0
	rel 82+4 t=14 type.int+0
	rel 92+4 t=14 runtime.writeBarrier+-1
	rel 120+4 t=7 runtime.gcWriteBarrierBX+0
	rel 129+4 t=7 runtime.deferproc+0
	rel 140+4 t=7 runtime.deferreturn+0
	rel 155+4 t=7 runtime.deferreturn+0
	rel 170+4 t=7 runtime.morestack_noctxt+0
main.deferInloop.func1 STEXT size=108 args=0x0 locals=0x30 funcid=0x15 align=0x0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	TEXT	main.deferInloop.func1(SB), WRAPPER|NEEDCTXT|ABIInternal, $48-0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	CMPQ	SP, 16(R14)
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$0, $-2
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	JLS	81
	0x0006 00006 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$0, $-1
	0x0006 00006 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	SUBQ	$48, SP
	0x000a 00010 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	BP, 40(SP)
	0x000f 00015 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	LEAQ	40(SP), BP
	0x0014 00020 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	32(R14), R12
	0x0018 00024 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	TESTQ	R12, R12
	0x001b 00027 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	JNE	88
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	NOP
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	FUNCDATA	$0, gclocals路g2BeySu+wFnoycgXfElmcg==(SB)
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	FUNCDATA	$1, gclocals路EaPwxsZ75yY1hHMVZLmk6g==(SB)
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	FUNCDATA	$2, main.deferInloop.func1.stkobj(SB)
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	FUNCDATA	$7, fmt.Println.wrapinfo(SB)
	0x001d 00029 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	8(DX), SI
	0x0021 00033 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	16(DX), DX
	0x0025 00037 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVUPS	X15, main..autotmp_0+24(SP)
	0x002b 00043 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	SI, main..autotmp_0+24(SP)
	0x0030 00048 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	DX, main..autotmp_0+32(SP)
	0x0035 00053 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	LEAQ	main..autotmp_0+24(SP), AX
	0x003a 00058 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVL	$1, BX
	0x003f 00063 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	BX, CX
	0x0042 00066 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$1, $0
	0x0042 00066 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	CALL	fmt.Println(SB)
	0x0047 00071 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	40(SP), BP
	0x004c 00076 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	ADDQ	$48, SP
	0x0050 00080 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	RET
	0x0051 00081 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	NOP
	0x0051 00081 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$1, $-1
	0x0051 00081 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$0, $-2
	0x0051 00081 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	CALL	runtime.morestack(SB)
	0x0056 00086 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	PCDATA	$0, $-1
	0x0056 00086 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	JMP	0
	0x0058 00088 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	LEAQ	56(SP), R13
	0x005d 00093 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	NOP
	0x0060 00096 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	CMPQ	(R12), R13
	0x0064 00100 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	JNE	29
	0x0066 00102 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	MOVQ	SP, (R12)
	0x006a 00106 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:7)	JMP	29
	0x0000 49 3b 66 10 76 4b 48 83 ec 30 48 89 6c 24 28 48  I;f.vKH..0H.l$(H
	0x0010 8d 6c 24 28 4d 8b 66 20 4d 85 e4 75 3b 48 8b 72  .l$(M.f M..u;H.r
	0x0020 08 48 8b 52 10 44 0f 11 7c 24 18 48 89 74 24 18  .H.R.D..|$.H.t$.
	0x0030 48 89 54 24 20 48 8d 44 24 18 bb 01 00 00 00 48  H.T$ H.D$......H
	0x0040 89 d9 e8 00 00 00 00 48 8b 6c 24 28 48 83 c4 30  .......H.l$(H..0
	0x0050 c3 e8 00 00 00 00 eb a8 4c 8d 6c 24 38 0f 1f 00  ........L.l$8...
	0x0060 4d 39 2c 24 75 b7 49 89 24 24 eb b1              M9,$u.I.$$..
	rel 67+4 t=7 fmt.Println+0
	rel 82+4 t=7 runtime.morestack+0
main.main STEXT size=40 args=0x0 locals=0x8 funcid=0x0 align=0x0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	TEXT	main.main(SB), ABIInternal, $8-0
	0x0000 00000 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	CMPQ	SP, 16(R14)
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	PCDATA	$0, $-2
	0x0004 00004 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	JLS	33
	0x0006 00006 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	PCDATA	$0, $-1
	0x0006 00006 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	SUBQ	$8, SP
	0x000a 00010 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	MOVQ	BP, (SP)
	0x000e 00014 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	LEAQ	(SP), BP
	0x0012 00018 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	FUNCDATA	$0, gclocals路g2BeySu+wFnoycgXfElmcg==(SB)
	0x0012 00018 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	FUNCDATA	$1, gclocals路g2BeySu+wFnoycgXfElmcg==(SB)
	0x0012 00018 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:12)	PCDATA	$1, $0
	0x0012 00018 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:12)	CALL	main.deferInloop(SB)
	0x0017 00023 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:13)	MOVQ	(SP), BP
	0x001b 00027 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:13)	ADDQ	$8, SP
	0x001f 00031 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:13)	NOP
	0x0020 00032 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:13)	RET
	0x0021 00033 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:13)	NOP
	0x0021 00033 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	PCDATA	$1, $-1
	0x0021 00033 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	PCDATA	$0, $-2
	0x0021 00033 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	CALL	runtime.morestack_noctxt(SB)
	0x0026 00038 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	PCDATA	$0, $-1
	0x0026 00038 (C:\Users\sakura\Desktop\unit_go\src\code.local\defer\calculateparams\paramsvals.go:10)	JMP	0
	0x0000 49 3b 66 10 76 1b 48 83 ec 08 48 89 2c 24 48 8d  I;f.v.H...H.,$H.
	0x0010 2c 24 e8 00 00 00 00 48 8b 2c 24 48 83 c4 08 90  ,$.....H.,$H....
	0x0020 c3 e8 00 00 00 00 eb d8                          ........
	rel 19+4 t=7 main.deferInloop+0
	rel 34+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.producer.main SDWARFCUINFO dupok size=0
	0x0000 2d 6c 20 72 65 67 61 62 69                       -l regabi
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
main..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
runtime.nilinterequal路f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64路f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
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
type..namedata.*struct { F uintptr; .autotmp_1 interface {} }- SRODATA dupok size=48
	0x0000 00 2e 2a 73 74 72 75 63 74 20 7b 20 46 20 75 69  ..*struct { F ui
	0x0010 6e 74 70 74 72 3b 20 2e 61 75 74 6f 74 6d 70 5f  ntptr; .autotmp_
	0x0020 31 20 69 6e 74 65 72 66 61 63 65 20 7b 7d 20 7d  1 interface {} }
type.*struct { F uintptr; main..autotmp_1 interface {} } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9a bb 3d 25 08 08 08 36 00 00 00 00 00 00 00 00  ..=%...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; .autotmp_1 interface {} }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; main..autotmp_1 interface {} }+0
runtime.gcbits.04 SRODATA dupok size=1
	0x0000 04                                               .
type..importpath.main. SRODATA dupok size=6
	0x0000 00 04 6d 61 69 6e                                ..main
type..namedata..F- SRODATA dupok size=4
	0x0000 00 02 2e 46                                      ...F
type..namedata..autotmp_1- SRODATA dupok size=12
	0x0000 00 0a 2e 61 75 74 6f 74 6d 70 5f 31              ...autotmp_1
type.noalg.struct { F uintptr; main..autotmp_1 interface {} } SRODATA dupok size=128
	0x0000 18 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	0x0010 9c 4a 9b 71 02 08 08 19 00 00 00 00 00 00 00 00  .J.q............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.04+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; .autotmp_1 interface {} }-+0
	rel 44+4 t=-32763 type.*struct { F uintptr; main..autotmp_1 interface {} }+0
	rel 48+8 t=1 type..importpath.main.+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; main..autotmp_1 interface {} }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata..autotmp_1-+0
	rel 112+8 t=1 type.interface {}+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals路J5F+7Qw7O7ve2QcWC7DpeQ== SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals路CnDyI2HjYXFz19SsOj98tw== SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals路g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals路EaPwxsZ75yY1hHMVZLmk6g== SRODATA dupok size=9
	0x0000 01 00 00 00 02 00 00 00 00                       .........
main.deferInloop.func1.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
fmt.Println.wrapinfo SRODATA static dupok size=4
	0x0000 00 00 00 00                                      ....
	rel 0+4 t=5 fmt.Println+0
