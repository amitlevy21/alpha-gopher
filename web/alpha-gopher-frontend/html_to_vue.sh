#!/bin/sh

src=$1
dst=$2

hxselect -c body < $src > $dst
sed -i '1s/^/<template>\n/' $dst
echo '</template>' >> $dst
