#局部变量(执⾏⽂件名称), 根据⾃⼰项⽬随便写
project_name="iyyzh"
#杀掉之前正在运⾏的程序
go_id=$(ps -ef | grep "./${project_name}" | grep -v "grep" | awk '{print $2}')
if [ -z "$go_id" ]; then
  echo "[go pid not found]"
else
  #杀掉进程
  kill -9 $go_id
  echo "killed $go_id"
fi
#清除旧的编译⽂件
echo "clean old file"
rm -rf ${project_name}
#执⾏⽇志，根据⾃⼰项⽬情况可选
rm -rf ${project_name}.log
if [ -f main ]; then
  echo "strat new process"
  mv main ${project_name}
  chmod -R 777 ${project_name}
  #这⾥要防⽌nohup不执⾏，添加了⼀个BUILD_ID
  BUILD_ID=DONTKILLME nohup ./${project_name} >${project_name}.log 2>&1 &
  echo "test BUILD_ID=DONTKILLME nohup ./${project_name} >${project_name}.log 2>&1 & over"
else
  echo "executable file not found,quit"
fi
