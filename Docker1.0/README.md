


Monolatic vs MIcroService



# Monolatic 

   Single Code base
   single war appication
   sigle Server



Maintainance is very less
loadbalace is very Risk














# Docker
  
# Virtual machine vs  Docker Container


1. installion and configuration is diffrent
1. time Consumming 
2. RAM and Cpu use More
3. Version will be same if different not Working
4. Peformance will be More



light weight
less cpu and Ram
Version will be allloed
build in os



Challenges
-----------


1. build and deployment we take care about thats
3. infrasturue containeration
Configuration Everything  in container

2. communication with Microservices



# Virtual machine
      1. big impact on os
      slow
      Higher disk space usage

      sharing and Re-building distribution is  dificult
      3.Encapsulate "whole machines"
      intend of "just app/environment"

# Docker Container
     1. Low impact on os
      very fast
      Minimal  disk space usage

      2. sharing and Re-building distribution is easy

      3. Encapsulate apps/environments
      intend of "whole machines"


# Docker
      Docker Containration tool
      Docker Create Multiple Docker Containers
      we can pull any sytem

      docker are lightweight containers
      deployment are easy and easy to containersation tool










How to Create my own Network:


Docker network ls



custom network will commicate with names and ports

default network will commicate with only port


# How to Create a new network
## docker network -d bridge networkname



# How to chek network one container to communicate with other containers

## docker exec -it dockercontainername bash




How to remove
docker  rm -f $(docker ps -aq)

Remove images from the docker 
docker rmi -f $(docker images -q)
docker rmi $(docker images -q)

How to create a container
docker run -d --name javawebapp -p --network flipcartnet 8080:8080 dockerusername/javaweb-app




docker run -d --name 



Dockerfile--------->Image--------container
      create a Image   push to repo


### docker build 

            docker build -t <registry/repo:tag> <dockerBuildContext>

            docker build -t vinaydocker/appname:1 .

            docker build -t 179.283.282.2823:8080/appname:1 .

            docker build -t ecr.123333.aws.com/app:1 .

### list images
       list all the iamges
            docker images or 
            docker images ls


### docker push 
     my images push to docker repository
            docker push vinaydocker/appname:1

### docker pull
      its will copy the image in docker repo
            docker pull vinaydocker/appname:1



### docker login
      if we need to push to docker repo 
      we need to login
            docker login -u vinaydocker
            password:***
            private login 
            docker login -u admin123 132.12.12.12:4044



### more details of docker images 
        docker image inspect <imageid/name>
        or
        docker insert <imageid/name

        example:
        docker image inspect vinaydocker/imagename:1
        docker inspect vinaydocker/imagename:1

### docker delete images
            remove the docker image
      
            docker rmi <imageid/name>
            docker rmi $(docker imaages -q)



## docker build

     it will create a docker image
      docker build -t dockerhandsom/maven-web:2 -f Dockerfile ~/build-context/

       example
      docker build -t dockerhandsom/maven-web:2 -f Dockerfile ~/maven-web-application/


      <!--  -->

      docker build -t dockerhandsom/maven-web:2 .

### image layer

   docker histroy imageid/iamged
  sudo /var/lib/docker/
  working of docker

  sudo ls /var/lib/docker/overlay2


# History,inspect
            docker image
                --list all the images
            docker History image
                 --chek the History image layer
            docker insect image
               ---rearding images



# tomcat

tomecat contain other image and other iamge depend on other image


if we execute same base image its take small time
in all conditions like push pull  


### How to Log out
      docker logout


### How to delecte images
     docker rmi <imageid/iamges:1>


### How to delect all images

      docker iamges -q
             its will list all imaages id
      docker rmi -f $(docker iamges -q)

#### How to delete image if there is running container
     


### How to list all containers
            docker container ls


### How to create a container:
      docker run -d --name mywebapp -p 8080:8080 imageid/imagename

### what is dangling image??
    Am image which doest have repo details.which in not tagged

    docker images -f dangling=true

    How to tag again?
    docker tag iamgeid  imageid/imagename

### How to remove the container?
       docker rm -f containername



### othe commends
            docker system prune
                    --- its will remove   
                    - all stopped containers
                    - all networks not used by at least one container
                    - all dangling images
                    - all dangling build cache
            docker image prune
                    
            docker container prune
            docker volume prune
### what is registery
    its collections of repositories
    collections if different version of same image  repository version diffrent

if we have ecr:
      











