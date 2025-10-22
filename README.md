# 🏠 iMob - Sistema de Gerenciamento de Contratos de Aluguel

[![Go](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-blue.svg)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-Compose-blue.svg)](https://docs.docker.com/compose/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Um sistema moderno e escalável para gerenciamento de contratos de aluguel de imóveis, desenvolvido em Go com arquitetura de microsserviços.

## 📋 Índice

- [Visão Geral](#-visão-geral)
- [Funcionalidades Atuais](#-funcionalidades-atuais)
- [Roadmap](#-roadmap)
- [Arquitetura](#-arquitetura)
- [Tecnologias](#-tecnologias)
- [Pré-requisitos](#-pré-requisitos)
- [Instalação](#-instalação)
- [Uso](#-uso)
- [API Endpoints](#-api-endpoints)
- [Testes](#-testes)
- [Contribuição](#-contribuição)

## 🎯 Visão Geral

O **iMob** é uma plataforma completa para gerenciamento de contratos de aluguel que permite:

- **Gestão de Imóveis**: Cadastro e controle de propriedades
- **Gestão de Usuários**: Proprietários, inquilinos e corretores
- **Contratos de Aluguel**: Criação, acompanhamento e renovação
- **Pagamentos**: Controle de mensalidades e taxas
- **Relatórios**: Dashboards e análises financeiras

### 🏗️ MVP Atual - Sistema Monolítico

Atualmente o sistema implementa um CRUD básico com:

- ✅ **Gestão de Usuários** (Proprietários/Inquilinos)
- ✅ **Gestão de Propriedades**
- ✅ **API RESTful** com Gin Framework
- ✅ **Banco PostgreSQL** com UUID v7
- ✅ **Docker Compose** para desenvolvimento
- ✅ **Arquitetura Clean** (Handlers, Services, Repositories)

## 🚀 Funcionalidades Atuais

### 👥 Gestão de Usuários

- Cadastro de usuários (proprietários e inquilinos)
- Busca por ID e email
- Atualização de dados
- Relacionamento com propriedades

### 🏠 Gestão de Propriedades

- Cadastro de imóveis
- Associação com proprietários
- Busca por ID e proprietário
- Informações básicas (título, endereço)

### 🔧 Infraestrutura

- API REST com validação de dados
- Banco PostgreSQL com UUID v7
- Logging estruturado
- Docker para containerização
- Configuração via variáveis de ambiente

## 🗺️ Roadmap

### 📋 Fase 1: CRUD Básico ✅ (Atual)

- [x] Gestão de Usuários
- [x] Gestão de Propriedades
- [x] API REST básica
- [x] Containerização

### 🔐 Fase 2: Autenticação & Autorização (Q1 2026)

- [ ] JWT Authentication
- [ ] Role-based Access Control (RBAC)
- [ ] OAuth2 integration
- [ ] API Rate Limiting
- [ ] Session Management

### 📄 Fase 3: Gestão de Contratos (Q1 2026)

- [ ] Modelo de Contratos de Aluguel
- [ ] Templates de contratos
- [ ] Assinatura digital
- [ ] Histórico de alterações
- [ ] Renovação automática

### 💰 Fase 4: Sistema de Pagamentos (Q2 2026)

- [ ] Integração com gateways de pagamento
- [ ] Controle de mensalidades
- [ ] Multas e juros
- [ ] Notificações de vencimento
- [ ] Relatórios financeiros

### 📊 Fase 5: Cache & Performance (Q2 2026)

- [ ] Couchbase para cache distribuído
- [ ] Redis para sessões
- [ ] Cache de consultas frequentes
- [ ] Otimização de queries

### 🔄 Fase 6: Mensageria & Eventos (Q3 2026)

- [ ] RabbitMQ para filas
- [ ] Event-driven architecture
- [ ] Processamento assíncrono
- [ ] Notificações em tempo real

### 🧩 Fase 7: Microsserviços (Q3-Q4 2026)

- [ ] **User Service**: Gestão de usuários
- [ ] **Property Service**: Gestão de imóveis
- [ ] **Contract Service**: Gestão de contratos
- [ ] **Payment Service**: Processamento de pagamentos
- [ ] **Notification Service**: Notificações
- [ ] **Report Service**: Relatórios e analytics

### ☸️ Fase 8: Orquestração Kubernetes (Q4 2026)

- [ ] Deployment em Kubernetes
- [ ] Service Mesh (Istio)
- [ ] Monitoring (Prometheus/Grafana)
- [ ] Logging centralizado (ELK Stack)
- [ ] CI/CD pipelines
- [ ] Auto-scaling

## 🏗️ Arquitetura

### Arquitetura Atual (Monolítica)

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   HTTP Client   │────│   API Gateway   │────│  Load Balancer  │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │
                       ┌─────────────────┐
                       │   iMob API      │
                       │   (Gin/Go)      │
                       └─────────────────┘
                                │
                       ┌─────────────────┐
                       │  PostgreSQL     │
                       │  (UUID v7)      │
                       └─────────────────┘
```

### Arquitetura Futura (Microsserviços)

```
                    ┌─────────────────┐
                    │   API Gateway   │
                    │   (Kong/Nginx)  │
                    └─────────────────┘
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
┌───────────────┐  ┌───────────────┐  ┌───────────────┐
│ User Service  │  │Property Service│  │Contract Service│
│    (Go)       │  │     (Go)       │  │     (Go)       │
└───────────────┘  └───────────────┘  └───────────────┘
        │                   │                   │
┌───────────────┐  ┌───────────────┐  ┌───────────────┐
│   PostgreSQL  │  │   PostgreSQL  │  │   PostgreSQL  │
│   + Couchbase │  │   + Couchbase │  │   + Couchbase │
└───────────────┘  └───────────────┘  └───────────────┘
                            │
                    ┌───────────────┐
                    │   RabbitMQ    │
                    │  (Messaging)  │
                    └───────────────┘
```

## 🛠️ Tecnologias

### Backend

- **Go 1.23**: Linguagem principal
- **Gin**: Framework web
- **GORM**: ORM para banco de dados
- **PostgreSQL 18**: Banco principal com UUID v7
- **UUID v7**: Identificadores ordenados por tempo

### Infraestrutura Atual

- **Docker & Docker Compose**: Containerização
- **PostgreSQL**: Banco de dados relacional

### Tecnologias Futuras

- **Couchbase**: Cache distribuído
- **RabbitMQ**: Message broker
- **Kubernetes**: Orquestração de containers
- **Istio**: Service mesh
- **Prometheus/Grafana**: Monitoring
- **ELK Stack**: Logging

## 📋 Pré-requisitos

- **Docker** >= 20.0
- **Docker Compose** >= 2.0
- **Go** >= 1.23 (para desenvolvimento local)
- **Git**

## 🚀 Instalação

### 1. Clone o repositório

```bash
git clone https://github.com/rafaelmfried/go-imob.git
cd go-imob
```

### 2. Configure as variáveis de ambiente

```bash
cp .env.example .env
# Edite o arquivo .env conforme necessário
```

### 3. Execute com Docker Compose

```bash
# Construir e executar
docker compose up --build -d

# Verificar logs
docker compose logs -f app

# Verificar status
docker compose ps
```

### 4. Acesse a aplicação

- **API**: http://localhost:8000
- **PostgreSQL**: localhost:5432

## 💻 Uso

### Executar localmente (desenvolvimento)

```bash
# Instalar dependências
go mod download

# Executar apenas o banco
docker compose up -d db

# Configurar variáveis de ambiente
export DBHOST=localhost
export DBPORT=5432
export DBUSER=postgres
export DBPASS=postgres
export DBNAME=imob

# Executar aplicação
go run cmd/api/main.go
```

### Executar testes

```bash
# Testes unitários
go test ./...

# Testes com coverage
go test -cover ./...

# Testes E2E (futuros)
make test-e2e
```

## 📚 API Endpoints

### 👥 Usuários

| Método   | Endpoint              | Descrição                |
| -------- | --------------------- | ------------------------ |
| `POST`   | `/users`              | Criar usuário            |
| `GET`    | `/users/:id`          | Buscar usuário por ID    |
| `GET`    | `/users?email=:email` | Buscar usuário por email |
| `PUT`    | `/users/:id`          | Atualizar usuário        |
| `DELETE` | `/users/:id`          | Deletar usuário          |

### 🏠 Propriedades

| Método   | Endpoint                   | Descrição                            |
| -------- | -------------------------- | ------------------------------------ |
| `POST`   | `/properties`              | Criar propriedade                    |
| `GET`    | `/properties/:id`          | Buscar propriedade por ID            |
| `GET`    | `/properties?owner_id=:id` | Buscar propriedades por proprietário |
| `PUT`    | `/properties/:id`          | Atualizar propriedade                |
| `DELETE` | `/properties/:id`          | Deletar propriedade                  |

### 📄 Exemplos de Uso

#### Criar Usuário

```bash
curl -X POST http://localhost:8000/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@example.com"
  }'
```

#### Criar Propriedade

```bash
curl -X POST http://localhost:8000/properties \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Apartamento 2 Quartos",
    "address": "Rua das Flores, 123",
    "owner_id": "019a0c8e-d27d-743b-991e-41fa5c06d9ce"
  }'
```

## 🧪 Testes

O projeto utiliza testes em HTTP REST Client:

### Arquivo .http

```http
# Teste completo do fluxo
@base_url=http://localhost:8000

### Criar usuário
# @name CreateUser
POST {{base_url}}/users
Content-Type: application/json

{
    "name": "Rafael Friederick",
    "email": "rafael@example.com"
}

### Criar propriedade
# @name CreateProperty
POST {{base_url}}/properties
Content-Type: application/json

{
    "title": "Casa Moderna",
    "address": "Rua Principal, 456",
    "owner_id": "{{CreateUser.response.body.$.id}}"
}
```

## 🏢 Modelo de Negócio - Contratos de Aluguel

### Entidades Principais (Futuras)

#### 📋 Contrato de Aluguel

- **Informações Básicas**: Número, data início/fim, valor
- **Partes**: Locador (proprietário) e Locatário
- **Imóvel**: Propriedade associada
- **Condições**: Reajuste, renovação, multas
- **Garantias**: Caução, fiador, seguro

#### 💰 Pagamentos

- **Mensalidades**: Valor, vencimento, status
- **Taxas**: Condomínio, IPTU, água, luz
- **Multas**: Atraso, quebra de contrato
- **Reajustes**: IGPM, IPCA, outros índices

#### 🔔 Notificações

- **Vencimentos**: Alertas de pagamento
- **Renovação**: Lembretes de fim de contrato
- **Manutenção**: Solicitações e agendamentos
- **Documentos**: Novos contratos, aditivos

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

### Padrões de Código

- Seguir [Effective Go](https://golang.org/doc/effective_go.html)
- Usar `gofmt` para formatação
- Escrever testes para novas funcionalidades
- Documentar funções públicas

## 📈 Status do Projeto

- 🟢 **Estável**: API básica funcional
- 🟡 **Em Desenvolvimento**: Novas funcionalidades
- 🔴 **Planejado**: Funcionalidades futuras

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👨‍💻 Autor

**Rafael Friederick** - [@rafaelmfried](https://github.com/rafaelmfried)

## 🙏 Agradecimentos

- Comunidade Go pela excelente documentação
- Contributors do projeto
- Inspiração em sistemas de gestão imobiliária existentes

---

⭐ **Star este projeto** se ele foi útil para você!

📧 **Entre em contato** se tiver dúvidas ou sugestões!
