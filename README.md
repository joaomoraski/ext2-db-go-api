# EXT2-DB: Um "mini-db" usando EXT2 com API em Go

Este projeto é a união de um trabalho realizado durante o curso de Ciencia da Computação, que tinha como objetivo a
criação de um Shell Interativo para rodar comando em cima de uma imagem(.iso ou .img) usando o Sistema de Arquivos EXT2
com uma API em Go, utilizando estas operaçǒes desenvolvidas no trabalho como um "banco de dados".

[Link para o projeto do EXT2](https://github.com/joaomoraski/ext2-fs-tools/tree/master) (trocar para branch `db-engine`
para este projeto)

## Arquitetura

O sistema é composto por dois componentes principais que se comunicam via I/O padrão (`stdin`/`stdout`):

1. **Banco de dados (`ext2-db-engine`):** Uma ferramenta de linha de comando em C, compilada a partir do
   projeto original, que manipula diretamente uma imagem de disco no formato EXT2. Ela é responsável por todas as
   operações de baixo nível, como alocação de inodes, manipulação de bitmaps e escrita/leitura de blocos de dados.
2. **API em Go:** Uma API RESTful simples, que expõe endpoints para operações de **Criar** e **Listar(com filtros)** em
   um "banco de dados" de usuários. Ela trata as requisições HTTP, valida os dados e chama o "banco de dados" em C para
   persistir ou recuperar as informações.

## Como Rodar o Projeto

### Pré-requisitos

* Compilador GCC
* Biblioteca `readline` (`sudo apt-get install libreadline-dev` ou similar)
* Go (versão 1.18 ou superior)

### Compilação e Execução

O projeto utiliza um `Makefile` na raiz do sub-módulo `ext2-fs-tools` para compilar o motor em C e um `Makefile` na raiz
do projeto `ext2-db-go-api` para orquestrar tudo.

1. **Compile o motor em C:**
   ```bash
   cd ext2-db-engine
   make all
   cd ..
   ```

   1.1 **Para gerar uma imagem nova**
   ```bash
      make generate-ext2
   ```

   1.2 **Criando a tabela**
   ```bash
      make run
   ```
   Dentro do terminal que sera iniciado, crie a database com o comando `touch user_record`

2. **Inicie a API em Go:**
   Na raiz do projeto, execute:
   ```bash
   cd api
   go run .
   ```
   O servidor estará rodando em `http://localhost:8080`.

---

## Como Usar a API

A API expõe os seguintes endpoints para a entidade `User`:

### Criar um Novo Usuário

* **Endpoint:** `POST /users`
* **Descrição:** Cria um novo registro de usuário no arquivo `/user_record` dentro da imagem EXT2.
* **Corpo da Requisição (JSON):**
  ```json
  {
      "id": 1,
      "is_active": 1,
      "username": "moraski",
      "email": "moraski@gmail.com"
  }

### Buscar Usuários

* **Endpoint:** `GET /users`
* **Descrição:** Retorna uma lista de todos os usuários cadastrados.
* **Query Parameters:**
    * limit=<numero>: Limita o número de resultados.
    * filters=<condições>: (Opcional) Filtra os resultados. As condições devem estar no padrão `field:operator:value`.
        * Suporta =, %
