#! /bin/zsh

num_of_test=$1
echo $num_of_test 

rm -f output
for ((i=0; i<$num_of_test; i++))
do
    x=$((1 + $RANDOM % 100))
    echo $x

    rm -f temp.y
    sum=0
    for ((j=0; j<$x; j++))
    do
        max=100
        min=-100
        y=$(( RANDOM % ( $max - $min + 1 ) + $min ))
        if [ $y -gt 0 ]; then
            let "sum=sum+y*y"
        fi
        printf "${y} " >> temp.y
    done
    echo $sum >> output
    cat temp.y | xargs
    rm -f temp.y

done




