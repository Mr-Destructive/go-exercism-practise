#exercism configure --workspace=/home/meet/code/playground/go-projects/exercism/practise

exercism download --exercise=$1 --track=go
mv go/$1 .
rm -rf go
