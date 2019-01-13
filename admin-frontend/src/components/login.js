import gql from 'graphql-tag';
import React from 'react';
import { Mutation } from 'react-apollo';
import { Redirect } from 'react-router-dom';
import { Button, Card, Form, Header, Message } from 'semantic-ui-react';
import styled from 'styled-components';

const HorizontalCentered = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const LoginCard = styled(Card)`
  &&& {
    padding: 16px;
    margin-top: 32px;
  }
`;

const LoginButton = styled(Button).attrs({
  primary: true
})`
  &&& {
    width: 100%;
  }
`;

const ErrorMessage = styled(Message).attrs({
  negative: true
})`
  margin-top: 32px;
`;

const LOGIN_USER_MUTATION = gql`
  mutation Login($username: String!, $password: String!) {
    authenticate(username: $username, password: $password) {
      token
    }
  }
`;

class Login extends React.Component {
  state = {
    username: '',
    password: ''
  };

  render() {
    if (localStorage.getItem('token')) {
      return <Redirect to="/" />;
    }

    return (
      <Mutation
        mutation={LOGIN_USER_MUTATION}
        onCompleted={({ authenticate: { token } }) => {
          localStorage.setItem('token', token);
          this.props.history.push('/');
        }}
      >
        {(authenticate, { loading, error }) => (
          <HorizontalCentered>
            <LoginCard>
              <Header as="h2">Hashbeam Admin</Header>
              <Form
                onSubmit={e => {
                  e.preventDefault();
                  authenticate({
                    variables: {
                      username: this.state.username,
                      password: this.state.password
                    }
                  });
                }}
              >
                <Form.Field>
                  <input
                    placeholder="Username"
                    value={this.state.username}
                    onChange={e => this.setState({ username: e.target.value })}
                  />
                </Form.Field>
                <Form.Field>
                  <input
                    type="password"
                    placeholder="Password"
                    value={this.state.password}
                    onChange={e => this.setState({ password: e.target.value })}
                  />
                </Form.Field>
                <LoginButton type="submit" loading={loading}>
                  Login
                </LoginButton>
              </Form>
            </LoginCard>
            {error && (
              <ErrorMessage>
                {error.graphQLErrors
                  ? error.graphQLErrors[0].message
                  : 'Network Error'}
              </ErrorMessage>
            )}
          </HorizontalCentered>
        )}
      </Mutation>
    );
  }
}

export default Login;
