import React from 'react';
import ReactDOM from 'react-dom';
import { SignUpRequest } from './protobufs/account/account_pb.js';
import { PublicApiClient } from './protobufs/account/account_grpc_web_pb.js';
// const { SignUpRequest } = require('./protobufs/account/account_pb.js');
// const { PublicApiClient } = require('./protobufs/account/account_grpc_web_pb.js');

class AccountZone extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            account: null,
            error: null,
        };

        const savedAccount = window.localStorage.getItem('account');
        if (savedAccount) {
            this.state.account = JSON.parse(savedAccount);
        }

        this.client = new PublicApiClient('https://localhost', null, null);
        this.signUp = this.signUp.bind(this);
        this.signIn = this.signIn.bind(this);
        this.signOut = this.signOut.bind(this);
    }

    componentWillUnmount() {
        this.client.Close();
    }

    signUp() {
        const request = new SignUpRequest();

        const email = prompt('What is your email dude?');
        const username = prompt('What is your username dude?');
        const password = prompt('What is your password dude?');

        request.setEmail(email);
        request.setUsername(username);
        request.setPassword(password);

        const signUpCall = this.client.signUp(request, {}, (err, res) => {
            if (err) {
                console.error('sign up error:', err);
                this.setState({ error: err });
                return;
            }
            const account = res.getAccount().toObject();
            this.setState({ account });
            window.localStorage.setItem('account', JSON.stringify(account));
            console.log('signed up:', res);
        });

        // signUpCall.on('status', status => console.log('status', status));
        // signUpCall.on('data', d => console.log('data', d));
        // signUpCall.on('end', d => console.log('end', d));
        // signUpCall.on('error', d => console.log('error', d));
    };

    signIn() {
        const request = new SignUpRequest();

        const email = prompt('What is your email dude?');
        const password = prompt('What is your password dude?');

        request.setEmail(email);
        request.setPassword(password);

        this.client.signIn(request, {}, (err, res) => {
            if (err) {
                console.error('sign in error:', err);
                this.setState({ error: err });
                return;
            }
            const account = res.getAccount().toObject();
            this.setState({ account });
            window.localStorage.setItem('account', JSON.stringify(account));
            console.log('signed in:', res);
        });
    }

    signOut() {
        localStorage.removeItem('account');
        this.setState({ account: null });
    }

    render() {
        const { account, error } = this.state;

        return (
            <Container error={error}>
                {!account ? (
                    <div>
                        <button onClick={this.signUp}>
                            Sign Up
                        </button>
                        <button onClick={this.signIn}>
                            Sign In
                        </button>
                    </div>
                ) : (
                    <div>
                        <div>
                            <button onClick={this.signOut}>
                                Sign Out
                            </button>
                            <Account account={account} />
                        </div>
                    </div>
                )}
            </Container>
        );
    }
}

function Container({ error, children }) {
    return (
        <div>
            {error && (
                <div style={{color: 'red'}}>
                    Error: {error.message} (code: {error.code});
                </div>
            )}
            {children}
        </div>
    )
}

function Account({ account }) {
    console.log(account);
    return (
        <div>
            <h1>Connected to: {account.username}</h1>
        </div>
    )
}

const accountRoot = document.querySelector('#accountRoot');
ReactDOM.render(<AccountZone />, accountRoot);
