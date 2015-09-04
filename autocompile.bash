#!/bin/bash

watchdir=$(pwd)/src/mipstest

#nano second
function gettimemills {
echo $(date +%s%N | cut -b1-13)
}

lasttime=0

function ratelimit {
  currenttime=$(gettimemills)
  t2=$(expr $(gettimemills) - $lasttime)
  if [ "$t2" -gt "4000" ];then
	lasttime=$currenttime
  	return 0
  fi
  echo "ratelimit ignore request time elpased $t2"
  return 1
}

# requires inotify-tools
# -m keep monitoring
# -r recursive
# -q quiet, print less infos
inotifywait -mrq $watchdir -e delete,modify | while read line
do
	# echo $line
	# split lines
	set -- $line
	dir=$1
	action=$2
	file=$3

	suffix=`echo $file | tail -c 4`

	if [[ $suffix == '.go' ]];then
		if [[ $action == 'MODIFY' ]];then
			echo "$dir $file $action"
			if ratelimit;then
				# 等待文件写入,否则会出现no buildable go source
				sleep 1
				make mips_demo
			fi
		fi
	fi
done
