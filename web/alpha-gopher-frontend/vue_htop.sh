#!/bin/sh

while true
do
    ./htop_to_html.sh ./public/htop.html
    ./html_to_vue.sh ./public/htop.html ./src/components/HTop.vue
    ./node_modules/eslint/bin/eslint.js ./src/components/HTop.vue --fix 
	sleep 8
done