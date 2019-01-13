import React from 'react';
import { Route, Switch } from 'react-router-dom';
import Login from './components/login';
import PrivateRoute from './components/private-route';

const App = () => (
  <Switch>
    <PrivateRoute path="/" exact component={() => null} />
    <Route path="/login" exact component={Login} />
  </Switch>
);

export default App;
