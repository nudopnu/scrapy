FROM node:23-alpine3.20 as builder
WORKDIR /app
COPY src/package*.json .
RUN npm install
COPY src .
RUN npm run build:prod -- --output-path=dist

FROM nginx:alpine3.20
COPY --from=builder /app/dist/browser /usr/share/nginx/html
COPY nginx.conf /etc/nginx.conf
EXPOSE 80
CMD [ "nginx", "-g", "daemon off;" ]