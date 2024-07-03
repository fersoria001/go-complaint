# Use the official Node.js 22 image
FROM node:22-alpine AS build

# Set the working directory
WORKDIR /app

# Copy your project files
COPY . /app

# Install dependencies
RUN npm install

# Build your project
RUN npm run build

# Use a multi-stage build with a slimmer image for production
FROM nginx:stable-alpine

# Expose port 80
EXPOSE 80

# Empty the default HTML directory before copying
RUN rm -rf /usr/share/nginx/html

# Copy the built application from the build stage
COPY --from=build /app/dist /usr/share/nginx/html

# Copy the custom configuration file (replace with your actual path)
COPY default.conf /etc/nginx/conf.d/default.conf

# Start Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]