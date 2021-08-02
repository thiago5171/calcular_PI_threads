
Esta atividade avaliativa tem como principal objetivo compreender melhor o conceito de thread e como elas podem ser usadas para utilizar melhor os recursos de processamento disponíveis nos computadores e sistemas operacionais modernos. Ela é inspirada nesta atividade(http://wiki.inf.ufpr.br/maziero/doku.php?id=so:calculo_de_pi_com_threads), do curso de sistemas operacionais do prof. Carlos Maziero (UFPR).

Construa um programa utilizando a sua linguagem de programação de preferência para calcular o valor de π utilizando N threads processando de forma concorrente. O valor de π deve ser aproximado pela série de Leibniz-Grégory(https://pt.wikipedia.org/wiki/F%C3%B3rmula_de_Leibniz_para_%CF%80):


Observações:

Devem ser calculados pelo menos 1 bilhão (109) de termos da série; use variáveis reais de dupla precisão (double) nos cálculos;
O programa deve dividir o espaço de cálculo uniformemente entre as N threads; cada thread efetua uma soma parcial de forma autônoma;
Lembrando que os resultados parciais de cada thread devem ser somados, pois o objetivo é chegar ao valor mais próximo do número π
Devem ser medidos os tempos de execução do programa para execuções com 1, 2, 4, 8, 16 e 32 threads (cenários experimentais). Para determinar o tempo de cada execução, você pode usar comando time do UNIX ou realizar a medição dentro do próprio programa;
Para que os resultados tenham valor estatístico, devem ser feitas pelo menos 5 execuções de cada cenário (repetições) e calculados o tempo médio de execução e o coeficiente de variação entre elas (coeficiente de variação = desvio-padrão / média); são aceitáveis coeficientes de variação de até 5%; caso contrário, as medições devem ser refeitas.
