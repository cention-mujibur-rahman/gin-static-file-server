# gin-static-file-server
gin framework static resource serve by web server
It will help this thread where some guys still asking example, [https://github.com/gin-gonic/gin/issues/75].

I took the help from manucorporat: [https://github.com/manucorporat]

# Run the following command to run
* git clone git@github.com:cention-mujibur-rahman/gin-static-file-server.git web
* go get github.com/gin-gonic/gin
* go get github.com/gin-gonic/contrib/static
* go build 
* ./web <-- It will start go web server on port :8080

# Open a browser 
[http://localhost:8080/index]

It will run including all the resources serve by web server through gin middleware.

Thanks