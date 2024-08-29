

# 1️⃣🐝🏎️ Desafio 1BRC em Go

Este repositório contém uma implementação em Go para o desafio proposto, que envolve a leitura de um arquivo de texto com dados de temperaturas para várias estações meteorológicas. O objetivo é calcular a temperatura mínima, média e máxima para cada estação e exibir os resultados de maneira organizada.

## Descrição do Projeto

O desafio consiste em processar um arquivo de texto onde cada linha contém uma medição de temperatura em uma estação meteorológica específica. Cada linha segue o formato:

```plaintext
<string: nome da estação>;<double: medição>
```

A medição é um valor numérico com exatamente uma casa decimal. Um exemplo de 10 linhas do arquivo seria:

```plaintext
Hamburg;12.0
Bulawayo;8.9
Palembang;38.8
St. John's;15.2
Cracow;12.6
Bridgetown;26.9
Istanbul;6.2
Roseau;34.4
Conakry;31.2
Istanbul;23.0
```

### Objetivo

A tarefa é escrever um programa em Go que:

1. Leia o arquivo de entrada.
2. Calcule os valores de temperatura mínima, média e máxima para cada estação meteorológica.
3. Exiba os resultados em stdout no seguinte formato, ordenado alfabeticamente pelo nome da estação:

```plaintext
<min>/<mean>/<max>
```

Os resultados devem ser arredondados para uma casa decimal e apresentados entre chaves `{}`.

Exemplo de saída:

```plaintext
{Hamburg=12.0/12.0/12.0, Istanbul=6.2/14.6/23.0, Palembang=38.8/38.8/38.8}
```

## Como Executar o Código em Go

### Pré-requisitos

- Go 1.19 ou superior instalado na sua máquina.

### Passos para Executar

1. Clone o repositório:

   ```bash
   git clone https://github.com/salmomascarenhas/1brc.git
   cd seu-repositorio
   ```

2. Prepare seu arquivo de entrada (`input.txt`) no formato especificado e coloque-o no diretório raiz do projeto.

3. Compile e execute o programa:

   ```bash
   go run main.go -input=input.txt
   ```

   - `-input`: Especifica o caminho para o arquivo de entrada.

4. O programa irá processar o arquivo e exibir o resultado diretamente no terminal.

### Exemplo de Uso

Para executar o programa com um arquivo de entrada chamado `input.txt`, use o comando:

```bash
go run main.go -input=input.txt
```

Este comando processará os dados, calculará as temperaturas mínima, média e máxima para cada estação e exibirá o resultado no terminal.

## Contribuição

Contribuições são bem-vindas! Se você encontrar algum problema ou tiver sugestões de melhoria, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
