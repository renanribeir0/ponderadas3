Este repositório contém os códigos baseados na implementação de um projeto utilizando o framework Go, com foco nas práticas de Desenvolvimento Dirigido por Testes (TDD). O objetivo é demonstrar as técnicas de TDD em projetos Go, incluindo a integração com rotas e manipulação de dados.

## Estrutura do Projeto

O projeto está organizado da seguinte forma:

```
/ponderadaRomualdo
|-- /handlers
|   |-- handler.go
|   |-- handler_test.go
|
|-- /db
|   |-- database.go
|   |-- database_test.go
|
|-- main.go
|-- go.mod
|-- go.sum
```

- **/handlers:** Contém os manipuladores de requisições HTTP para a aplicação.
- **/db:** Gerencia as conexões e consultas ao banco de dados.
- **main.go:** Ponto de entrada principal da aplicação.
- **go.mod e go.sum:** Gerenciam as dependências do projeto Go.

## Configuração do Ambiente de Testes

Para executar os testes, certifique-se de ter o Go configurado corretamente e de que as dependências estejam instaladas. Use o seguinte comando para garantir que todas as dependências estão instaladas:

```bash
go mod tidy
```

## Execução dos Testes

Os testes podem ser executados utilizando o comando:

```bash
go test ./...
```

### Testes Incluídos

#### Teste de Inserção de Dados

- **Descrição:** Verifica se a inserção de dados no banco ocorre corretamente.
- **Detalhes:**
  - Um objeto fictício é criado e inserido no banco.
  - A operação de inserção é validada utilizando a biblioteca Testify para garantir que não houve erros.

#### Teste de Recuperação de Dados

- **Descrição:** Verifica se os dados podem ser recuperados corretamente após serem inseridos.
- **Detalhes:**
  - O objeto inserido anteriormente é recuperado e seus dados são comparados para garantir a consistência.

#### Teste de Atualização de Dados

- **Descrição:** Verifica se a atualização de dados no banco ocorre corretamente.
- **Detalhes:**
  - Um objeto existente é atualizado no banco.
  - O teste verifica se a operação ocorreu sem erros e se os dados foram atualizados corretamente.

#### Teste de Deleção de Dados

- **Descrição:** Verifica se a deleção de um objeto no banco ocorre corretamente.
- **Detalhes:**
  - Um objeto é deletado do banco, e o teste verifica se ele não pode mais ser encontrado.

## Considerações Finais

Este projeto foi desenvolvido para demonstrar como aplicar TDD em um projeto Go. Todos os exemplos seguem o ciclo Red-Green-Refactor. Para mais detalhes sobre as implementações, consulte os arquivos de código comentados.

