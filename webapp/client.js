const { SignUpRequest } = require('./protobufs/account/account_pb.js');
const { PublicApiClient } = require('./protobufs/account/account_grpc_web_pb.js');

const accountClient = new PublicApiClient('https://localhost', null, null);

const accountRequest = new SignUpRequest();

const email = prompt("What is your email dude?");
const username = prompt("What is your username dude?");

accountRequest.setEmail(email);
accountRequest.setPassword('pass');
accountRequest.setUsername(username);

const signUpCall = accountClient.signUp(accountRequest, {}, console.log);

signUpCall.on('status', console.log('status', status));
