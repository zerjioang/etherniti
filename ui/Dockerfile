FROM node:9.10.0-alpine as builder

MAINTAINER team@etherniti.org

WORKDIR /opt/app

COPY .babelrc /opt/app/.babelrc
COPY build/ /opt/app/build/
COPY config/ /opt/app/config/
COPY index.html /opt/app/index.html
COPY index.template /opt/app/index.template
COPY package.json /opt/app/package.json
COPY package-lock.json /opt/app/package-lock.json
COPY .postcssrc.js /opt/app/.postcssrc.js
COPY src/ /opt/app/src/
COPY static/ /opt/app/static/

RUN npm install -g yarn && npm install && /usr/local/bin/yarn build

# Pull dist files into a second stage deploy alpine container
FROM nginx:1.13.10-alpine

COPY --from=builder /opt/app/dist /usr/share/nginx/html
