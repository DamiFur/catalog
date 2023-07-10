#!/bin/bash

# usage: sh test.sh [-c] [-v] [-m] [-t <sepcific_test>]
# -c produces coverage report
# -v runs go test with verbose option
# -t <specific_test> runs go test with the option "-run specific_test"

# parses for options -c, -v, -m and -t
function opts() {
	while getopts "t:cvm" opt; do
		case $opt in
			t) specific_test="-run $OPTARG"; ;;
			c) cover="true"; ;;
			v) verbose="-v"; ;;
		esac
	done
}

opts $@

for dir in $(find . -name "*_test.go" -not -path "./vendor/*" |sed 's#\(.*\)/.*#\1#' | sort -u)
do
       cd $dir
       go test $verbose $specific_test
       test_result="$?"
       cd -
       if [ $test_result != 0 ]; then
            break
       fi
done

exit $((test_result))
