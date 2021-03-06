#!/bin/bash

test -f .env && source .env

if ! test "$AccountKey"
then
	echo env AccountKey is unset
	exit
fi

pwd=$PWD
cd "$(mktemp -d)" || exit

count=0
while :
do
curl -s -f -X GET http://datamall2.mytransport.sg/ltaodataservice/BusStops/?\$skip=$count -H "AccountKey: $AccountKey" |
	jq .value[] > $count.json
	test -s "$count.json" || break
	count=$((count+500))
done

jq . ./*.json | jq -s . > "$pwd/all.json"
echo "Bus stop count $(jq 'length' "$pwd/all.json")"
