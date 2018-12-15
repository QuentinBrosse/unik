const { HelloRequest } = require('./protobufs/api_pb.js');
const { ApiPromiseClient } = require('./protobufs/api_grpc_web_pb.js');

const client = new ApiPromiseClient('https://localhost', null, null);

const request = new HelloRequest();
request.setName('you');

client.sayHello(request, {})
    .then(response => {
        console.log(response.getMessage());
    })
    .catch(console.error);
