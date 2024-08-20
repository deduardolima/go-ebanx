<p align="center">
  <a href="#sobre">Go Ebanx</a> |
  <a href="#tecnologia">Tecnologias</a> |
  <a href="#features">Features</a> |
  <a href="#back">Rodando o back-end</a> |
  <a href="#documentacao">Documentação</a> |
  <a href="#ngrok">Usando Ngrok</a>
  <a href="#testes">Testes</a> |
  <a href="#desenvolvedores">Desenvolvedores</a> |
  <a href="#voltar">Voltar para o topo ⬆️</a>
</p>

<h2 id="sobre"> Go Ebanx</h2> |


Este projeto é uma API desenvolvida em Go, seguindo a arquitetura Clean Architecture, para gerenciar operações bancárias básicas como depósitos, saques, transferências e consulta de saldo. Ele disponibiliza uma interface HTTP para interagir com as funcionalidades bancárias.


<h2 id="tecnologia">🛠 Tecnologias</h2> |

- [**Go**](https://golang.org/doc/): Linguagem de programação desenvolvida pela Google.
- [**Docker**](https://www.docker.com/): Ferramenta de conteinerização, distribuição e execução de aplicações em contêineres.
- [**Google Cloud Run**](https://cloud.google.com/run): Serviço de computação sem servidor que executa contêineres Docker de forma escalável.
- [**GitHub Actions**](https://github.com/features/actions): Automação de CI/CD diretamente no GitHub para testar e implantar o código.
- [**Ngrok**](https://ngrok.com/): Ferramenta para criar túneis HTTP para expor localmente seu servidor ou aplicação a partir de uma URL pública.


<h2 id="features">🚀 Features</h2> |

A API possui endpoints para para realizar as transações de depósitos, saques e transferências entre contas

- [x] Depósito em contas bancárias
- [x] Saque de contas bancárias
- [x] Transferências entre contas.
- [x] Consulta de saldo de contas.
- [x] Testes automatizados para garantir a funcionalidade.
- [x] Pipeline de CI/CD implementada com GitHub Actions.


<h2 id="back"> 🚀 Rodando o back-end</h2> |

#### Clonando o Repositório
Primeiro, clone o repositório do projeto:

```bash
git clone https://github.com/deduardolima/go-ebanx.git
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

<h2 id="documentacao">📖 Documentação da API</h2> |

A documentação da API está disponível via Swagger. Após iniciar o servidor, acesse:

```bash
http://localhost:8080/swagger/index.html

```
Ou:
- [**Go Ebanx Documentation (Swagger)**](https://ebanx-service-qaetfstifq-uc.a.run.app/api/swagger/index.html)


<h2 id="ngrok">🌐 Usando Ngrok</h2> |

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

<h2 id="access-api">🌍 Acessando a API</h2> |

A API também está disponível publicamente através do Google Cloud Run. Você pode acessá-la no seguinte link:

- [**Go Ebanx API (Google Cloud Run)**](https://ebanx-service-qaetfstifq-uc.a.run.app/)
- [**Go Ebanx Documentation (Swagger)**](https://ebanx-service-qaetfstifq-uc.a.run.app/swagger/index.html)

Este link é uma URL temporária.


<h2 id="testes">🧪 Rodando Testes</h2> |

Este projeto possui testes automatizados. Para rodar os testes:

```
go test ./...
```

Ou acesse a API para testar as funcionalidades:

```bash
https://ipkiss.pragmazero.com/
```

Após executar os testes, você terá seguinte resultado :
| Testes |
|-------|

| <img src="https://github.com/user-attachments/assets/8b430429-f011-46da-99f6-b398c7d99191" alt="Login" width="300"> |

<h2 id="desenvolvedores">👨‍💻 Desenvolvedores</h2> |
<table>         
  <tr>
    <td align="center">
      <a href="https://github.com/deduardolima">
        <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/98969787?v=4" width="100px;" alt="Imagem profile Diego Lima desenvolvedor"/>
        <br />
        <sub><b> Diego Lima</b></sub>
    </td>
  </tr>
</table>

<h2 id="voltar">Voltar para o topo</h2>
