# -----
# 1st stage: prepare unprivileged user
# -----
FROM node:10.21.0-alpine3.9 as frontend-base

ENV USER "echo"
ENV GROUP "echo"
ENV HOME "/echo"

RUN addgroup "$GROUP" && adduser -D -G "$GROUP" "$USER" \
    && mkdir -p "$HOME" \
    && chown -R "$USER:$GROUP" "$HOME"

# -----
# 2nd stage: install Node.js modules
# -----
FROM frontend-base as built
WORKDIR "$HOME"
USER "$USER"

COPY --chown="$USER" package.json .
COPY --chown="$USER" package-lock.json .
RUN npm install

# -----
# 3rd stage: set up application
# -----
FROM frontend-base as frontend

ENV PORT 8080
EXPOSE "$PORT"

WORKDIR "$HOME"
USER "$USER"

COPY --from=built --chown="$USER" "$HOME/node_modules" "$HOME/node_modules"
COPY --chown="$USER" . .

CMD ["/echo/entrypoint.sh"]
