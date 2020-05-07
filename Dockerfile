FROM node:lts-alpine

WORKDIR /app
COPY package.json /app/
COPY yarn.lock /app/

RUN apk add --update make python g++
RUN yarn install

COPY . /app/
CMD [ "npm", "start" ]
