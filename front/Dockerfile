
# ビルド環境
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN yarn install
COPY . .
RUN  sed -i  -e "s#^(function() { var firebase=require('firebase/app');require('firebase/auth')#(function() { var firebase=require('firebase/compat/app');require('firebase/compat/auth')#" ./node_modules/firebaseui-ja/dist/npm__ja.js
RUN yarn build

# 本番環境
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
