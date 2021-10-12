# Hello fresh code challenge

## How to run back-end

```
#cd back-end
#make run-docker
```
It will run a PostgreSQL docker image and it will link it to our backend deployed in another docker image

If you want to run it without docker, you must define on your /etc/host a host called "database" with the postgreSQL address
```
##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1       database
```
And then, execute:

```
#cd back-end
#make run
```

## How to run front-end

You can run it locally by:

```
#cd front-end
#yarn serve
```

Otherwise, you can enjoy the lastest version deployed by Github actions, on [https://gtrrz-victor.github.io/hf](https://gtrrz-victor.github.io/hf)

# Solution Design

You can find the app design [here](./design.md)
