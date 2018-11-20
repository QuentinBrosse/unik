const { HelloRequest } = require('./protobufs/api_pb.js');
const { ApiPromiseClient } = require('./protobufs/api_grpc_web_pb.js');

const client = new ApiPromiseClient('http://localhost:8010', null, null);

// simple unary call
const request = new HelloRequest();
request.setName('toi');

client.sayHello(request, {})
    .then(response => {
        console.log(response.getMessage());
    })
    .catch(console.error);
