#!/bin/bash

# Compila e roda o programa

# Dependências
#	go
#	bash

SRCDIR=$GOPATH/src/github.com/Redes-2019
OBJDIRS="userinterface connection"

for dir in $OBJDIRS ; do
	echo "Compilando $dir"
	go build $SRCDIR/$dir
done;

printf "\nRodando\n"

go run $SRCDIR/ircclient