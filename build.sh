# 停止和删除旧容器
OLD_CONTAINER_ID=`docker container ls -aq --filter name=wechat-official-account`
if [ -n $OLD_CONTAINER_ID ];then
    docker container stop $OLD_CONTAINER_ID
    docker container rm $OLD_CONTAINER_ID
fi
echo "stop and rm old container, done!"


# 删除旧镜像
OLD_IMAGE_ID=`docker image ls wechat-official-account -q`
if [ -n $OLD_IMAGE_ID ]; then
    docker image rm $OLD_IMAGE_ID
fi

echo "rm old image, done!"

# 加载新镜像
docker image load -i wechat-official-account.tar.gz
echo "load new container, done!"

# 启用新容器
IMAGEID=`docker image ls wechat-official-account -q`
docker container run --name wechat-official-account -p 8080:8080 -v /home/work/data/wechat-official-account/logs:/app/logs -d $IMAGEID
echo "run new container done"
