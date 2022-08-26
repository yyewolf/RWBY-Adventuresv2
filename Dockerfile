FROM node:lts-alpine AS BUILD_FRONT_IMAGE

# définit le dossier 'app' comme dossier de travail
WORKDIR /app

# copie les fichiers et dossiers du projet dans le dossier de travail (par exemple : le dossier 'app')
COPY ./arenas_front/ .

RUN npm install

# construit l'app pour la production en la minifiant
RUN npm run build

FROM golang:1.19 AS BUILD_BACK_IMAGE

WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .
COPY --from=BUILD_FRONT_IMAGE /app/dist /app/arenas_back/static/www/

WORKDIR /app/arenas_back

RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o arenas

FROM alpine

ENV UID 1005
ENV GID 1005

WORKDIR /app

RUN addgroup -S arenas -g $GID && \
    adduser -S arenas -G arenas -u $UID && \
    chown -R arenas:arenas /app
    
USER arenas

COPY --from=BUILD_BACK_IMAGE --chown=arenas:arenas /app/arenas_back/arenas /app/arenas
RUN chmod +x /app/arenas

EXPOSE 50001
EXPOSE 9001

ENTRYPOINT [ "/app/arenas" ]