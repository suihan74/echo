FROM golang:alpine as backend-base

ENV USER "echo"
ENV GROUP "echo"
ENV HOME "/echo"

RUN addgroup "$GROUP" && adduser -D -G "$GROUP" "$USER" \
    && mkdir -p "$HOME"\
    && chown -R "$USER:$GROUP" "$HOME"


FROM backend-base as built

WORKDIR "$HOME"
USER "$USER"

COPY --chown="$USER" *.go ./
COPY --chown="$USER" go.mod ./
RUN go build -o backend *.go


FROM backend-base as backend

WORKDIR "$HOME"
USER "$USER"
COPY --from=built --chown="$USER" "$HOME/backend" .
COPY --chown="$USER" entrypoint.sh .
COPY --chown="$USER" *.json .
COPY --chown="$USER" echo-sns-firebase-adminsdk.json .

EXPOSE 8000
CMD ["/echo/entrypoint.sh"]
