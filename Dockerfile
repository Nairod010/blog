# Use the official PostgreSQL image as the base image
FROM postgres:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=dorian
ENV POSTGRES_PASSWORD=pass
ENV POSTGRES_DB=dbase

# Set the timezone non-interactively (optional)
RUN ln -fs /usr/share/zoneinfo/Europe/London /etc/localtime

# Expose the PostgreSQL default port
EXPOSE 5432 

