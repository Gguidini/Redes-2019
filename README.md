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

É necessário ter a linguagem Go instalada e configurada do modo padrão.  
Também é necessário rodar o script run.sh em um computador com bash (Linux em geral).

## Build

Para compilar o projeto é necessário compilar cada pacote separadamente, e por fim rodar ou instalar o módulo principal, que é o **ircclient**. Assumindo que se esteja na pasta definida em $GOPATH, é possível rodar o script de build automatizado.

Para rodar o programa sem gerar um executável:

```
$ ./main.sh --run
```

Para gerar um executável:

```
$ ./main.sh --install
```


