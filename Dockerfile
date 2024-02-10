# Use the official PostgreSQL image as the base image
FROM postgres:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_PASSWORD=pass


# Set the timezone non-interactively
RUN ln -fs /usr/share/zoneinfo/Europe/London /etc/localtime

# Expose the PostgreSQL default port
EXPOSE 5432 

