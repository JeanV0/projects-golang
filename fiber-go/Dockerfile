# Eu me baseio nesse link mas algumas coisas mudei para meus propositos https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

ARG arch=AMD64

FROM golang:1.19-alpine as builder

ENV USER=appuser
ENV UID=10001

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

RUN chown userapp:userapp -R "/home/${USER}/"

WORKDIR /srv/app

USER appuser:appuser

COPY ./ ./

RUN go install

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /srv/app/app

FROM scratch as runner

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /srv/app

COPY --from=builder /srv/app/.env.local /srv/app/.env.local
COPY --from=builder /srv/app/app /srv/app/app

USER appuser:appuser

EXPOSE 8080

CMD ["/srv/app/app"]