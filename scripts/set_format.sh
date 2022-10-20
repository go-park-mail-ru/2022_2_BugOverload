#!/bin/zsh 

for file in `find $1 -type f -name "*"`
do
	mv $file "${file%.*}.$2"
done 

