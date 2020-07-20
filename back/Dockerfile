# -----
# 1st stage: prepare unprivileged user
# -----
FROM golang:alpine as backend-base

ENV USER "echo"
ENV GROUP "echo"
ENV HOME "/echo"

RUN addgroup "$GROUP" && adduser -D -G "$GROUP" "$USER" \
    && mkdir -p "$HOME"\
    && chown -R "$USER:$GROUP" "$HOME"

# -----
# 2nd stage: build executable from Go codes
# -----
FROM backend-base as built

WORKDIR "$HOME"
USER "$USER"

COPY --chown="$USER" *.go ./
COPY --chown="$USER" go.mod ./
RUN go build -o backend *.go

# -----
# 3rd stage: set up application
# -----
FROM backend-base as backend

ENV PORT 8080
EXPOSE "$PORT"

WORKDIR "$HOME"
USER "$USER"

COPY --from=built --chown="$USER" "$HOME/backend" "$HOME/backend"
COPY --chown="$USER" entrypoint.sh .
COPY --chown="$USER" echo-sns-firebase-adminsdk.json .

CMD ["/echo/entrypoint.sh"]
