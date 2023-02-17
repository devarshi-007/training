
#!bin/bash
for i in {0..10}
  do
    echo "Hi:" $i
  done

for j in {0..10..2}
  do
    echo  $j
  done

read a
if test $a -lt 5;then
	echo "a less 5"
else
	echo "a grete 5"
fi

read z
	case $z in
		1|2 ) echo "1|2" $z
	   	;;
		3|4 ) echo "3|4" $z
    		;;
    		* ) echo "sorry"
	    	;;
  	esac

declare -i counter
counter=10
	while [ $counter -gt 2 ]; do
		echo The counter is $counter
		counter=counter-1
	done
exit 0
