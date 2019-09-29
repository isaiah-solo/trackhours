// @flow

import React, {useCallback} from 'react';

import HomePage from './home/HomePage';
import LoginPage from './login/LoginPage';
import Page from './component/Page';
import Query from './component/Query';

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
    ({is_logged_in: isLoggedIn}) => {
      return isLoggedIn ? <HomePage /> : <LoginPage />
    },
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
