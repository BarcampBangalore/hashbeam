import ApolloClient from 'apollo-boost';
import React from 'react';
import { ApolloProvider } from 'react-apollo';
import { render } from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import 'semantic-ui-css/semantic.min.css';
import App from './app';
import PageContainer from './components/page-container';
import config from './config.json';

const client = new ApolloClient({
  uri: config.server_url,
  request: async operation => {
    const token = localStorage.getItem('token');
    const headers = {};
    if (token) {
      headers.Authorization = token;
    }
    operation.setContext({ headers });
  },
  onError: err => {
    console.error(err);
  }
});

const Root = () => (
  <ApolloProvider client={client}>
    <BrowserRouter>
      <PageContainer>
        <App />
      </PageContainer>
    </BrowserRouter>
  </ApolloProvider>
);

render(<Root />, document.getElementById('root'));
