# Use official Node.js v22.14 image as base
FROM node:22-alpine

# Set working directory
WORKDIR /usr/src/app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install --production

# Copy the rest of the application code
COPY . .

# Build the NestJS app (if using TypeScript)
RUN npm run build

# Expose the port the app runs on
EXPOSE 3000

# Start the NestJS app
CMD ["node", "dist/main.js"]