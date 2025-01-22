FROM node:18 as dev
WORKDIR /app

# Install dependencies
COPY package.json ./
RUN npm install

COPY . .

EXPOSE 5173
CMD ["npm", "run", "dev"]