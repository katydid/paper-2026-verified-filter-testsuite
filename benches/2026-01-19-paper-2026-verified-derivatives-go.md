# 2026-01-19

Benchmark ran using:
* paper-2026-verified-derivatives-go: e8569912be1d3829abf9368680437e5c1e80146e
* paper-2026-verified-derivatives-testsuite: 623a2dae55d609cd6567495681f1aaa80750cd23

```
% make paper_benchmarks
go test -v -run=none -bench=BenchmarkSuite/Conf ./testsuite/intern/
goos: darwin
goarch: arm64
pkg: github.com/katydid/paper-2026-verified-derivatives/testsuite/intern
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                  317000              3763 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14               215257              5176 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14          64069             18547 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                            313927              3599 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14                 256488              4391 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14            70022             16644 ns/op
PASS
ok      github.com/katydid/paper-2026-verified-derivatives/testsuite/intern     8.264s
go test -v -run=none -bench=BenchmarkSuite/Conf -benchtime=100x ./testsuite/mem/
goos: darwin
goarch: arm64
pkg: github.com/katydid/paper-2026-verified-derivatives/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                     100               643.3 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14                  100              1302 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14            100              1856 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                               100               403.8 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14                    100               449.2 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14              100              1091 ns/op
PASS
ok      github.com/katydid/paper-2026-verified-derivatives/testsuite/mem        0.656s
go test -v -run=none -bench=BenchmarkSuite/Conf -benchtime=1000x ./testsuite/mem/
goos: darwin
goarch: arm64
pkg: github.com/katydid/paper-2026-verified-derivatives/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                    1000               410.5 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14                 1000              1051 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14           1000              1100 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                              1000               200.3 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14                   1000               239.1 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14             1000               508.9 ns/op
PASS
ok      github.com/katydid/paper-2026-verified-derivatives/testsuite/mem        0.652s
go test -v -run=none -bench=BenchmarkSuite/Conf -benchtime=10000x ./testsuite/mem/
goos: darwin
goarch: arm64
pkg: github.com/katydid/paper-2026-verified-derivatives/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                   10000               391.1 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14                10000              1064 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14          10000              1101 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                             10000               191.4 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14                  10000               221.5 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14            10000               342.8 ns/op
PASS
ok      github.com/katydid/paper-2026-verified-derivatives/testsuite/mem        0.682s
go test -v -run=none -bench=BenchmarkSuite/Conf ./testsuite/auto/
goos: darwin
goarch: arm64
pkg: github.com/katydid/paper-2026-verified-derivatives/testsuite/auto
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                 3743166               321.7 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14              1221672               980.9 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14        1225294               979.2 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                          10845369               109.4 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14                7600720               154.4 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14          4572258               261.9 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/katydid/paper-2026-verified-derivatives/testsuite/auto       10.721s
```