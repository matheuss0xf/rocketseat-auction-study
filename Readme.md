# Projeto de Leilão - Rocketseat

Este é um projeto de leilão desenvolvido durante o curso da Rocketseat, utilizando Go e o framework `chi` para a API.

## Tecnologias Utilizadas

- **Go**
- **chi (roteamento e middleware)**
- **SQLC (geração de código a partir de SQL)**
- **Tern (migração de banco de dados)**
- **Air (hot reload para desenvolvimento)**
- **WebSockets**
- **Autenticação e Sessões**
- **Docker - Ambiente de desenvolvimento**
- **PostgreSQL - Banco de dados**

## Instalação

1. Clone este repositório:
   ```sh
   git clone https://github.com/seu-usuario/rocketseat-auction-study.git
   ```
2. Acesse o diretório do projeto:
   ```sh
   cd rocketseat-auction-study
   ```
3. Configurar as Variáveis de Ambiente
Crie um arquivo .env na raiz do projeto e configure as seguintes variáveis:
    ```
    GOBID_DB_USER=<usuario_do_banco>
    GOBID_DB_NAME=<nome_do_banco>
    GOBID_DB_PASSWORD=<senha_do_banco>
    GOBID_DB_HOST=localhost
    GOBID_DB_PORT=5432
    ```
4. Subir o Banco de Dados com Docker
    Se ainda não tiver o Docker instalado, faça o download e instale-o primeiro.

    Para rodar o PostgreSQL via Docker, use:
    ```sh
   docker-compose up -d
    ```
5. Instale as dependências:
   ```sh
   go mod tidy
   ```
4. Execute o projeto:
   ```sh
   go run main.go
   ```

## Rotas da API

### Autenticação de Usuário

- `POST /api/v1/users/signup` - Cadastro de um novo usuário.
- `POST /api/v1/users/login` - Login de um usuário.
- `POST /api/v1/users/logout` - Logout do usuário (requer autenticação).

### Produtos e Leilões

- `POST /api/v1/products/` - Cria um novo produto para leilão (requer autenticação).
- `GET /api/v1/products/ws/subscribe/{product_id}` - Inscreve um usuário no leilão via WebSocket (requer autenticação).

## Middleware

- `RequestID` - Adiciona um ID único para cada requisição.
- `Recoverer` - Recupera pânicos e evita que o servidor caia.
- `Logger` - Loga as requisições no console.
- `AuthMiddleware` - Middleware para verificar a autenticação dos usuários.
- `Sessions` - Gerencia sessões de usuários.

---
📜 Licença
Este projeto foi criado para fins de estudo e aprendizado.
---