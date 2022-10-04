FROM node:lts-alpine AS ARENA_FRONT_IMAGE
WORKDIR /app
COPY ./arenas_front/ .
RUN npm install
RUN npm run build

FROM node:lts-alpine AS DUNGEON_FRONT_IMAGE
WORKDIR /app
COPY ./dungeons_front/ .
RUN npm install
RUN npm run build

FROM node:lts-alpine AS MARKET_FRONT_IMAGE
WORKDIR /app
COPY ./market_front/ .
RUN npm install
RUN npm run build

FROM node:lts-alpine AS OC_FRONT_IMAGE
WORKDIR /app
COPY ./oc_contest_front/ .
RUN npm install
RUN npm run build

FROM golang:1.19 AS BUILD_GO_IMAGE
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .

#### FRONTS
COPY --from=ARENA_FRONT_IMAGE /app/dist /app/arenas_back/static/www/
COPY --from=DUNGEON_FRONT_IMAGE /app/dist /app/dungeons_back/static/www/

#### BACKS
WORKDIR /app/arenas_back
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o arenas -buildvcs=false

WORKDIR /app/dungeons_back
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o dungeons -buildvcs=false

WORKDIR /app/market_back
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o market -buildvcs=false

WORKDIR /app/topgg
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o topgg -buildvcs=false

WORKDIR /app/gambles
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o gambles -buildvcs=false

WORKDIR /app/cdn
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o cdn -buildvcs=false

WORKDIR /app/auth
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o auth -buildvcs=false

WORKDIR /app/oc_contest_back/cmd/server
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o ocback -buildvcs=false

WORKDIR /app/main
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o main -buildvcs=false


### FINAL IMAGE
FROM node:lts-alpine

RUN npm install -g http-server

ENV UID 1005
ENV GID 1005

WORKDIR /app

RUN addgroup -S bot -g $GID && \
    adduser -S bot -G bot -u $UID && \
    chown -R bot:bot /app
    
USER bot

COPY --from=BUILD_GO_IMAGE --chown=bot:bot /app/arenas_back/arenas /app/arenas
RUN chmod +x /app/arenas

COPY --from=BUILD_GO_IMAGE --chown=bot:bot /app/dungeons_back/dungeons /app/dungeons
RUN chmod +x /app/dungeons

COPY --from=BUILD_GO_IMAGE --chown=bot:bot /app/market_back/market /app/market
RUN chmod +x /app/market

COPY --from=BUILD_GO_IMAGE --chown=bot:bot /app/topgg/topgg /app/topgg
RUN chmod +x /app/topgg

COPY --from=BUILD_GO_IMAGE --chown=bot:bot /app/gambles/gambles /app/gambles
RUN chmod +x /app/gambles

COPY --from=BUILD_GO_IMAGE --chown=bot:bot /app/cdn/cdn /app/cdn
RUN chmod +x /app/cdn

COPY --from=BUILD_GO_IMAGE --chown=bot:bot /app/auth/auth /app/auth
RUN chmod +x /app/auth

COPY --from=BUILD_GO_IMAGE --chown=bot:bot /app/oc_contest_back/cmd/server/ocback /app/ocback
RUN chmod +x /app/ocback

COPY --from=BUILD_GO_IMAGE --chown=bot:bot /app/main/main /app/main
RUN chmod +x /app/main

COPY --from=MARKET_FRONT_IMAGE --chown=bot:bot /app/dist /app/market_front/dist
COPY --from=OC_FRONT_IMAGE --chown=bot:bot /app/dist /app/oc_front/dist

# HTTP
EXPOSE 80 
# WS
EXPOSE 81
# MICROSERVICE
EXPOSE 82

# MARKET FRONT
EXPOSE 8080