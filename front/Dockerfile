FROM node:alpine as frontend-base

ENV USER "echo"
ENV GROUP "echo"
ENV HOME "/echo"

RUN addgroup "$GROUP" && adduser -D -G "$GROUP" "$USER" \
    && mkdir -p "$HOME" \
    && chown -R "$USER:$GROUP" "$HOME"


FROM frontend-base as built
WORKDIR "$HOME"
USER "$USER"

COPY --chown="$USER" package.json .
COPY --chown="$USER" package-lock.json .
RUN npm install


FROM frontend-base as frontend

WORKDIR "$HOME"
USER "$USER"

COPY --from=built --chown="$USER" "$HOME/node_modules" "$HOME/node_modules"
COPY --chown="$USER" . .

ENV PORT 8080
EXPOSE "$PORT"
CMD ["/echo/entrypoint.sh"]
