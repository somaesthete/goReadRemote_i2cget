ssh root@192.168.1.132 'PIN=51; echo 1 > /sys/class/gpio/gpio$PIN/value && sleep 0.5 && echo 0 > /sys/class/gpio/gpio$PIN/value'

sleep 0.5

ssh root@192.168.1.132 'PIN=2; echo 1 > /sys/class/gpio/gpio$PIN/value && sleep 0.5 && echo 0 > /sys/class/gpio/gpio$PIN/value'


