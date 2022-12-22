make STATIC=1
cd build/
echo "FROM scratch" > Dockerfile
echo "ADD paste /paste" >> Dockerfile
echo "CMD /paste" >> Dockerfile
docker build -t paste .
