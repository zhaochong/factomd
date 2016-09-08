
echo "starting main"
./test -profiler="localhost:7070" -name=main -netdebug=1 -networkPort=7766 -peers=127.0.0.1:7777 > test-main.out &
echo "starting second"
./test -profiler="localhost:7071" -name=second -netdebug=1 -networkPort=7777 -peers=127.0.0.1:7766 > test-second.out &
echo "sleeping 2 minutes"

sleep 120
pkill test
