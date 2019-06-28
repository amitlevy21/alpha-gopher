#!/bin/sh

while true
do
    ./htop_to_html.sh ./public/htop.html
    ./html_to_vue.sh ./public/htop.html ./src/components/HTop.vue
	sleep 8
done