# Go 1.23.0 regression - panic in NewMethodSet

When using Go 1.22.4, the program prints:

```
processing type: type github.com/varungandhi-src/underpanic._Ctype_void [0]byte
processing type: type github.com/varungandhi-src/underpanic.MyCStructInGo struct{x github.com/varungandhi-src/underpanic._Ctype_int}
processing type: type _ struct{_ runtime/internal/sys.NotInHeap}
processing type: type github.com/varungandhi-src/underpanic._Ctype_int int32
processing type: type github.com/varungandhi-src/underpanic._Ctype_my_c_struct = github.com/varungandhi-src/underpanic._Ctype_struct___0
processing type: type github.com/varungandhi-src/underpanic._Ctype_struct___0 struct{x github.com/varungandhi-src/underpanic._Ctype_int}
```
