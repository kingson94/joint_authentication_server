mkdir -p "build"

task=$1
output_file_path="$PWD/build/joint-authen-server"
src_path="$PWD/src"

cd $src_path

if [ $task == "-b" ]
then
	go build -o $output_file_path
	exit 0
fi

if [ $task == "-br" ]
then
	go build -o $output_file_path
	$output_file_path
	exit 0
fi

if [ $task == "-r" ]
then
	$output_file_path
	exit 0
fi