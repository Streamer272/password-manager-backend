# Password Manager Backend

### Setup DB
`sudo -S docker-compose -f db.yml up -d`
### Build container
`sudo -S docker build -t password-manager-backend_image .`
### Run container from image
`sudo -S docker run --net host -it --name password-manager-backend password-manager-backend_image`
<br>
If you want to run it in detached mode, add `-d` flag.
<br>
If container already exists, run it with `sudo docker start -a -i password-manager-backend`
<br>
If you don't want to start it in attached mode, remove `-a` flag.
