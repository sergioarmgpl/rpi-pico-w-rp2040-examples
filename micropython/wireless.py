#examples from https://www.coderdojotc.org/micropython/basics/06-wireless/
#Sensor Part
import dht
from machine import Pin
import time
#Wireless
import network
from utime import sleep
import urequests


secrets = {"SSID":"waifai","PASSWORD":"demo123-"}
ssid = secrets["SSID"]
password = secrets["PASSWORD"]

wlan = network.WLAN(network.STA_IF)
wlan.active(True)
wlan.connect(ssid, password)

# Wait for connect or fail
max_wait = 10
while max_wait > 0:
  if wlan.status() < 0 or wlan.status() >= 3:
    break
  max_wait -= 1
  print('waiting for connection...')
  time.sleep(1)

# Handle connection error
if wlan.status() != 3:
   raise RuntimeError('network connection failed')
else:
  print('connected')
  status = wlan.ifconfig()
  print( 'ip = ' + status[0] )
  

astronauts = urequests.get("http://api.open-notify.org/astros.json").json()
number = astronauts['number']
print('There are', number, 'astronauts in space.')

d = dht.DHT11(machine.Pin(28))

while True:
    d.measure()
    t=d.temperature()
    h=d.humidity()
    time.sleep(1)
    print('{"t":'+str(t)+',"h":'+str(h)+'}')

