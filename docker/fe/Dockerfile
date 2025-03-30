FROM node:21-alpine

WORKDIR /code
RUN mkdir /code/ls-app
# 先にpackage.jsonとpackage-lock.jsonをマウントさせる
COPY ls-app/package*.json /code/ls-app

WORKDIR /code/ls-app

# RUN npm install