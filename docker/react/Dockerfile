FROM node:18.0.0

COPY ./client /app/client
WORKDIR /app/client

RUN cd /app/client \
    && yarn install \
    && yarn run build

CMD ["yarn","dev"]