// autogenerated from gen/generic.rules: do not edit!
// generated with: cd gen; go run *.go
package ssa

func rewriteValuegeneric(v *Value, config *Config) bool {
	b := v.Block
	switch v.Op {
	case OpAdd64:
		// match: (Add64 (Const64 [c]) (Const64 [d]))
		// cond:
		// result: (Const64 [c+d])
		{
			if v.Args[0].Op != OpConst64 {
				goto end8c46df6f85a11cb1d594076b0e467908
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst64 {
				goto end8c46df6f85a11cb1d594076b0e467908
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst64
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c + d
			return true
		}
		goto end8c46df6f85a11cb1d594076b0e467908
	end8c46df6f85a11cb1d594076b0e467908:
		;
	case OpAddPtr:
		// match: (AddPtr (ConstPtr [c]) (ConstPtr [d]))
		// cond:
		// result: (ConstPtr [c+d])
		{
			if v.Args[0].Op != OpConstPtr {
				goto end145c1aec793b2befff34bc8983b48a38
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConstPtr {
				goto end145c1aec793b2befff34bc8983b48a38
			}
			d := v.Args[1].AuxInt
			v.Op = OpConstPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c + d
			return true
		}
		goto end145c1aec793b2befff34bc8983b48a38
	end145c1aec793b2befff34bc8983b48a38:
		;
	case OpArrayIndex:
		// match: (ArrayIndex (Load ptr mem) idx)
		// cond:
		// result: (Load (PtrIndex <v.Type.PtrTo()> ptr idx) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end4894dd7b58383fee5f8a92be08437c33
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			idx := v.Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := b.NewValue0(v.Line, OpPtrIndex, TypeInvalid)
			v0.Type = v.Type.PtrTo()
			v0.AddArg(ptr)
			v0.AddArg(idx)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end4894dd7b58383fee5f8a92be08437c33
	end4894dd7b58383fee5f8a92be08437c33:
		;
	case OpCom16:
		// match: (Com16 (Com16 x))
		// cond:
		// result: (Copy x)
		{
			if v.Args[0].Op != OpCom16 {
				goto end388d572e5a72fd87a07da5cab243ebdc
			}
			x := v.Args[0].Args[0]
			v.Op = OpCopy
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto end388d572e5a72fd87a07da5cab243ebdc
	end388d572e5a72fd87a07da5cab243ebdc:
		;
	case OpCom32:
		// match: (Com32 (Com32 x))
		// cond:
		// result: (Copy x)
		{
			if v.Args[0].Op != OpCom32 {
				goto end5b2b3834acc7313649923604f685e7c5
			}
			x := v.Args[0].Args[0]
			v.Op = OpCopy
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto end5b2b3834acc7313649923604f685e7c5
	end5b2b3834acc7313649923604f685e7c5:
		;
	case OpCom64:
		// match: (Com64 (Com64 x))
		// cond:
		// result: (Copy x)
		{
			if v.Args[0].Op != OpCom64 {
				goto end6d6312f25d06a327d92f028b1ce50566
			}
			x := v.Args[0].Args[0]
			v.Op = OpCopy
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto end6d6312f25d06a327d92f028b1ce50566
	end6d6312f25d06a327d92f028b1ce50566:
		;
	case OpCom8:
		// match: (Com8 (Com8 x))
		// cond:
		// result: (Copy x)
		{
			if v.Args[0].Op != OpCom8 {
				goto end70cbd85c4b8e82c170dba7c23f8bc0f3
			}
			x := v.Args[0].Args[0]
			v.Op = OpCopy
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto end70cbd85c4b8e82c170dba7c23f8bc0f3
	end70cbd85c4b8e82c170dba7c23f8bc0f3:
		;
	case OpConstString:
		// match: (ConstString {s})
		// cond:
		// result: (StringMake (Addr <config.Frontend().TypeBytePtr()> {config.fe.StringData(s.(string))} (SB <config.Frontend().TypeUintptr()>)) (ConstPtr <config.Frontend().TypeUintptr()> [int64(len(s.(string)))]))
		{
			s := v.Aux
			v.Op = OpStringMake
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := b.NewValue0(v.Line, OpAddr, TypeInvalid)
			v0.Type = config.Frontend().TypeBytePtr()
			v0.Aux = config.fe.StringData(s.(string))
			v1 := b.NewValue0(v.Line, OpSB, TypeInvalid)
			v1.Type = config.Frontend().TypeUintptr()
			v0.AddArg(v1)
			v.AddArg(v0)
			v2 := b.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v2.Type = config.Frontend().TypeUintptr()
			v2.AuxInt = int64(len(s.(string)))
			v.AddArg(v2)
			return true
		}
		goto end68cc91679848c7c30bd8b0a8ed533843
	end68cc91679848c7c30bd8b0a8ed533843:
		;
	case OpEq16:
		// match: (Eq16 x x)
		// cond:
		// result: (ConstBool {true})
		{
			x := v.Args[0]
			if v.Args[1] != x {
				goto enda503589f9b617e708a5ad3ddb047809f
			}
			v.Op = OpConstBool
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = true
			return true
		}
		goto enda503589f9b617e708a5ad3ddb047809f
	enda503589f9b617e708a5ad3ddb047809f:
		;
	case OpEq32:
		// match: (Eq32 x x)
		// cond:
		// result: (ConstBool {true})
		{
			x := v.Args[0]
			if v.Args[1] != x {
				goto endc94ae3b97d0090257b02152e437b3e17
			}
			v.Op = OpConstBool
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = true
			return true
		}
		goto endc94ae3b97d0090257b02152e437b3e17
	endc94ae3b97d0090257b02152e437b3e17:
		;
	case OpEq64:
		// match: (Eq64 x x)
		// cond:
		// result: (ConstBool {true})
		{
			x := v.Args[0]
			if v.Args[1] != x {
				goto end4d21cead60174989467a9c8202dbb91d
			}
			v.Op = OpConstBool
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = true
			return true
		}
		goto end4d21cead60174989467a9c8202dbb91d
	end4d21cead60174989467a9c8202dbb91d:
		;
	case OpEq8:
		// match: (Eq8 x x)
		// cond:
		// result: (ConstBool {true})
		{
			x := v.Args[0]
			if v.Args[1] != x {
				goto end73dce8bba164e4f4a1dd701bf8cfb362
			}
			v.Op = OpConstBool
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = true
			return true
		}
		goto end73dce8bba164e4f4a1dd701bf8cfb362
	end73dce8bba164e4f4a1dd701bf8cfb362:
		;
	case OpEqFat:
		// match: (EqFat x y)
		// cond: x.Op == OpConstNil && y.Op != OpConstNil
		// result: (EqFat y x)
		{
			x := v.Args[0]
			y := v.Args[1]
			if !(x.Op == OpConstNil && y.Op != OpConstNil) {
				goto endcea7f7399afcff860c54d82230a9a934
			}
			v.Op = OpEqFat
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(y)
			v.AddArg(x)
			return true
		}
		goto endcea7f7399afcff860c54d82230a9a934
	endcea7f7399afcff860c54d82230a9a934:
		;
		// match: (EqFat (Load ptr mem) (ConstNil))
		// cond:
		// result: (EqPtr (Load <config.Frontend().TypeUintptr()> ptr mem) (ConstPtr <config.Frontend().TypeUintptr()> [0]))
		{
			if v.Args[0].Op != OpLoad {
				goto end540dc8dfbc66adcd3db2d7e819c534f6
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			if v.Args[1].Op != OpConstNil {
				goto end540dc8dfbc66adcd3db2d7e819c534f6
			}
			v.Op = OpEqPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := b.NewValue0(v.Line, OpLoad, TypeInvalid)
			v0.Type = config.Frontend().TypeUintptr()
			v0.AddArg(ptr)
			v0.AddArg(mem)
			v.AddArg(v0)
			v1 := b.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v1.Type = config.Frontend().TypeUintptr()
			v1.AuxInt = 0
			v.AddArg(v1)
			return true
		}
		goto end540dc8dfbc66adcd3db2d7e819c534f6
	end540dc8dfbc66adcd3db2d7e819c534f6:
		;
	case OpIsInBounds:
		// match: (IsInBounds (ConstPtr [c]) (ConstPtr [d]))
		// cond:
		// result: (ConstPtr {inBounds(c,d)})
		{
			if v.Args[0].Op != OpConstPtr {
				goto enddfd340bc7103ca323354aec96b113c23
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConstPtr {
				goto enddfd340bc7103ca323354aec96b113c23
			}
			d := v.Args[1].AuxInt
			v.Op = OpConstPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = inBounds(c, d)
			return true
		}
		goto enddfd340bc7103ca323354aec96b113c23
	enddfd340bc7103ca323354aec96b113c23:
		;
	case OpLoad:
		// match: (Load <t> ptr mem)
		// cond: t.IsString()
		// result: (StringMake (Load <config.Frontend().TypeBytePtr()> ptr mem) (Load <config.Frontend().TypeUintptr()> (OffPtr <config.Frontend().TypeBytePtr()> [config.PtrSize] ptr) mem))
		{
			t := v.Type
			ptr := v.Args[0]
			mem := v.Args[1]
			if !(t.IsString()) {
				goto end18afa4a6fdd6d0b92ed292840898c8f6
			}
			v.Op = OpStringMake
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := b.NewValue0(v.Line, OpLoad, TypeInvalid)
			v0.Type = config.Frontend().TypeBytePtr()
			v0.AddArg(ptr)
			v0.AddArg(mem)
			v.AddArg(v0)
			v1 := b.NewValue0(v.Line, OpLoad, TypeInvalid)
			v1.Type = config.Frontend().TypeUintptr()
			v2 := b.NewValue0(v.Line, OpOffPtr, TypeInvalid)
			v2.Type = config.Frontend().TypeBytePtr()
			v2.AuxInt = config.PtrSize
			v2.AddArg(ptr)
			v1.AddArg(v2)
			v1.AddArg(mem)
			v.AddArg(v1)
			return true
		}
		goto end18afa4a6fdd6d0b92ed292840898c8f6
	end18afa4a6fdd6d0b92ed292840898c8f6:
		;
	case OpMul64:
		// match: (Mul64 (Const64 [c]) (Const64 [d]))
		// cond:
		// result: (Const64 [c*d])
		{
			if v.Args[0].Op != OpConst64 {
				goto end7aea1048b5d1230974b97f17238380ae
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst64 {
				goto end7aea1048b5d1230974b97f17238380ae
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst64
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c * d
			return true
		}
		goto end7aea1048b5d1230974b97f17238380ae
	end7aea1048b5d1230974b97f17238380ae:
		;
	case OpMulPtr:
		// match: (MulPtr (ConstPtr [c]) (ConstPtr [d]))
		// cond:
		// result: (ConstPtr [c*d])
		{
			if v.Args[0].Op != OpConstPtr {
				goto end808c190f346658bb1ad032bf37a1059f
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConstPtr {
				goto end808c190f346658bb1ad032bf37a1059f
			}
			d := v.Args[1].AuxInt
			v.Op = OpConstPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c * d
			return true
		}
		goto end808c190f346658bb1ad032bf37a1059f
	end808c190f346658bb1ad032bf37a1059f:
		;
	case OpNeq16:
		// match: (Neq16 x x)
		// cond:
		// result: (ConstBool {false})
		{
			x := v.Args[0]
			if v.Args[1] != x {
				goto end192755dd3c2be992e9d3deb53794a8d2
			}
			v.Op = OpConstBool
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = false
			return true
		}
		goto end192755dd3c2be992e9d3deb53794a8d2
	end192755dd3c2be992e9d3deb53794a8d2:
		;
	case OpNeq32:
		// match: (Neq32 x x)
		// cond:
		// result: (ConstBool {false})
		{
			x := v.Args[0]
			if v.Args[1] != x {
				goto endeb23619fc85950a8df7b31126252c4dd
			}
			v.Op = OpConstBool
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = false
			return true
		}
		goto endeb23619fc85950a8df7b31126252c4dd
	endeb23619fc85950a8df7b31126252c4dd:
		;
	case OpNeq64:
		// match: (Neq64 x x)
		// cond:
		// result: (ConstBool {false})
		{
			x := v.Args[0]
			if v.Args[1] != x {
				goto endfc6eea780fb4056afb9e4287076da60c
			}
			v.Op = OpConstBool
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = false
			return true
		}
		goto endfc6eea780fb4056afb9e4287076da60c
	endfc6eea780fb4056afb9e4287076da60c:
		;
	case OpNeq8:
		// match: (Neq8 x x)
		// cond:
		// result: (ConstBool {false})
		{
			x := v.Args[0]
			if v.Args[1] != x {
				goto endcccf700d93c6d57765b80f92f7b3fa81
			}
			v.Op = OpConstBool
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = false
			return true
		}
		goto endcccf700d93c6d57765b80f92f7b3fa81
	endcccf700d93c6d57765b80f92f7b3fa81:
		;
	case OpNeqFat:
		// match: (NeqFat x y)
		// cond: x.Op == OpConstNil && y.Op != OpConstNil
		// result: (NeqFat y x)
		{
			x := v.Args[0]
			y := v.Args[1]
			if !(x.Op == OpConstNil && y.Op != OpConstNil) {
				goto end94c68f7dc30c66ed42e507e01c4e5dc7
			}
			v.Op = OpNeqFat
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(y)
			v.AddArg(x)
			return true
		}
		goto end94c68f7dc30c66ed42e507e01c4e5dc7
	end94c68f7dc30c66ed42e507e01c4e5dc7:
		;
		// match: (NeqFat (Load ptr mem) (ConstNil))
		// cond:
		// result: (NeqPtr (Load <config.Frontend().TypeUintptr()> ptr mem) (ConstPtr <config.Frontend().TypeUintptr()> [0]))
		{
			if v.Args[0].Op != OpLoad {
				goto end67d723bb0f39a5c897816abcf411e5cf
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			if v.Args[1].Op != OpConstNil {
				goto end67d723bb0f39a5c897816abcf411e5cf
			}
			v.Op = OpNeqPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := b.NewValue0(v.Line, OpLoad, TypeInvalid)
			v0.Type = config.Frontend().TypeUintptr()
			v0.AddArg(ptr)
			v0.AddArg(mem)
			v.AddArg(v0)
			v1 := b.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v1.Type = config.Frontend().TypeUintptr()
			v1.AuxInt = 0
			v.AddArg(v1)
			return true
		}
		goto end67d723bb0f39a5c897816abcf411e5cf
	end67d723bb0f39a5c897816abcf411e5cf:
		;
	case OpPtrIndex:
		// match: (PtrIndex <t> ptr idx)
		// cond:
		// result: (AddPtr ptr (MulPtr <config.Frontend().TypeUintptr()> idx (ConstPtr <config.Frontend().TypeUintptr()> [t.Elem().Size()])))
		{
			t := v.Type
			ptr := v.Args[0]
			idx := v.Args[1]
			v.Op = OpAddPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(ptr)
			v0 := b.NewValue0(v.Line, OpMulPtr, TypeInvalid)
			v0.Type = config.Frontend().TypeUintptr()
			v0.AddArg(idx)
			v1 := b.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v1.Type = config.Frontend().TypeUintptr()
			v1.AuxInt = t.Elem().Size()
			v0.AddArg(v1)
			v.AddArg(v0)
			return true
		}
		goto endf7546737f42c76a99699f241d41f491a
	endf7546737f42c76a99699f241d41f491a:
		;
	case OpSliceCap:
		// match: (SliceCap (Load ptr mem))
		// cond:
		// result: (Load (AddPtr <ptr.Type> ptr (ConstPtr <config.Frontend().TypeUintptr()> [config.PtrSize*2])) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end6696811bf6bd45e505d24c1a15c68e70
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := b.NewValue0(v.Line, OpAddPtr, TypeInvalid)
			v0.Type = ptr.Type
			v0.AddArg(ptr)
			v1 := b.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v1.Type = config.Frontend().TypeUintptr()
			v1.AuxInt = config.PtrSize * 2
			v0.AddArg(v1)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end6696811bf6bd45e505d24c1a15c68e70
	end6696811bf6bd45e505d24c1a15c68e70:
		;
	case OpSliceLen:
		// match: (SliceLen (Load ptr mem))
		// cond:
		// result: (Load (AddPtr <ptr.Type> ptr (ConstPtr <config.Frontend().TypeUintptr()> [config.PtrSize])) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end9844ce3e290e81355493141e653e37d5
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := b.NewValue0(v.Line, OpAddPtr, TypeInvalid)
			v0.Type = ptr.Type
			v0.AddArg(ptr)
			v1 := b.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v1.Type = config.Frontend().TypeUintptr()
			v1.AuxInt = config.PtrSize
			v0.AddArg(v1)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end9844ce3e290e81355493141e653e37d5
	end9844ce3e290e81355493141e653e37d5:
		;
	case OpSlicePtr:
		// match: (SlicePtr (Load ptr mem))
		// cond:
		// result: (Load ptr mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end459613b83f95b65729d45c2ed663a153
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		goto end459613b83f95b65729d45c2ed663a153
	end459613b83f95b65729d45c2ed663a153:
		;
	case OpStore:
		// match: (Store dst (Load <t> src mem) mem)
		// cond: t.Size() > 8
		// result: (Move [t.Size()] dst src mem)
		{
			dst := v.Args[0]
			if v.Args[1].Op != OpLoad {
				goto end324ffb6d2771808da4267f62c854e9c8
			}
			t := v.Args[1].Type
			src := v.Args[1].Args[0]
			mem := v.Args[1].Args[1]
			if v.Args[2] != mem {
				goto end324ffb6d2771808da4267f62c854e9c8
			}
			if !(t.Size() > 8) {
				goto end324ffb6d2771808da4267f62c854e9c8
			}
			v.Op = OpMove
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = t.Size()
			v.AddArg(dst)
			v.AddArg(src)
			v.AddArg(mem)
			return true
		}
		goto end324ffb6d2771808da4267f62c854e9c8
	end324ffb6d2771808da4267f62c854e9c8:
		;
		// match: (Store dst str mem)
		// cond: str.Type.IsString()
		// result: (Store (OffPtr <config.Frontend().TypeBytePtr()> [config.PtrSize] dst) (StringLen <config.Frontend().TypeUintptr()> str) (Store <TypeMem> dst (StringPtr <config.Frontend().TypeBytePtr()> str) mem))
		{
			dst := v.Args[0]
			str := v.Args[1]
			mem := v.Args[2]
			if !(str.Type.IsString()) {
				goto enddf0c5a150f4b4bf6715fd2bd4bb4cc20
			}
			v.Op = OpStore
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := b.NewValue0(v.Line, OpOffPtr, TypeInvalid)
			v0.Type = config.Frontend().TypeBytePtr()
			v0.AuxInt = config.PtrSize
			v0.AddArg(dst)
			v.AddArg(v0)
			v1 := b.NewValue0(v.Line, OpStringLen, TypeInvalid)
			v1.Type = config.Frontend().TypeUintptr()
			v1.AddArg(str)
			v.AddArg(v1)
			v2 := b.NewValue0(v.Line, OpStore, TypeInvalid)
			v2.Type = TypeMem
			v2.AddArg(dst)
			v3 := b.NewValue0(v.Line, OpStringPtr, TypeInvalid)
			v3.Type = config.Frontend().TypeBytePtr()
			v3.AddArg(str)
			v2.AddArg(v3)
			v2.AddArg(mem)
			v.AddArg(v2)
			return true
		}
		goto enddf0c5a150f4b4bf6715fd2bd4bb4cc20
	enddf0c5a150f4b4bf6715fd2bd4bb4cc20:
		;
	case OpStringLen:
		// match: (StringLen (StringMake _ len))
		// cond:
		// result: len
		{
			if v.Args[0].Op != OpStringMake {
				goto end0d922460b7e5ca88324034f4bd6c027c
			}
			len := v.Args[0].Args[1]
			v.Op = len.Op
			v.AuxInt = len.AuxInt
			v.Aux = len.Aux
			v.resetArgs()
			v.AddArgs(len.Args...)
			return true
		}
		goto end0d922460b7e5ca88324034f4bd6c027c
	end0d922460b7e5ca88324034f4bd6c027c:
		;
	case OpStringPtr:
		// match: (StringPtr (StringMake ptr _))
		// cond:
		// result: ptr
		{
			if v.Args[0].Op != OpStringMake {
				goto end061edc5d85c73ad909089af2556d9380
			}
			ptr := v.Args[0].Args[0]
			v.Op = ptr.Op
			v.AuxInt = ptr.AuxInt
			v.Aux = ptr.Aux
			v.resetArgs()
			v.AddArgs(ptr.Args...)
			return true
		}
		goto end061edc5d85c73ad909089af2556d9380
	end061edc5d85c73ad909089af2556d9380:
		;
	case OpStructSelect:
		// match: (StructSelect [idx] (Load ptr mem))
		// cond:
		// result: (Load (OffPtr <v.Type.PtrTo()> [idx] ptr) mem)
		{
			idx := v.AuxInt
			if v.Args[0].Op != OpLoad {
				goto end16fdb45e1dd08feb36e3cc3fb5ed8935
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := b.NewValue0(v.Line, OpOffPtr, TypeInvalid)
			v0.Type = v.Type.PtrTo()
			v0.AuxInt = idx
			v0.AddArg(ptr)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end16fdb45e1dd08feb36e3cc3fb5ed8935
	end16fdb45e1dd08feb36e3cc3fb5ed8935:
	}
	return false
}
func rewriteBlockgeneric(b *Block) bool {
	switch b.Kind {
	case BlockIf:
		// match: (If (Not cond) yes no)
		// cond:
		// result: (If cond no yes)
		{
			v := b.Control
			if v.Op != OpNot {
				goto endebe19c1c3c3bec068cdb2dd29ef57f96
			}
			cond := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockIf
			b.Control = cond
			b.Succs[0] = no
			b.Succs[1] = yes
			return true
		}
		goto endebe19c1c3c3bec068cdb2dd29ef57f96
	endebe19c1c3c3bec068cdb2dd29ef57f96:
		;
		// match: (If (ConstBool {c}) yes no)
		// cond: c.(bool)
		// result: (Plain nil yes)
		{
			v := b.Control
			if v.Op != OpConstBool {
				goto end9ff0273f9b1657f4afc287562ca889f0
			}
			c := v.Aux
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c.(bool)) {
				goto end9ff0273f9b1657f4afc287562ca889f0
			}
			b.Func.removePredecessor(b, no)
			b.Kind = BlockPlain
			b.Control = nil
			b.Succs = b.Succs[:1]
			b.Succs[0] = yes
			return true
		}
		goto end9ff0273f9b1657f4afc287562ca889f0
	end9ff0273f9b1657f4afc287562ca889f0:
		;
		// match: (If (ConstBool {c}) yes no)
		// cond: !c.(bool)
		// result: (Plain nil no)
		{
			v := b.Control
			if v.Op != OpConstBool {
				goto endf401a4553c3c7c6bed64801da7bba076
			}
			c := v.Aux
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(!c.(bool)) {
				goto endf401a4553c3c7c6bed64801da7bba076
			}
			b.Func.removePredecessor(b, yes)
			b.Kind = BlockPlain
			b.Control = nil
			b.Succs = b.Succs[:1]
			b.Succs[0] = no
			return true
		}
		goto endf401a4553c3c7c6bed64801da7bba076
	endf401a4553c3c7c6bed64801da7bba076:
	}
	return false
}
