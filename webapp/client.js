const { SignUpRequest } = require('./protobufs/account/account_pb.js');
const { PublicApiPromiseClient: AccountPublicApiPromiseClient } = require('./protobufs/account/account_grpc_web_pb.js');

const accountClient = new AccountPublicApiPromiseClient('https://localhost', null, null);

const accountRequest = new SignUpRequest();
accountRequest.setEmail('qb8@gmail.com');
accountRequest.setPassword('pass');
accountRequest.setUsername('bross');

accountClient.signUp(accountRequest, {})
    .then(response => {
        console.log(response.getMessage());
    })
    .catch(console.error);
