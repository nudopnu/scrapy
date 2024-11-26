FROM node:22.11.0 as builder
WORKDIR /app
COPY src/package*.json .
RUN npm install
COPY src .
RUN npm run build:prod -- --base-href ${BASE_HREF} --output-path=dist

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx.conf
EXPOSE 80
CMD [ "nginx", "-g", "daemon off;" ]