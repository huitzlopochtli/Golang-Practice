#! /bin/zsh

zsh test.sh 100 > input
go run sum_of_squares.go < input > gooutput
diff gooutput output > tmp.diff

if [ $(cat tmp.diff | wc -l | sed -e "s/ //g") > 0  ]; then
    go run post_hennge_mission_3.go
else
    echo "Error!"
fi
