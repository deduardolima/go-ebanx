
# Go Ebanx

Este projeto é uma API desenvolvida em Go, seguindo a arquitetura Clean Architecture, para gerenciar operações bancárias básicas como depósitos, saques, transferências e consulta de saldo. Ele disponibiliza uma interface HTTP para interagir com as funcionalidades bancárias.


## Funcionalidades

- Depósito em contas bancárias.
- Saque de contas bancárias.
- Transferências entre contas.
- Consulta de saldo de contas.
- Arquitetura escalável e de fácil manutenção.
- Testes automatizados para garantir a funcionalidade.


## Instalação

#### Clonando o Repositório
Primeiro, clone o repositório do projeto:

```bash
git clone https://github.com/seu-usuario/go-ebanx.git
cd go-ebanx

```
#### Construindo e Subindo o Container
Este projeto usa Docker para simplificar a execução. Para construir a imagem e subir o container, use:

```bash
docker-compose up --build

```
Este comando irá:

- Construir a imagem Docker para o projeto.
- Subir o container na porta 8080.



## Deploy

### Usando Ngrok

Para expor a API localmente utilizando Ngrok, siga os passos abaixo:

1. **Instale o Ngrok**:

   Baixe e instale o Ngrok diretamente do site oficial:

   [Baixar Ngrok](https://ngrok.com/download)

   Siga as instruções de instalação de acordo com seu sistema operacional.

2. **Inicie o Ngrok para expor a porta 8080**:

   Após instalar o Ngrok, inicie-o com o seguinte comando:

   ```bash
   ngrok http 8080
   ```

3.  Copie o link gerado pelo Ngrok e utilize-o para acessar a API de qualquer lugar.



## Rodando os testes

Este projeto possui testes automatizados. Para rodar os testes:

1. Acesse a API para testar as funcionalidades:

```bash
https://ipkiss.pragmazero.com/
```

Após executar os testes, você terá seguinte resultado


## Autores

- [@deduardolima](https://github.com/deduardolima)