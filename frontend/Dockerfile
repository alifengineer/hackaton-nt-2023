# we will use node:alpine docker image
FROM node:14.16.0-alpine3.13 as builder

WORKDIR /web

COPY ./package.json ./

# install npm dependencies
RUN npm install

# copy your project files to work directory
COPY . .

RUN npm run build