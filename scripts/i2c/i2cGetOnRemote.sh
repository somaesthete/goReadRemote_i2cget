pullUps=('/usr/sbin/i2cset -y 2 0x20 0x0D 0xFF' '/usr/sbin/i2cset -y 2 0x20 0x0C 0xFF' '/usr/sbin/i2cset -y 2 0x21 0x0C 0xFF' '/usr/sbin/i2cset -y 2 0x21 0x0D 0xFF')

for cmd in "${pullUps[@]}"; do
  echo $cmd
  $cmd
done

while :;
do
  echo $(/usr/sbin/i2cget -y 2 0x20 0x12) $(/usr/sbin/i2cget -y 2 0x20 0x13) $(/usr/sbin/i2cget -y 2 0x21 0x12) $(/usr/sbin/i2cget -y 2 0x21 0x13); sleep 0.1;
done
