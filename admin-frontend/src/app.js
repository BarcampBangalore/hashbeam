import React from 'react';
import { Route, Switch } from 'react-router-dom';
import Login from './components/login';
import MainScreen from './components/main-screen';
import PrivateRoute from './components/private-route';

const App = () => (
  <Switch>
    <PrivateRoute path="/" exact component={MainScreen} />
    <Route path="/login" exact component={Login} />
  </Switch>
);

export default App;
