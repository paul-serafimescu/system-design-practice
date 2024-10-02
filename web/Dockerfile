# FROM node:20-alpine

# WORKDIR /app

# COPY package*.json ./

# RUN npm install

# COPY . .

# RUN npm run build

# RUN npm install -g serve

# EXPOSE 3000

# CMD ["serve", "-s", "dist"]

# above is using builder: this is decent for prod but for dev, npm install and npm run build take forever

FROM node:20-alpine

WORKDIR /app

COPY dist /app/dist

RUN npm install -g serve

EXPOSE 3000

CMD ["serve", "-s", "dist"]

