#!/bin/bash

# run http server (8000)
/http_test.py 8000 &

# run http server (8080)
/http_test.py 8080 &

# start apache2
service apache2 start

# infinite loop
/usr/bin/tail -f /dev/null
