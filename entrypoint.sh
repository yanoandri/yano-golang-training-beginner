#!/bin/sh

# Start the run once job.
echo "Docker container has been started"

# Setup a cron schedule
echo "* * * * * /myapp/start cron >> /dev/stdout" > /etc/crontabs/root

crontab /etc/crontabs/root
crontab -l