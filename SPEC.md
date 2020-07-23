## Desenvolvedor(a) Go Sênior ##

1) Criamos uma base com os dados([data.json](data.json)) na seguinte estrutura:

```json
{
    "hostname": "server7",
    "cpu_load": {
        "value": 0.5627961727877718,
        "unit": "%"
    },
    "memory_size": {
        "value": 16,
        "unit": "GB"
    },
    "memory_usage": {
        "value": 5.881018176090025,
        "unit": "GB"
    },
    "disk_size": {
        "value": 100,
        "unit": "GB"
    },
    "disk_usage": {
        "value": 33.89572263431528,
        "unit": "GB"
    }
}
```


2) A partir dessa base, crie uma API que receba a requisição abaixo:
```json
{
    "hostname": "xxxxx"
}
```

3) E apresente a média, moda e tendência (*i.e.* mediana) de utilização para os itens abaixo por servidor:
    - CPU
    - Memória
    - Disco

4) Você define formato da resposta

5) Crie uma documentação de como rodar o seu projeto.

6) Leve em conta os seguintes aspectos durante o desenvolvimento:
   - Manutenibilidade
   - Arquitetura
   - Qualidade
   - Desempenho

