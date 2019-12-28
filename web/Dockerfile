FROM node:12.13.0-alpine

WORKDIR /app

COPY package*.json /app/

RUN npm install

COPY . .

EXPOSE 8080

# USER zafer

CMD [ "npm", "run", "serve" ]
