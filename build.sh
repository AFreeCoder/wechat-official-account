###
 # @Description: 构建脚本
 # @Date: 2021-04-19 23:23:09
 # @LastEditors: wanghaijie01
 # @LastEditTime: 2021-04-20 00:37:53
### 

# prepare
HOMEDIR=`pwd`
OUTPUT=${HOMEDIR}/output
SERVERNAME=wechat-server
RUNMODE=debug

if [ "$1" = "release" ]; then
RUNMODE="$1"
fi

# init go env
go env -w GO111MODULE="on"
go env -w GOPROXY="https://goproxy.cn,direct"

# build
go build -o ${HOMEDIR}/${SERVERNAME}

# build dir
rm -rf ${OUTPUT}
mkdir -p ${OUTPUT}/bin
mkdir -p ${OUTPUT}/conf
cp ${HOMEDIR}/${SERVERNAME} ${OUTPUT}/bin/
cp -r conf/* ${OUTPUT}/conf/

# set run env
cat conf/httpserver.toml | sed -E "s/^run_mode.*/run_mode = \"${RUNMODE}\"/g" > ${OUTPUT}/conf/httpserver.toml

# tar
tar -czPf output.tar output

# clean
rm -rf ${OUTPUT}
rm -rf ${HOMEDIR}/${SERVERNAME}
