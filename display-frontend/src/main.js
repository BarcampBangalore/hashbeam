import { InMemoryCache } from 'apollo-cache-inmemory';
import { ApolloClient } from 'apollo-client';
import { split } from 'apollo-link';
import { HttpLink } from 'apollo-link-http';
import { WebSocketLink } from 'apollo-link-ws';
import { getMainDefinition } from 'apollo-utilities';
import Vue from 'vue';
import VueApollo from 'vue-apollo';
import App from './components/App';
import config from './config.json';

const httpLink = new HttpLink({
  uri: config.server_url
});

const wsLink = new WebSocketLink({
  uri: config.server_url.replace(new RegExp('https://'), 'wss://'),
  options: {
    reconnect: true
  }
});

const link = split(
  // split based on operation type
  ({ query }) => {
    const { kind, operation } = getMainDefinition(query);
    return kind === 'OperationDefinition' && operation === 'subscription';
  },
  wsLink,
  httpLink
);

const apolloClient = new ApolloClient({
  link,
  cache: new InMemoryCache(),
  connectToDevTools: true
});

const apolloProvider = new VueApollo({
  defaultClient: apolloClient
});

Vue.use(VueApollo);

new Vue({
  el: '#app',
  apolloProvider,
  render: h => h(App)
});
