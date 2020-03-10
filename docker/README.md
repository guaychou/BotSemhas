##Docker Container Support 

No | Key | Default Value | TYPE
--- | --- | --- | --- |
1 | NAMA | John Doe | String
2 | NIM | 16515020XXXXXXX| String
3 | DOKUMEN | 279/UN10.F15.11/PP/2020| String
4 | TIME | 15.00 | String

How to run
```cassandraql
$ docker run -d --name botsemhas lordchou/botsemhas:1.0.0
```

Adjust name, nim, document, and time with __YOUR__ own.

Example
```cassandraql
$ docker run -d -e NAMA="Smithy Weenie Bittie" -e NIM=165150208293892 -e DOKUMEN=279/UN10.F15.11/PP/2020 \
-e TIME=13.00 --name botsemhas lordchou/botsemhas:1.0.0
```