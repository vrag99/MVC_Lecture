# express-server/Dockerfile
FROM node:18-alpine

WORKDIR /app

# Create package.json for the Express server
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy server code
COPY server.js .

# Expose port
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3000/health || exit 1

# Start the server
CMD ["node", "server.js"]
