# Go 1.23.0 regression - panic in NewMethodSet

When using Go 1.22.4 (SHA: 82f8fb218211493f069cfa18b38892a03932a134), the program prints:

```
processing type: type github.com/varungandhi-src/underpanic._Ctype_void [0]byte
processing type: type github.com/varungandhi-src/underpanic.MyCStructInGo struct{x github.com/varungandhi-src/underpanic._Ctype_int}
processing type: type _ struct{_ runtime/internal/sys.NotInHeap}
processing type: type github.com/varungandhi-src/underpanic._Ctype_int int32
processing type: type github.com/varungandhi-src/underpanic._Ctype_my_c_struct = github.com/varungandhi-src/underpanic._Ctype_struct___0
processing type: type github.com/varungandhi-src/underpanic._Ctype_struct___0 struct{x github.com/varungandhi-src/underpanic._Ctype_int}
```

When using Go 1.23.0 (updating `go.mod`), the program prints:

```
processing type: type github.com/varungandhi-src/underpanic._Ctype_int int32
processing type: type github.com/varungandhi-src/underpanic._Ctype_struct___0 struct{x github.com/varungandhi-src/underpanic._Ctype_int}
panic: Named.check == nil but type is incomplete

goroutine 1 [running]:
go/types.(*Named).under(0x140001a6850)
        /Users/varun/.asdf/installs/golang/1.23.0/go/src/go/types/named.go:548 +0x38c
go/types.under({0x1007ca020, 0x140001a6850})
        /Users/varun/.asdf/installs/golang/1.23.0/go/src/go/types/under.go:16 +0x2c
go/types.writeObject(0x14000467890, {0x1007cdb38, 0x140001a4780}, 0x0)
        /Users/varun/.asdf/installs/golang/1.23.0/go/src/go/types/object.go:576 +0x734
go/types.ObjectString({0x1007cdb38, 0x140001a4780}, 0x0)
        /Users/varun/.asdf/installs/golang/1.23.0/go/src/go/types/object.go:613 +0x44
go/types.(*TypeName).String(...)
        /Users/varun/.asdf/installs/golang/1.23.0/go/src/go/types/object.go:619
main.main()
        /Users/varun/Code/play/go/underpanic/main.go:33 +0x164
```
