#!/bin/sh
/usr/bin/supervisord -n &
odh-data-monitor d start