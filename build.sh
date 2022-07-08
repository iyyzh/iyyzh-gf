#局部变量(执⾏⽂件名称), 根据⾃⼰项⽬随便写
project_name="iyyzh"
#杀掉之前正在运⾏的程序
go_id=$(ps -ef | grep "./iyyzh" | grep -v "grep" | awk '{print $2}')
if [ -z "$go_id" ]; then
  echo "[go pid not found]"
else
  #杀掉进程
  kill -9 $go_id
  echo "killed $go_id"
fi
#清除旧的编译⽂件
echo "clean old file"
rm -rf iyyzh
#执⾏⽇志，根据⾃⼰项⽬情况可选
rm -rf iyyzh.log
if [ -f main ]; then
  echo "strat new process"
  mv main iyyzh
  chmod -R 777 iyyzh
  #这⾥要防⽌nohup不执⾏，添加了⼀个BUILD_ID
#  BUILD_ID=DONTKILLME nohup ./${project_name} >${project_name}.log 2>&1 &
#  nohup ./${project_name} >${project_name}.log 2>&1 &
  ./iyyzh
  echo "test ${project_name} over"
else
  echo "executable file not found,quit"
fi
