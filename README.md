# calcular_PI_threads

como executar:
    <br /> vá para o repósitorio para onde o arquivo esta localizado e digite no terminal
    ```
    go run main.go
    ```
    <br /> 
    Abaixo estão os resultados com os paramentros pedidos pelo professor.
    <br />
 

- 32 Threads
```
Numero de termos 1000000000
Numero de threads: 32
Tempo Médio -> 28.613925040000005
Desvio padrão ->  0.49692319666389334
Coeficiente de variação ->  1.7366481388667718
Valores obtidos para Pi ->  [3.141592652590108 3.141592652590108 3.141592652590108 3.141592652590108 3.141592652590108]
```

- 16 Threads
```
Numero de termos 1000000000
Numero de threads: 16
Tempo Médio -> 29.10392024
Desvio padrão ->  0.8128574811588379
Coeficiente de variação ->  2.792948422259825
Valores obtidos para Pi ->  [3.141592652590205 3.141592652590205 3.141592652590205 3.141592652590205 3.141592652590205]
```

- 8 Threads
```
Numero de termos 1000000000
Numero de threads: 8
Tempo Médio -> 28.871871340000002
Desvio padrão ->  0.6651578652872431
Coeficiente de variação ->  2.3038266465454647
Valores obtidos para Pi ->  [3.141592652589324 3.141592652589324 3.141592652589324 3.141592652589324 3.141592652589324]
```

- 4 Threads
```
Numero de termos 1000000000
Numero de threads: 4
Tempo Médio -> 37.53117068
Desvio padrão ->  0.3911083209215754
Coeficiente de variação ->  1.0420893189190956
Valores obtidos para Pi ->  [3.1415926525892104 3.1415926525892104 3.1415926525892104 3.1415926525892104 3.1415926525892104]
```

- 2 Threads
```
Numero de termos 1000000000
Numero de threads: 2
Tempo Médio -> 59.26495366
Desvio padrão ->  1.2106235817566073
Coeficiente de variação ->  2.0427310020385616
Valores obtidos para Pi ->  [3.141592652589258 3.141592652589258 3.141592652589258 3.141592652589258 3.141592652589258]
```

- 1 Threads
```
Numero de termos 1000000000
Numero de threads: 1
Tempo Médio -> 101.86898643999999
Desvio padrão ->  2.1955561200238525
Coeficiente de variação ->  2.155274334958675
Valores obtidos para Pi ->  [3.1415926525880504 3.1415926525880504 3.1415926525880504 3.1415926525880504 3.1415926525880504]

```


<br /> <br />
É interessante observar que estes resultados foram gerados por um computador com 8 nucles de processamento,
<br />
Desse modo, quando  testamos com mais de 8 threads é perceptivel que o tempo não se altera de forma singficativa

