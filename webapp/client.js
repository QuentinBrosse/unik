const { HelloRequest } = require('./protobufs/api/api_pb.js');
const { SignUpRequest } = require('./protobufs/auth/auth_pb.js');
const { PublicApiPromiseClient: ApiPublicApiPromiseClient } = require('./protobufs/api/api_grpc_web_pb.js');
const { PublicApiPromiseClient: AuthPublicApiPromiseClient } = require('./protobufs/auth/auth_grpc_web_pb.js');


const apiClient = new ApiPublicApiPromiseClient('https://localhost', null, null);

const apiRequest = new HelloRequest();
apiRequest.setName('you');

apiClient.sayHello(apiRequest, {})
    .then(response => {
        console.log(response.getMessage());
    })
    .catch(console.error);

const authClient = new AuthPublicApiPromiseClient('https://localhost', null, null);

const authRequest = new SignUpRequest();
authRequest.setEmail('qb8@gmail.com')
authRequest.setPassword('pass')
authRequest.setUsername('bross');

authClient.signUp(authRequest, {})
    .then(response => {
        console.log(response.getMessage());
    })
    .catch(console.error);
