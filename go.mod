module gitee.com/quant1x/gotdx

go 1.21.1

require (
	gitee.com/quant1x/gox v1.15.1
	gitee.com/quant1x/pkg v0.1.3
	golang.org/x/exp v0.0.0-20231127185646-65229373498e
	golang.org/x/text v0.14.0
)

//replace gitee.com/quant1x/gox v1.15.1 => ../../mymmsc/gox
//replace gitee.com/quant1x/pkg v0.1.3 => ../pkg

require (
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/google/pprof v0.0.0-20231203200248-ad67f76aa53d // indirect
	golang.org/x/sys v0.15.0 // indirect
)
