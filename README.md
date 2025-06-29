# Todo API - Go Backend

Uma API RESTful para gerenciamento de tarefas (todos) desenvolvida em Go com MongoDB como banco de dados.

## 🚀 Funcionalidades

- ✅ Criar novos todos
- 📋 Listar todos os todos
- 🔍 Buscar todo por ID
- ✏️ Atualizar todos existentes
- 🗑️ Deletar todos
- 🏥 Health check da API

## 🛠️ Tecnologias Utilizadas

- **Go 1.24** - Linguagem de programação
- **MongoDB** - Banco de dados NoSQL
- **Chi Router** - Roteamento HTTP
- **Docker & Docker Compose** - Containerização

## 📋 Pré-requisitos

- [Go 1.24+](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/) (opcional)

## ⚙️ Configuração do Ambiente

### 1. Clone o repositório
```bash
git clone <url-do-repositorio>
cd go-backend
```

### 2. Configure as variáveis de ambiente
```bash
cp .env.sample .env
```

Edite o arquivo `.env` conforme necessário:
```bash
BINARY=mongo-todos

CONTAINER_NAME=mongo-todos-container
DB_CONTAINER_NAME=mongo-todos-db

MONGO_DB=todos-db
MONGO_DB_USERNAME=admin
MONGO_DB_PASSWORD=admin123
```

### 3. Instale as dependências do Go
```bash
go mod tidy
```

## 🐳 Executando com Docker (Recomendado)

### Usando Docker Compose
```bash
# Subir MongoDB e aplicação
docker-compose up -d

# Para parar os containers
docker-compose down
```

### Usando Makefile
```bash
# Subir os serviços
make up

# Parar os serviços
make down
```

## 🏃‍♂️ Executando Localmente

### 1. Inicie o MongoDB
Certifique-se de que o MongoDB está rodando em `localhost:27017` ou use Docker:
```bash
docker run -d \
  --name mongo-todos-db \
  -p 27017:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=admin \
  -e MONGO_INITDB_ROOT_PASSWORD=admin123 \
  -e MONGO_INITDB_DATABASE=todos-db \
  mongo:latest
```

### 2. Configure as variáveis de ambiente
```bash
export MONGO_DB_USERNAME=admin
export MONGO_DB_PASSWORD=admin123
```

### 3. Execute a aplicação
```bash
# Compilar e executar
go run cmd/api/main.go

# Ou usando o Makefile
make build
make start

# Ou restart (build + start)
make restart
```

A API estará disponível em: `http://localhost:8080`

## 📚 Endpoints da API

### Health Check
```http
GET /api/v1/health
```

### Todos
```http
# Listar todos os todos
GET /api/v1/todos

# Buscar todo por ID
GET /api/v1/todos/{id}

# Criar novo todo
POST /api/v1/todos/create
Content-Type: application/json

{
  "title": "Minha tarefa",
  "description": "Descrição da tarefa",
  "done": false
}

# Atualizar todo
PUT /api/v1/todos/update/{id}
Content-Type: application/json

{
  "title": "Tarefa atualizada",
  "description": "Nova descrição",
  "done": true
}

# Deletar todo
DELETE /api/v1/todos/delete/{id}
```

## 📁 Estrutura do Projeto

```
go-backend/
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada da aplicação
├── db/
│   └── db.go               # Configuração do MongoDB
├── handlers/
│   ├── handlers.go         # Health check handler
│   ├── router.go           # Configuração das rotas
│   └── todo_handlers.go    # Handlers dos todos
├── services/
│   ├── main.go             # Configuração dos serviços
│   └── todo.go             # Lógica de negócio dos todos
├── docker-compose.yml      # Configuração do Docker Compose
├── Dockerfile              # Imagem Docker da aplicação
├── Makefile               # Comandos de automação
├── go.mod                 # Dependências do Go
├── go.sum                 # Checksums das dependências
├── .env.sample            # Exemplo de variáveis de ambiente
└── README.md              # Este arquivo
```

## 🔧 Comandos Úteis

```bash
# Compilar o projeto
go build -o mongo-todos ./cmd/api/main.go

# Executar testes
go test ./...

# Formatar código
go fmt ./...

# Verificar dependências
go mod tidy

# Ver logs do container
docker logs mongo-todos-container

# Acessar o container da aplicação
docker exec -it mongo-todos-container sh

# Acessar o MongoDB
docker exec -it mongo-todos-db mongosh -u admin -p admin123
```

## 🐛 Troubleshooting

### Erro de conexão com MongoDB
- Verifique se o MongoDB está rodando na porta 27017
- Confirme as credenciais no arquivo `.env`
- Verifique se as variáveis de ambiente estão sendo carregadas

### Erro "package not found"
- Execute `go mod tidy` para sincronizar as dependências
- Verifique se o nome do módulo no `go.mod` está correto

### Porta já em uso
- Verifique se não há outros serviços rodando na porta 8080
- Para MongoDB, verifique a porta 27017

## 📝 Estrutura do Todo

```go
type Todo struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Done        bool      `json:"done"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
