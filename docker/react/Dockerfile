FROM node:16.0.0

WORKDIR /app/client
COPY . /client

RUN cd /app \
    && yarn start \
    && yarn run build

CMD ["yarn","dev"]