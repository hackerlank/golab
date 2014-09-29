go test -bench=".*"


go test -memprofile mem.out -test.memprofilerate=1
go tool pprof memory.test mem.out --alloc_space


go test -cpuprofile cpu.out
go tool pprof memory.test cpu.out

cpu消耗在runtime.mach_semaphore_wait， 有点奇怪