# API Cliente-Servidor em Go

Este projeto é uma API cliente-servidor desenvolvida em Go como parte do desafio Full Cycle. Ele demonstra a implementação de uma arquitetura simples de cliente-servidor usando as bibliotecas padrão do Go.

## Estrutura do Projeto

O repositório está organizado nas seguintes pastas:

- `client/`: Contém o código da aplicação cliente.
- `server/`: Contém o código da aplicação servidor.

## Pré-requisitos

- Go 1.16 ou superior instalado na sua máquina. Você pode baixá-lo no [site oficial do Go](https://golang.org/dl/).

## Funcionalidades

- O servidor obtém a cotação do dólar em reais por meio de uma API externa.
- O resultado de cada chamada é armazenado pelo servidor em um banco de dados SQLite.
- O cliente solicita a cotação ao servidor e armazena o resultado em um arquivo chamado `cotacao.txt`.

## Começando

Siga as instruções abaixo para configurar e executar o projeto na sua máquina local.

### Clonar o Repositório

```bash
git clone https://github.com/brunocordeiro180/go-fc-client-server-api.git
cd go-fc-client-server-api
```

### Executando o Servidor

1. Navegue até o diretório `server`:

   ```bash
   cd server
   ```

2. Baixe as dependências:

   ```bash
   go tidy
   ```

3. Compile e execute o servidor:

    ```bash
   go run server.go
    ```

    O servidor será iniciado e ficará aguardando conexões de clientes.

### Executando o Cliente

1. Abra um novo terminal.
2. Navegue até o diretório `client`

   ```bash
   cd client
   ```

1. Compile e execute o cliente:

    ```bash
   go run client.go
    ```

    O servidor será iniciado e ficará aguardando conexões de clientes.

