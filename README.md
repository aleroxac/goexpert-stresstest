# goexpert-stresstest
Projeto do Desafio Técnico "Sistema de Stress test" do treinamento GoExpert(FullCycle).



## O desafio
Criar um sistema CLI em Go para realizar testes de carga em um serviço web. 
O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.
O sistema deverá gerar um relatório com informações específicas após a execução dos testes.



## Como rodar o projeto: manual
``` shell
docker build -t aleroxac/goexpert-stresstest:v1 .
docker run aleroxac/goexpert-stresstest:v1 -—url=http://google.com -—requests=10 —-concurrency=2
```

## Como rodar o projeto: make
``` shell
make build
make runc
```



## Funcionalidades da Linguagem Utilizadas
- cli: cobra
- net/http
- contexts
- waitgroups
- channels
- mutex



## Requisitos: parâmetros
Entrada de Parâmetros via CLI:
- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requests.
- `--concurrency`: Número de chamadas simultâneas.

## Requisitos: execução do teste
- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.

## Requisitos: relatório
Apresentar um relatório ao final dos testes contendo:
- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Requisitos: execução do app
Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:
``` shell
docker run <sua imagem docker> --url=http://google.com -—requests=1000 —-concurrency=10
```
