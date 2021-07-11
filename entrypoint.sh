#!/bin/sh

# Start the run once job.
echo "Docker container has been started"

# Setup a cron schedule
echo "* * * * * /myapp/start cron >> /var/log/cron.log 2>&1" > scheduler.txt

# Register cron
crontab scheduler.txt

# Start the cron
/usr/sbin/crond -f -l 8