./test -profiler="localhost:7070" -name=main -netdebug=1 -networkPort=7766 -peers=192.168.1.49:7777 > test-main.out &
./test -profiler="localhost:7071" -name=second -netdebug=1 -networkPort=7777 -peers=192.168.1.49:7766 > test-second.out &
