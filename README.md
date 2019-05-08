# Redes-2019
Repositório para trabalhos de Redes, UnB (2019/1)

Diego Vaz Fernandes
Giovanni M Guidini
Matheus de Sousa Lemos Fernandes

## Estrutura

O projeto está estruturado seguindo a arquitetura padrão Go.  
Arquivos .go estão presentes na pasta *src*, e executáveis são armazenados na pasta *bin*.  

## Configuracão de pastas
É necessário criar uma estrutura de pastas no seguinte esquema:  

```
$ $GOPATH/src/github.com/ 
```

A pasta Redes-2019 deve ser colocada na pasta *$GOPATH/src/github.com/*

## Dependências

É necessário ter a linguagem Go instalada e configurada do modo padrão.  
Também é necessário rodar o script run.sh em um computador com bash (Linux em geral).

## Execucão

### Local

Para compilar o projeto é necessário compilar cada pacote separadamente, e por fim rodar ou instalar o módulo principal, que é o **ircclient**. Assumindo que se esteja na pasta definida em $GOPATH, é possível rodar o script de build automatizado.

Para rodar o programa sem gerar um executável:

```
$ ./main.sh --run
```

Para gerar um executável:

```
$ ./main.sh --install
```

Para fins de teste é possível conectar-se ao Freenode, um servidor aberto de IRC.  
No client as credenciais seriam

* Remote server name: chat.freenode.net
* Porta: 6667
* Senha: não 

### Docker

O client tem a opcão de ser executado num container docker. Para executá-lo, basta executar o Dockerfile

```
$ docker build . -t ircclient
```

```
$ docker run -i ircclient
```

As credenciais são as mesmas da conexão local.



