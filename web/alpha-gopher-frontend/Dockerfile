FROM node:lts-alpine

RUN apk add --no-cache bash htop aha html-xml-utils

# install simple http server for serving static content
RUN npm install -g http-server

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

# build app for production with minification
RUN npm run build

CMD [ "./start.sh" ]