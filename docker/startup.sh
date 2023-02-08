#!/bin/sh
apisense daemon start --background
/usr/bin/supervisord -n
