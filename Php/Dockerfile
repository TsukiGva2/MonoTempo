# Use the official PHP image with Apache
FROM php:8.2-apache

RUN apt-get update && apt-get install -y git supervisor unzip libzip-dev libssl-dev

# Install mysqli and sockets extensions
RUN docker-php-ext-install mysqli sockets zip

# Set the working directory
WORKDIR /var/www

# Copy the composer.json file and install dependencies
COPY ./src/composer.json ./
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" && \
    php composer-setup.php && \
    php -r "unlink('composer-setup.php');" && \
    php composer.phar install

# Copy the rest of the application files
COPY ./src /var/www/html/
COPY ./src/server.php /var/www/server.php


# Copy the supervisor configuration file
COPY ./supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Expose ports for Apache and WebSocket server
EXPOSE 80 8080

# Command to run supervisor
CMD ["/usr/bin/supervisord"]
