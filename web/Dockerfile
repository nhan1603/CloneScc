# web image for local env
FROM node:16.1.0
WORKDIR /web
COPY package.json yarn.lock ./
COPY ./ ./
RUN yarn cache clean && yarn install
RUN npm rebuild node-sass
