#! /bin/zsh

zsh test.sh 100 > input
go run sum_of_squares2.go < input > gooutput
diff gooutput output > tmp.diff

if [ $(echo tmp.diff | wc - l) -gt 0  ]; then
    go run post_hennge_mission_3.go
else
    echo "Error!"
fi
