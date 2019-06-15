exercise="a b c d e f"

for a in $exercise
do
    mkdir $a
    cp ../boilerplate/a.go ${a}/${a}.go
    cp ../boilerplate/a_test.go ${a}/${a}_test.go
done
