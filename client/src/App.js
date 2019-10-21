// @flow strict

import type {Element} from 'react';

import React, {useCallback} from 'react';
import {
  BrowserRouter as Router,
  Link,
  Redirect,
  Route,
  Switch,
} from "react-router-dom";

import HomePage from './home/HomePage';
import LoginPage from './login/LoginPage';
import Page from './component/Page';
import Query from './component/Query';

type Response = {
  is_logged_in: boolean,
};

type Props = {};

function App() {
  const errorRenderer = useCallback(
    (error) => (
      <div>
        Error
      </div>
    ),
    [],
  );
  const successRenderer = useCallback(
    ({is_logged_in: isLoggedIn}: Response): Element<typeof Router> => (
      <Router>
        <Route exact
          path='/login'
          render={({location}): Element<typeof LoginPage | typeof Redirect> => (
            isLoggedIn
              ? <Redirect to={{
                pathname: '/home',
                state: { from: location }
              }} />
              : <LoginPage />
          )} />
        <Route exact
          path='/home'
          render={({location}): Element<typeof HomePage | typeof Redirect> => (
            isLoggedIn
              ? <HomePage />
              : <Redirect to={{
                pathname: '/login',
                state: { from: location }
              }} />
          )} />
        <Redirect to={{
          pathname: '/home',
        }} />
      </Router>
    ),
    [],
  );
  return (
    <Query
      apiPath="/api/check_login"
      errorRenderer={errorRenderer}
      loadingComponent={<Page />}
      successRenderer={successRenderer} />
  );
};

export default React.memo<Props>(App);
