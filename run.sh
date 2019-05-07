#!/bin/bash

# Compila e roda o programa

# Dependências
#	go
#	bash

# Parâmetros (só um deles)
#	--run - executa o programa
#	--install - gera um executável 'ircclient'

SRCDIR=$GOPATH/src/github.com/Redes-2019
OBJDIRS="userinterface connection tui"

for dir in $OBJDIRS ; do
	echo "Compilando $dir"
	go build $SRCDIR/$dir
done;

if [[ $1 == '--run' ]]; then
	printf "\nRodando\n"
	go run $SRCDIR/ircclient
fi; 
if [[ $1 == '--install' ]]; then
	printf "Gerando executável\n"
	go install $SRCDIR/ircclient
fi;
