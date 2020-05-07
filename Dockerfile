FROM node:lts-alpine

WORKDIR /app
COPY package.json /app/
COPY yarn.lock /app/

RUN apk add --update make python
RUN yarn install

COPY . /app/
CMD [ "npm", "start" ]
