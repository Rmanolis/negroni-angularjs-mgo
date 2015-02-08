app.factory('SocketSrv', function (socketFactory) {
  return socketFactory({
    ioSocket: io.connect('http://localhost:5000')
    
  });
})
