#!/bin/ash
NAMA=${NAMA:-"John Doe"}
NIM=${PHP_PORT:-"16515020XXXXXXX"}
DOKUMEN=${DOKUMEN:-"279/UN10.F15.11/PP/2020"}
TIME=${TIME:-"15.00"}

/app/BotSemhas --nama="$NAMA" --nim=$NIM --dokumen=$DOKUMEN --time=$TIME