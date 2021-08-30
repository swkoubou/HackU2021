# 参考: https://nodejs.org/de/docs/guides/nodejs-docker-webapp/

FROM node:16.8


# アプリのディレクトリーを作成
WORKDIR /usr/src/app


# package関連のjsonをコピー
COPY test-app/package*.json ./


# yarnと、必要なパッケージをインストール
#RUN npm install --global yarn
# => npm ERR! EEXIST: file already exists
# おそらく最初から入っている
RUN yarn install


# Bundle app source(? 多分ソースコードをコピー)
COPY test-app .

# port 9090を開通(8080は、nginxで使う?)
EXPOSE 9090

CMD ["yarn", "serve"]



## memo dockerfileをcomposeを使わずに起動する方法
# ${container_name} = hacku2021-web-test
# [ビルド]
# docker build -t ${container_name} .
# [コンテナ起動]
# docker run -d -p 9090:9090 ${container_name}
# [起動中のコンテナ確認]
# docker ps -a
# (CONTAINER IDが表示される)
# [コンテナ停止]
# docker stop ${CONTAINER ID}
