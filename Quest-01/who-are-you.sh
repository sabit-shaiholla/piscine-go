#!/bin/bash

abcd=$(curl --silent https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq '.[] | select( .id == 70) | .name')
echo $abcd
