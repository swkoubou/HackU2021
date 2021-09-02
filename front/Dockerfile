
# ビルド環境
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN yarn install
COPY . .
RUN sed -i -e "s#^import firebase from 'firebase/app'#import firebase from 'firebase/compat/app'#" ./node_modules/firebaseui/dist/esm.js
RUN yarn build

# 本番環境
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
