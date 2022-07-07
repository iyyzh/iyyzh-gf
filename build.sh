project_name="iyyzh"
go_id=$(ps -ef | grep "./${project_name}" | grep -v "grep" | awk '{print $2}')
if [-z "$go_id"]; then
  echo "[go pid not found]"
else
  kill -9 $go_id
  echo "killed $go_id"
fi
echo "clean old file"
rm -rf ${project_name}
rm -rf ${project_name}.log
if [-f main]; then
  echo "strat new process"
  mv main ${project_name}
  chmod-R 777 ${project_name}
  BUILD_ID=DONTKILLME nohup./${project_name} >${project_name}log 2>&1 &
else
  echo "executable+file+not+found,quit"
fi
