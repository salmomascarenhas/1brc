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

## Como Executar o Código

### Pré-requisitos

- Go 1.19 ou superior instalado na sua máquina.
- Python 3.x instalado para geração do arquivo de entrada.

### Passos para Executar

1. Clone o repositório:

   ```bash
   git clone https://github.com/salmomascarenhas/1brc.git
   cd seu-repositorio
   ```

2. **Geração do Arquivo de Entrada**:

   Antes de executar o programa em Go, você precisa gerar o arquivo de entrada `measurements.txt` com os dados simulados. Para isso, utilize o script em Python disponível no repositório:

   ```bash
   python3 create_measurements.py 100000
   ```

   Onde `100000` representa a quantidade de linhas a serem geradas. O valor máximo recomendado é de aproximadamente 1 bilhão de linhas, o que resultará em um arquivo de cerca de 12GB.

   **Nota**: O arquivo `measurements.txt` deve ser sobrescrito com o conteúdo gerado pelo script.

3. **Execução do Programa em Go**:

   Após gerar o arquivo de entrada, você pode compilar e executar o programa em Go:

   ```bash
   go run main.go
   ```

   O programa irá processar o arquivo `measurements.txt` e exibir os resultados diretamente no terminal.

### Exemplo de Uso

Para gerar um arquivo de 100 mil linhas e processá-lo com o programa em Go, utilize os seguintes comandos:

```bash
python3 create_measurements.py 100000
go run main.go
```

Esses comandos irão criar o arquivo de entrada `measurements.txt`, calcular as temperaturas mínima, média e máxima para cada estação, e exibir o resultado no terminal.

## Contribuição

Contribuições são bem-vindas! Se você encontrar algum problema ou tiver sugestões de melhoria, sinta-se à vontade para abrir uma issue ou enviar um pull request.

### **Desafio de Performance 🚀**

Se você tem experiência com Go e conhece técnicas para reduzir o tempo de execução, como concorrência, otimização de algoritmos ou manipulação eficiente de arquivos, suas sugestões serão muito bem-vindas! Contribua com seu conhecimento para aprimorar ainda mais este projeto e torná-lo ainda mais eficiente.

## Ideia original
Créditos da ideia original em Java pode ser acessada no repositório do [gunnarmorling](https://github.com/gunnarmorling/1brc).

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
