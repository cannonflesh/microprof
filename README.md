# microprof

Экспортирует один метод, логирующий использование памяти и нагрузку на ЦПУ в точке вызова.

Пример использования - `example/main.go`

Все настройки осуществляются за счет передаваемых методу параметров, например:
```go
l := NewCustomLogger(opt...)

microprof.PrintProfilingInfo(l, microprof.UnitsKb, false)
```

В точке вызова функция измерит и распечатает в переданный логгер следующие данные:
```
INFO[0000] Allocated Memory: 0.5215 Kb                  
INFO[0000] Total Allocated Memory: 7.2755 Kb            
INFO[0000] Heap Memory: 0.5215 Kb                       
INFO[0000] Heap System Memory: 7.5312 Kb                
INFO[0000] Garbage Collector Memory: 2.2604 Kb          
INFO[0000] CPU usage: 0.0000 
```

Если же в последнем параметре передать `true`, последняя строчка будет выглядеть так:
```
INFO[0000] CPU usage: 0: 0.0000 | 1: 0.0000 | 2: 0.0000 | 3: 0.0000 | 4: 0.0000 | 5: 100.0000 | 6: 0.0000 | 7: 0.0000
```
если в системе 8 виртуальных ядер.