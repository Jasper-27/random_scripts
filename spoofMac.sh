#!/bin/sh
newMac=$(hexdump -n3 -e'/3 "a4-83-2F" 3/1 "-%02X"' /dev/random) 
#echo $newMac
ip=$(ifconfig en0 | grep ether)
echo "old MAC address: $ip" 
sudo ifconfig en0 ether $newMac
ip=$(ifconfig en0 | grep ether)
echo "New MAC address: $ip" 