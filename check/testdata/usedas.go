package foo

func FuncAsParam(fn func(FooType) string) { fn(0) }

func PassedAsParam(f FooType) string {
	doWork()
	return "foo"
}

func (f FooType) FuncAsParam2(fn func(FooType) []byte) { fn(0) }

func PassedAsParam2(f FooType) []byte {
	doWork()
	return nil
}

func AnonType() {
	for _, f := range []func(FooType, int32){
		func(f FooType, i int32) {
			doWork()
			println(i)
		},
	} {
		f(1, 2)
	}
	for _, f := range []struct {
		f2 func(f FooType, i int64)
	}{
		{f2: func(f FooType, i int64) {
			doWork()
			println(i)
		}},
	} {
		f.f2(3, 4)
	}
}

func UsedAsArg() {
	foo := func(f func(f FooType, u uint32)) {
		f(5, 6)
	}
	bar := func(v interface{}) {
		doWork()
		println(v)
	}
	foo(func(f FooType, u uint32) {
		println(f)
	})
	bar(func(f FooType, u uint64) {
		println(f)
	})
}

func globalParam(f func(f FooType, i int8)) {
	f(7, 8)
}

func UsedAsGlobalArg(f FooType, i int8) {
	doWork()
	println(f)
}

func globalParamIface(v interface{}) {
	println(v)
}

func UsedAsGlobalArgIface(f FooType, i int16) {
	doWork()
	println(f)
}

func GlobArgUse() {
	globalParam(UsedAsGlobalArg)
	globalParamIface(UsedAsGlobalArgIface)
}

type barIface interface {
	bar(FooType, uint16)
}

type barType struct{}

func (b *barType) bar(f FooType, u uint16) {
	doWork()
	println(f)
}

func barImpl() barIface { return &barType{} }

func BarIfaceUse() {
	b := barImpl()
	b.bar(0, 1)
}

func (f FooType) MethodPassedAsParam(f2 FooType) bool {
	if f == 3 {
		doWork()
		return true
	}
	return true
}

func (f FooType) MethodPassedAsParam2() bool {
        if f == 4 {
                doWork()
                return true
        }
        return true
}

func MethodUsedAsArg() {
	foo := func(f func(f FooType) bool) {
		f(2)
	}
	var f FooType
	foo(f.MethodPassedAsParam)
	foo((FooType).MethodPassedAsParam2)
}