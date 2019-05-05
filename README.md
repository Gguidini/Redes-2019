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

## Build

Para compilar o projeto é necessário compilar cada pacote separadamente, e por fim rodar ou instalar o módulo principal, que é o **ircclient**. Assumindo que se esteja na pasta definida em $GOPATH, a compilacão de cada módulo é feita da seguinte forma:


```
$ go build github.com/Redes-2019/userinterface
```

Após compilados todos os módulos secundários pode-se rodar o programa principal:

```
$ go run github.com/Redes-2019/ircclient
```

(No futuro essa compilacão será feita através de um Makefile)

