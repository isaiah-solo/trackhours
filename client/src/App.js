import React from 'react';

import HomePage from './HomePage';
import LoginPage from './LoginPage';
import {useFetchInitialData} from './api/useFetchInitialData';

const App = (props) => {
  const {error, isLoading, response} = useFetchInitialData('/api/check_login');
  if (isLoading || response === null) {
    return <div>Loading...</div>;
  } else if (error !== null) {
    return <div>Error: {error.message}</div>;
  }
  const {is_logged_in: isLoggedIn} = response;
  return (
    isLoggedIn ? <HomePage /> : <LoginPage />
  );
};

export default React.memo(App);
