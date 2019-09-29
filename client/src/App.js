import React, {useCallback, useMemo} from 'react';

import HomePage from './home/HomePage';
import LoginPage from './login/LoginPage';
import Page from './component/Page';
import Query from './component/Query';

function App(props) {
  const errorRenderer = useCallback(
    (error) => <div>Error: {error.message}</div>,
    [],
  );
  const successRenderer = useCallback(
    ({isLoggedIn}) => (
      isLoggedIn ? <HomePage /> : <LoginPage />
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

export default React.memo(App);
