# -- stage 1: The build stage

FROM node:18-alpine AS build-stage

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

# -- stage 2 The production stage ---

FROM nginx:stable-alpine

RUN rm /etc/nginx/conf.d/default.conf

COPY nginx.conf /etc/nginx/conf.default

COPY --from=build-stage /app/dist /usr/share/nginx/html

EXPOSE 80

