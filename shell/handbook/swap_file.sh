#!/bin/bash

if [[ `cat /proc/meminfo | grep SwapTotal | awk '{ print $2 }'` != 0 ]]; then
  print "Swap is already enabled!"
  free -h | grep Swap
else
  print "Enter with number of gigabytes that you want to create [4]: "
  read answer
  if [ -z "$answer" ]; then
    answer=4
  fi
  SWAP_SIZE=$(echo $answer | grep -o [0-9] | head -n1)
  COUNT=$(echo "${SWAP_SIZE}*1024/128" | bc)
  print "Creating swap file..."
  sudo dd if=/dev/zero of=/var/swap.img bs=128M  count=${COUNT}
  sudo chmod 600 /var/swap
  sudo mkswap /var/swap
  sudo swapon /var/swap
  sudo sed -i '$a # SwapFile\n/var/swap    none     swap    sw    0    0' /etc/fstab

  print "SwapFile has been created successfuly!"
  free -h | grep Swap
fi
