#!/bin/sh
odh-data-monitor daemon start --background
/usr/bin/supervisord -n
