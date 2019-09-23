FROM node:carbon as builder
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx
COPY --from=builder /usr/src/app/build/ /usr/share/nginx/html
COPY --from=builder /usr/src/app/default.conf /etc/nginx/conf.d/
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

