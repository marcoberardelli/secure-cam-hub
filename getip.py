import urllib.request

ip = urllib.request.urlopen('https://api.ipify.org').read().decode('utf8')
print(ip)
