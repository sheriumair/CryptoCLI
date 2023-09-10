# Use the official Node.js image as the base image
FROM node:14

# Set the working directory inside the container
WORKDIR /app

# Copy your frontend application files
COPY . .

# Install dependencies
RUN npm install

# Expose the port your application will run on
EXPOSE  5173

# Command to start your frontend application
CMD ["npm", "run", "dev"]