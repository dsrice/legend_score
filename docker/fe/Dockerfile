FROM node:21-alpine

WORKDIR /code
RUN mkdir /code/myapp
# 先にpackage.jsonとpackage-lock.jsonをマウントさせる
COPY ls-app/package*.json /code/myapp

WORKDIR /code/myapp

# RUN npm install