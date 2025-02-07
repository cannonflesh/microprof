# microprof

Экспортирует один метод, логирующий использование памяти и нагрузку на ЦПУ в точке вызова.

Пример использования - `example/main.go`

Все настройки осуществляются за счет передаваемых методу параметров, например:
```go
l := NewCustomLogger()

microprof.PrintProfilingInfo(l, microprof.UnitsKb, false)
```

В точке вызова функция измерит и распечатает в переданный логгер следующие данные:
```
INFO[0006] Allocated Memory: 8.8218 Mb                  
INFO[0006] Total Allocated Memory: 178.5753 Mb          
INFO[0006] Heap Memory: 8.8218 Mb                       
INFO[0006] Heap System Memory: 23.1562 Mb               
INFO[0006] Garbage Collector Memory: 2.9423 Mb          
INFO[0006] CPU usage: 7.8044
```

Если же в последнем параметре передать `true`, последняя строчка будет выглядеть так:
```
INFO[0000] CPU usage: 0: 0.0000 | 1: 0.0000 | 2: 0.0000 | 3: 0.0000 | 4: 0.0000 | 5: 100.0000 | 6: 0.0000 | 7: 0.0000
```
если в системе 8 виртуальных ядер.