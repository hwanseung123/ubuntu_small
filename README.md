Docker Installation on Linux

https://github.com/WCSCourses/index/blob/main/Docker_guide.md

=> to avoid using **sudo** for docker commands

# add the user to the docker group 
1. sudo usermod -aG docker $USER
# changed group permissions will take effect immediately
2. newgrp docker
