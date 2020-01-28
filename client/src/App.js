// @flow strict

import type { Element } from 'react';

import React, { useCallback, useState } from 'react';
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

import { useFetchInitialData } from './api/useFetchInitialData';

type Response = {
  is_logged_in: boolean,
};

type Props = {};

function App() {
  const [submitting, setSubmitting] = useState(false);
  const startLoginSubmit = useCallback((): void => (
    setSubmitting(true)
  ), []);
  const finishLoginSubmit = useCallback((): void => (
    setSubmitting(false)
  ), []);
  const {
    data,
    error,
    isLoading,
  } = useFetchInitialData<Response>('/api/check_login', finishLoginSubmit, [submitting]);
  if (isLoading || data == null) {
    return <Page />;
  } else if (error != null) {
    return <div>{error.message}</div>;
  }
  const { is_logged_in: isLoggedIn } = data;
  return (
    <Router>
      <Switch>
        <Route exact
          path='/login'
          render={({ location }): Element<typeof LoginPage | typeof Redirect> => (
            isLoggedIn
              ? <Redirect to={{
                pathname: '/home',
                state: { from: location }
              }} />
              : <LoginPage onSubmit={startLoginSubmit} />
          )} />
        <Route exact
          path='/home'
          render={({ location }): Element<typeof HomePage | typeof Redirect> => (
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
      </Switch>
    </Router>
  );
};

export default React.memo<Props>(App);
