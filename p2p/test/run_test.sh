#!/usr/bin/env bash

echo "This sets up vagrant boxes, builds factomd, installs it, and then runs it."

CWD=`pwd`
DEST="$CWD/bin/factomd"
echo "changing directory to factomd"

cd "$GOPATH/src/github.com/FactomProject/factomd"

echo "Building linux factomd and putting it in $DEST"
CGO_ENABLED=0 GOOS=linux go build -ldflags "-X github.com/FactomProject/factomd/engine.Build=`git rev-parse HEAD`" -installsuffix cgo -o $DEST
if [ $? -eq 0 ]; then
  echo "was binary updated? Current:`date`"
  ls -G -lh $DEST

  echo "changing directory to back to Vagrant Root ( $CWD )"
  cd $CWD

  # echo "Start the Vagrant boxes"
  # vagrant reload --provision
  vagrant up
  
  echo "About to delete .factom"
  vagrant ssh leader -c "rm -rf ~/.factom"
  vagrant ssh follower -c "rm -rf ~/.factom"
  echo "About to create .factom"
  vagrant ssh leader -c "mkdir ~/.factom"
  vagrant ssh follower -c "mkdir ~/.factom"
  
  echo "About to create .factom/m2"
  vagrant ssh leader -c "mkdir ~/.factom/m2"
  vagrant ssh follower -c "mkdir ~/.factom/m2"

  echo "Start the leader"
  vagrant ssh leader -c "cd /vagrant/bin/ && ./leader.sh" 

  vagrant ssh leader -c "nohup /vagrant/bin/factomd -peers=\"10.0.99.2:8110\" -networkPort=8110 -network=LOCAL -blktime=20 -netdebug=1 -exclusive=true >> /vagrant/output/leader.out 2>&1 & "

  # echo "Start the wallet on leader"
  # vagrant ssh leader -c "cd /vagrant/bin/ && nohup ./fctwallet > /vagrant/output/leader-wallet.out 2>&1 &"  

  # echo "Sleep while waiting for the leader to make 12 blocks."
  # sleep 240

  # echo "Add entries"
  # vagrant ssh leader -c "cd /vagrant/bin/ && ./entries.sh"  

  # echo "Sleep while waiting for the leader to make 6 blocks."
  # sleep 120

  # echo "Start the follower"
  # vagrant ssh follower -c "cd /vagrant/bin/ && ./follower.sh"

  # while true ;
  # do
  #   echo "Block Heights. CTRL-C to quit."
  #   echo "Leader:"
  #   vagrant ssh leader -c "/vagrant/bin/factom-cli get height"  
  #   echo "Follower:"
  #   vagrant ssh follower -c "/vagrant/bin/factom-cli get height"  
  #   sleep 5
  # done

fi