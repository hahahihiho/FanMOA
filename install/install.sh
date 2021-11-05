sudo apt install docker.io docker-compose -y

# add User in Group
sudo groupadd docker
sudo usermod -aG docker $USER
su $USER

# # docker 실행
# systemctl start docker 

# golang 설치
sudo apt install golang -y

# fabric-samples 설치
curl -sSL http://bit.ly/2ysbOFE | bash -s -- 1.4.12 1.4.9 0.4.22

# # nodejs 설치
wget https://nodejs.org/dist/v16.13.0/node-v16.13.0-linux-x64.tar.xz
tar -Jxvf node-v16.13.0-linux-x64.tar.xz