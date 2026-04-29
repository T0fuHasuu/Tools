FROM node:22-slim as build


WORKDIR /app

COPY pakage*.json ./
RUN npm install --production 

COPY . .

FROM gcr.io/distroless/nodejs22-debian12
WORKDIR /app

COPY --from=build /app /app

EXPOSE 3000

CMD ["node", "index.js"]