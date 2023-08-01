# tabnews-rss

> Receba as atualizações dos seus usuários favoritos

## Build & Deploy

Você pode rodar esse serviço usando Docker ou mesmo seu sistema de uso normal, primeiramente clone o repositório:

```bash
git clone https://github.com/dheison0/tabnews-rss
```

As seguintes variáveis de ambiente são aceitas:
    - `PORT` Posta onde o servidor vai ficar rodando(padrão: 8080);
    - `DATABASE_FILE` Onde o arquivo de banco de dados SQLite3 vai ficar.

Agora rode usando:

  - [Docker](#Docker)
  - [Sistema](#Sistema)

### Docker

Construa a imagem usando o seguinte comando:

```bash
docker build -t tabrss .
```

Caso queira salvar o banco de dados para poder atualizar o serviço futuramente, crie um volume:

```bash
docker volume create tabrss-db
```

Agora inicie o servidor:

```bash
docker run -d \
    --name tabrss                    \
    -v tabrss-db:/data               \
    -e DATABASE_FILE=/data/tabrss.db \
    -p 8181:8080                     \
    --restart unless-stopped         \
    tabrss:latest
```

Pronto, só acessar o painel na rota http://localhost:8181

### Sistema

Instale o compilador Go(ex: Debian/Ubuntu):

```bash
sudo apt install golang
```

E compile o projeto:

```bash
go build -o tabrss .
```

Agora é só rodar com `./tabrss`, caso queira definir as variáveis, aqui vai um exemplo:

```bash
export PORT=1900
export DATABASE_FILE=./meuBancoDeDados.db
./tabrss
```
#### Atenção!
**O servidor deve ser iniciado de dentro da pasta do projeto para que tenha acesso ao conteúdo da pasta web**