#!/bin/sh
odh-data-monitor d start
/usr/bin/supervisord -n &
