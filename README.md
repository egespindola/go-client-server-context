# 🚀 Webserver HTTP, Contexts, database in memory (SQLite) e manupulação de arquivo

## Olá 👋

Seja bem-vindo(a)

### 🎯 Sobre o projeto

Neste projeto, existem dois sistemas: **client.go** e **server.go**.

O ***server*** é responsável por consultar o último câmbio de Dolár para Real consumindo a API https://economia.awesomeapi.com.br/json/last/USD-BRL e retornar o valor (bid) buscado. A cada requisição, os valores retornados pela API são salvos no banco em memória do SQLite e retornado ao client.
É utilizado o package **Context** para limitar o timeout de 200ms na requisição da API externa, e um timeout máximo de 10ms para registrar no banco de dados.

O ***client*** é responsável por consumir a API (server), **salvar o resultado no arquivo cotacao.txt** e **realizar logs dos possíveis erros retornados**. Também é utilizado um timeout de 300ms para a requisição.
A cotação atual é salva no arquivo no formato: Dólar: {valor}

Obs.: o endpoint do ws é /cotacao na porta 8080.

### 📍 Para testar:
1. Faça o clone do projeto com ***git clone***.
2. Instale os binários do Go (se não os tiver) em [go.dev](https://go.dev/dl/).
3. Instale as dependências com ***go mod tidy***.
4. Inicialize o web server com o comando ***go run cmd/ws/server.go***.
5. Rode o client através do comando ***go run cmd/client/client.go***.
6. Verifique o arquivo na raíz ***cotacao.txt***.

---

_"Programar é esculpir lógica com arte"_ ✨
