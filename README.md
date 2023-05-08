# Erros

## Sobre

Este conjunto de programas serve à demonstração da manipulação de erros na linguagem de programação [Go](https://go.dev).

## Estrutura do repositório

Os arquivos fonte deste repositório, cada um demonstrando um exemplo, estão organizados de acordo com a seguinte estrutura:

```
+─errors            ---> Diretório com exemplo envolvendo a criação de variáveis de erros
  ├─── concat       ---> Diretório com exemplos utilizando generics
  ├─── division     ---> Diretório com exemplo envolvendo a criação de tipos customizados de erro
  ├─── division-rem ---> Diretório com exemplo envolvendo a criação de variáveis de erros
  ├─── request      ---> Diretório com exemplo envolvendo manipulação de múltiplos erros
```

## Requisitos

Para a compilação e execução dos programas, os seguintes elementos devem estar devidamente instalados no ambiente de desenvolvimento:

- [Git](https://git-scm.com), como sistema de controle de versões
- [Go](https://go.dev), incluindo compilador, ambiente de execução e outras ferramentas associadas à linguagem de programação Go

## Download, compilação e execução dos programas

No terminal do sistema operacional, insira os seguintes comandos para realizar o *download* da implementação a partir deste repositório Git:

```bash
# Download da implementação a partir do repositório Git
git clone https://github.com/ufrn-golang/errors.git
```

Para executar um programa, deve-se primeiro navegar para o respectivo diretório no qual ele se encontra e utilizar o comando `go run` no terminal do sistema operacional. Por exemplo, para executar o programa referente ao arquivo [`errors/erequest.go`](https://github.com/ufrn-golang/errors/blob/master/request/erequest.go), deve-se utilizar os seguintes comandos:

```bash
# Navegar para o diretório do programa
cd errors/request

# Compilar e executar o programa
go run erequest.go
```

Caso deseje gerar um executável para o programa em questão, deve-se utilizar o comando `go build` e, em seguida, invocar o arquivo executável gerado. Por exemplo, Por exemplo, para compilar e executar o programa referente ao arquivo [`errors/erequest.go`](https://github.com/ufrn-golang/errors/blob/master/request/erequest.go), deve-se utilizar os seguintes comandos:

```bash
# Navegar para o diretório do programa
cd generics/request

# Compilação do programa e geração de arquivo executável (com o mesmo nome do arquivo fonte)
go build erequest.go

# Execução do programa
./erequest
```
