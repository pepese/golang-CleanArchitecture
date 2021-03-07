#!/bin/sh

set -e

host="$1"
shift
user="$1"
shift
password="$1"
shift
cmd="$@"

echo "Waiting for mysql"
until mysql -h"$host" -u"$user" -p"$password" &> /dev/null
do
        >&2 echo -n "."
        sleep 1
done

>&2 echo "MySQL is up - executing command"
exec $cmd