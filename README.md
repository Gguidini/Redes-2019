# Redes-2019
Repositório para trabalhos de Redes, UnB (2019/1)

## Estrutura

O projeto está estruturado seguindo a arquitetura padrão Go.  
Arquivos .go estão presentes na pasta *src*, e executáveis são armazenados na pasta *bin*.  

## Configuracão de pastas
É necessário criar uma estrutura de pastas no seguinte esquema:  

```
$ $GOPATH/src/github.com/ 
```

A pasta Redes-2019 deve ser colocada na pasta *$GOPATH/src/github.com/*

# Dependências

O projeto utiliza o projeto marcusolsson/tui-go para implementacão da interface de terminal (TUI). É necessário baixar como dependência

```
$ go get github.com/marcusolsson/tui-go
```


## Build

Para compilar o projeto é necessário compilar cada pacote separadamente, e por fim rodar ou instalar o módulo principal, que é o **ircclient**. Assumindo que se esteja na pasta definida em $GOPATH, é possível rodar o script de build automatizado.

```
$ ./run.sh
```


